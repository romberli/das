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

var _ dependency.Repository = (*DbRepo)(nil)

type DbRepo struct {
	Database middleware.Pool
}

// NewDbRepo returns *DbRepo with given middleware.Pool
func NewDbRepo(db middleware.Pool) *DbRepo {
	return &DbRepo{db}
}

// NewDbRepo returns *DbRepo with global mysql pool
func NewDbRepoWithGlobal() *DbRepo {
	return NewDbRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (er *DbRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := er.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata DbRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

func (er *DbRepo) Transaction() (middleware.Transaction, error) {
	return er.Database.Transaction()
}

// GetAll returns all available entities
func (er *DbRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, db_name, cluster_id, cluster_type, owner_id, owner_group, env_id, del_flag, create_time, last_update_time
		from t_meta_db_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata DbRepo.GetAll() sql: \n%s", sql)

	result, err := er.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*DbInfo
	dbInfoList := make([]*DbInfo, result.RowNumber())
	for i := range dbInfoList {
		dbInfoList[i] = NewEmptyDbInfoWithGlobal()
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

func (er *DbRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, db_name, cluster_id, cluster_type, owner_id, owner_group, env_id, del_flag, create_time, last_update_time
		from t_meta_db_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata DbRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := er.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata DbInfo.GetByID(): data does not exists, id: %s", id))
	case 1:
		dbInfo := NewEmptyDbInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(dbInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return dbInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata DbInfo.GetByID(): duplicate key exists, id: %s", id))
	}
}

// GetID checks identity of given entity from the middleware
func (er *DbRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_db_info where del_flag = 0 and db_name = ? and owner_id = ? and env_id = ? order by id desc;`
	log.Debugf("metadata DbRepo.GetID() select sql: %s", sql)
	result, err := er.Execute(sql, entity.(*DbInfo).DbName, entity.(*DbInfo).OwnerId, entity.(*DbInfo).EnvId)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (er *DbRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `insert into t_meta_db_info(db_name, owner_id, env_id) values(?,?,?);`
	log.Debugf("metadata DbRepo.Create() insert sql: %s", sql)
	// execute
	_, err := er.Execute(sql, entity.(*DbInfo).DbName, entity.(*DbInfo).OwnerId, entity.(*DbInfo).EnvId)
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
func (er *DbRepo) Update(entity dependency.Entity) error {
	sql := `update t_meta_db_info set db_name = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata DbRepo.Update() update sql: %s", sql)
	dbInfo := entity.(*DbInfo)
	_, err := er.Execute(sql, dbInfo.DbName, dbInfo.DelFlag, dbInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (er *DbRepo) Delete(id string) error {
	sql := `update t_meta_db_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata DbRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = er.Execute(sql, idInt)

	return err
}
