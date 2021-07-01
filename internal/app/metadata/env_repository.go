package metadata

import (
	"errors"
	"fmt"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"

	"github.com/romberli/log"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/metadata"
)

var _ metadata.EnvRepo = (*EnvRepo)(nil)

type EnvRepo struct {
	Database middleware.Pool
}

// NewEnvRepo returns *EnvRepo with given middleware.Pool
func NewEnvRepo(db middleware.Pool) *EnvRepo {
	return &EnvRepo{db}
}

// NewEnvRepoWithGlobal returns *EnvRepo with global mysql pool
func NewEnvRepoWithGlobal() *EnvRepo {
	return NewEnvRepo(global.DASMySQLPool)
}

// Execute executes given command and placeholders on the middleware
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

// Transaction returns middleware.PoolConn, so it can run multiple statements in the same transaction
func (er *EnvRepo) Transaction() (middleware.Transaction, error) {
	return er.Database.Transaction()
}

// GetAll gets all environments from the middleware
func (er *EnvRepo) GetAll() ([]metadata.Env, error) {
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
	entityList := make([]metadata.Env, result.RowNumber())
	for i := range entityList {
		entityList[i] = envInfoList[i]
	}

	return entityList, nil
}

// GetByID gets an environment by the identity from the middleware
func (er *EnvRepo) GetByID(id int) (metadata.Env, error) {
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
		return nil, errors.New(fmt.Sprintf("metadata EnvInfo.GetByID(): data does not exists, id: %d", id))
	case 1:
		envInfo := NewEmptyEnvInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(envInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return envInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata EnvInfo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetID gets the identity with given environment name from the middleware
func (er *EnvRepo) GetID(envName string) (int, error) {
	sql := `select id from t_meta_env_info where del_flag = 0 and env_name = ?;`
	log.Debugf("metadata EnvRepo.GetID() select sql: %s", sql)
	result, err := er.Execute(sql, envName)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// Create creates an environment in the middleware
func (er *EnvRepo) Create(env metadata.Env) (metadata.Env, error) {
	sql := `insert into t_meta_env_info(env_name) values(?);`
	log.Debugf("metadata EnvRepo.Create() insert sql: %s", sql)
	// execute
	_, err := er.Execute(sql, env.GetEnvName())
	if err != nil {
		return nil, err
	}
	// get id
	id, err := er.GetID(env.GetEnvName())
	if err != nil {
		return nil, err
	}
	// get env
	return er.GetByID(id)
}

// Update updates the environment in the middleware
func (er *EnvRepo) Update(env metadata.Env) error {
	sql := `update t_meta_env_info set env_name = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata EnvRepo.Update() update sql: %s", sql)
	_, err := er.Execute(sql, env.GetEnvName(), env.GetDelFlag(), env.Identity())

	return err
}

// Delete deletes the environment in the middleware
func (er *EnvRepo) Delete(id int) error {
	tx, err := er.Transaction()
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Close()
		if err != nil {
			log.Errorf("metadata EnvRepo.Delete(): close database connection failed.\n%s", err.Error())
		}
	}()

	err = tx.Begin()
	if err != nil {
		return err
	}
	//delete from t_meta_env_info where id =
	sql := `delete from t_meta_env_info where id = ?;`
	log.Debugf("metadata EnvRepo.Delete() delete sql(t_meta_app_info): %s", sql)
	_, err = tx.Execute(sql, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// GetEnvByName gets Env of given environment name
func (er *EnvRepo) GetEnvByName(envName string) (metadata.Env, error) {
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
