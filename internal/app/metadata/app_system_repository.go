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

var _ dependency.Repository = (*AppSystemRepo)(nil)

type AppSystemRepo struct {
	Database middleware.Pool
}

// NewAppSystemRepo returns *AppSystemRepo with given middleware.Pool
func NewAppSystemRepo(db middleware.Pool) *AppSystemRepo {
	return &AppSystemRepo{db}
}

// NewAppSystemRepoWithGlobal returns *AppSystemRepo with global mysql pool
func NewAppSystemRepoWithGlobal() *AppSystemRepo {
	return NewAppSystemRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (asr *AppSystemRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := asr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata AppSystemRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

func (asr *AppSystemRepo) Transaction() (middleware.Transaction, error) {
	return asr.Database.Transaction()
}

// GetAll returns all available entities
func (asr *AppSystemRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, system_name, del_flag, create_time, last_update_time, level,owner_id, owner_group
		from t_meta_app_system_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata AppSystemRepo.GetAll() sql: \n%s", sql)

	result, err := asr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*AppSystemInfo
	appSystemInfoList := make([]*AppSystemInfo, result.RowNumber())
	for i := range appSystemInfoList {
		appSystemInfoList[i] = NewEmptyAppSystemInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(appSystemInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []dependency.Entity
	entityList := make([]dependency.Entity, result.RowNumber())
	for i := range entityList {
		entityList[i] = appSystemInfoList[i]
	}

	return entityList, nil
}

func (asr *AppSystemRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, system_name, del_flag, create_time, last_update_time,level,owner_id,owner_group
		from t_meta_app_system_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata AppSystemRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := asr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata AppSystemInfo.GetByID(): data does not exists, id: %s", id))
	case 1:
		appSystemInfo := NewEmptyAppSystemInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(appSystemInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return appSystemInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata AppSystemInfo.GetByID(): duplicate key exists, id: %s", id))
	}
}

// GetID checks identity of given entity from the middleware
func (asr *AppSystemRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_app_system_info where del_flag = 0 and system_name = ?;`
	log.Debugf("metadata AppSystemRepo.GetID() select sql: %s", sql)
	result, err := asr.Execute(sql, entity.(*AppSystemInfo).AppSystemName)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (asr *AppSystemRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `insert into t_meta_app_system_info(system_name,level,owner_id,owner_group) values(?,?,?,?);`
	log.Debugf("metadata AppSystemRepo.Create() insert sql: %s", sql)
	// execute
	_, err := asr.Execute(sql, entity.(*AppSystemInfo).AppSystemName, entity.(*AppSystemInfo).Level, entity.(*AppSystemInfo).OwnerID, entity.(*AppSystemInfo).OwnerGroup)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := asr.GetID(entity)
	if err != nil {
		return nil, err
	}
	// get entity
	return asr.GetByID(id)
}

// Update updates data with given entity in the middleware
func (asr *AppSystemRepo) Update(entity dependency.Entity) error {
	sql := `update t_meta_app_system_info set system_name = ?, del_flag = ?,level = ?,owner_id = ?,owner_group = ? where id = ?;`
	log.Debugf("metadata AppSystemRepo.Update() update sql: %s", sql)
	appSystemInfo := entity.(*AppSystemInfo)
	_, err := asr.Execute(sql, appSystemInfo.AppSystemName, appSystemInfo.DelFlag, appSystemInfo.Level, appSystemInfo.OwnerID, appSystemInfo.OwnerGroup, appSystemInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (asr *AppSystemRepo) Delete(id string) error {
	sql := `update t_meta_app_system_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata AppSystemRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = asr.Execute(sql, idInt)

	return err
}
