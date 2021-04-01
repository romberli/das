package metadata

import (
	"errors"
	"fmt"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"

	"github.com/romberli/log"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Repository = (*EnvRepo)(nil)

type EnvRepo struct {
	Database middleware.Pool
}

// NewEnvRepo returns *EnvRepo with given middleware.Pool
func NewEnvRepo(db middleware.Pool) *EnvRepo {
	return &EnvRepo{db}
}

// NewEnvRepoWithGlobal returns *EnvRepo with global mysql pool
func NewEnvRepoWithGlobal() *EnvRepo {
	return NewEnvRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (er *EnvRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := er.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata EnvRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction implements dependency.Repository interface
func (er *EnvRepo) Transaction() (middleware.Transaction, error) {
	return er.Database.Transaction()
}

// GetAll returns all available entities
func (er *EnvRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, env_name, del_flag, create_time, last_update_time
		from t_meta_env_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata EnvRepo.GetAll() sql: \n%s", sql)

	result, err := er.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*EnvInfo
	envInfoList := make([]*EnvInfo, result.RowNumber())
	for i := range envInfoList {
		envInfoList[i] = NewEmptyEnvInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(envInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []dependency.
	entityList := make([]dependency.Entity, result.RowNumber())
	for i := range entityList {
		entityList[i] = envInfoList[i]
	}

	return entityList, nil
}

func (er *EnvRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, env_name, del_flag, create_time, last_update_time
		from t_meta_env_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata EnvRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := er.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata EnvInfo.GetByID(): data does not exists, id: %s", id))
	case 1:
		envInfo := NewEmptyEnvInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(envInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return envInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata EnvInfo.GetByID(): duplicate key exists, id: %s", id))
	}
}

// GetID checks identity of given entity from the middleware
func (er *EnvRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_env_info where del_flag = 0 and env_name = ?;`
	log.Debugf("metadata EnvRepo.GetID() select sql: %s", sql)
	envName, err := entity.Get(envNameStruct)
	if err != nil {
		return constant.EmptyString, err
	}
	result, err := er.Execute(sql, envName)
	if err != nil {
		return constant.EmptyString, err
	}

	id, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	if err != nil {
		return constant.EmptyString, err
	}

	return id, err
}

// Create creates data with given entity in the middleware
func (er *EnvRepo) Create(env dependency.Entity) (dependency.Entity, error) {
	sql := `insert into t_meta_env_info(env_name) values(?);`
	log.Debugf("metadata EnvRepo.Create() insert sql: %s", sql)
	// execute
	envName, err := env.Get(envNameStruct)
	if err != nil {
		return nil, err
	}
	_, err = er.Execute(sql, envName)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := er.GetID(env)
	if err != nil {
		return nil, err
	}
	// get env
	return er.GetByID(id)
}

// Update updates data with given entity in the middleware
func (er *EnvRepo) Update(env dependency.Entity) error {
	sql := `update t_meta_env_info set env_name = ? where id = ?;`
	log.Debugf("metadata EnvRepo.Update() update sql: %s", sql)
	envName, err := env.Get(envNameStruct)
	if err != nil {
		return err
	}

	_, err = er.Execute(sql, envName, env.Identity())

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (er *EnvRepo) Delete(id string) error {
	sql := `update t_meta_env_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata EnvRepo.Delete() update sql: %s", sql)
	_, err := er.Execute(sql, id)

	return err
}

func (er *EnvRepo) GetEnvByName(envName string) (dependency.Entity, error) {
	sql := `
		select id, env_name, del_flag, create_time, last_update_time
		from t_meta_env_info
		where del_flag = 0
		and env_name = ?;
	`
	log.Debugf("metadata EnvRepo.GetEnvByName() sql: \n%s\nplaceholders: %s", sql, envName)

	result, err := er.Execute(sql, envName)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata EnvInfo.GetEnvByName(): data does not exists, env_name: %s", envName))
	case 1:
		envInfo := NewEmptyEnvInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(envInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return envInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata EnvInfo.GetEnvByName(): duplicate key exists, env_name: %s", envName))
	}
}
