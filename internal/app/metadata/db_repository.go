package metadata

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"

	"github.com/romberli/log"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Repository = (*DBRepo)(nil)

type DBRepo struct {
	Database middleware.Pool
}

// NewDBRepo returns *DBRepo with given middleware.Pool
func NewDBRepo(db middleware.Pool) *DBRepo {
	return &DBRepo{db}
}

// NewDBRepo returns *DBRepo with global mysql pool
func NewDBRepoWithGlobal() *DBRepo {
	return NewDBRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (er *DBRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := er.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata DBRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

func (er *DBRepo) Transaction() (middleware.Transaction, error) {
	return er.Database.Transaction()
}

// GetAll returns all available entities
func (er *DBRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, db_name, cluster_id, cluster_type, owner_id, owner_group, env_id, del_flag, create_time, last_update_time
		from t_meta_db_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata DBRepo.GetAll() sql: \n%s", sql)

	result, err := er.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*DBInfo
	dbInfoList := make([]*DBInfo, result.RowNumber())
	for i := range dbInfoList {
		dbInfoList[i] = NewEmptyDBInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(dbInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []dependency.Entity
	entityList := make([]dependency.Entity, result.RowNumber())
	for i := range entityList {
		entityList[i] = dbInfoList[i]
	}

	return entityList, nil
}

func (er *DBRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, db_name, cluster_id, cluster_type, owner_id, owner_group, env_id, del_flag, create_time, last_update_time
		from t_meta_db_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata DBRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := er.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata DBInfo.GetByID(): data does not exists, id: %s", id))
	case 1:
		dbInfo := NewEmptyDBInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(dbInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return dbInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata DBInfo.GetByID(): duplicate key exists, id: %s", id))
	}
}

// GetID checks identity of given entity from the middleware
func (er *DBRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_db_info where del_flag = 0 and db_name = ? and owner_id = ? and env_id = ? order by id desc;`
	log.Debugf("metadata DBRepo.GetID() select sql: %s", sql)
	result, err := er.Execute(sql, entity.(*DBInfo).DBName, entity.(*DBInfo).OwnerID, entity.(*DBInfo).EnvID)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (er *DBRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `insert into t_meta_db_info(db_name, owner_id, env_id) values(?,?,?);`
	log.Debugf("metadata DBRepo.Create() insert sql: %s", sql)
	// execute
	_, err := er.Execute(sql, entity.(*DBInfo).DBName, entity.(*DBInfo).OwnerID, entity.(*DBInfo).EnvID)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := er.GetID(entity)
	if err != nil {
		return nil, err
	}
	// get entity
	return er.GetByID(id)
}

// Update updates data with given entity in the middleware
func (er *DBRepo) Update(entity dependency.Entity) error {
	sql := `update t_meta_db_info set db_name = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata DBRepo.Update() update sql: %s", sql)
	dbInfo := entity.(*DBInfo)
	_, err := er.Execute(sql, dbInfo.DBName, dbInfo.DelFlag, dbInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (er *DBRepo) Delete(id string) error {
	sql := `update t_meta_db_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata DBRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = er.Execute(sql, idInt)

	return err
}
