package metadata

import (
	"errors"
	"fmt"
	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/log"
	"strconv"
)

var _ dependency.Repository = (*MiddlewareServerRepo)(nil)

type MiddlewareServerRepo struct {
	Database middleware.Pool
}

// NewMiddlewareServerRepo returns *MiddlewareServerRepo with given middleware.Pool
func NewMiddlewareServerRepo(db middleware.Pool) *MiddlewareServerRepo {
	return &MiddlewareServerRepo{db}
}

// NewMiddlewareServerRepo returns *MiddlewareServerRepo with global mysql pool
func NewMiddlewareServerRepoWithGlobal() *MiddlewareServerRepo {
	return NewMiddlewareServerRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (er *MiddlewareServerRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := er.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata MiddlewareServerRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

func (msr *MiddlewareServerRepo) Transaction() (middleware.Transaction, error) {
	return msr.Database.Transaction()
}

// GetAll returns all available entities
func (msr *MiddlewareServerRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
      select id, cluster_id_middleware, server_name, middleware_role, host_ip, port_num, del_flag, create_time, last_update_time
      from t_meta_middleware_server_info
      where del_flag = 0
      order by id;
   `
	log.Debugf("metadata MiddlewareServerRepo.GetAll() sql: \n%s", sql)

	result, err := msr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*MiddlewareServerInfo
	middlewareServerInfoList := make([]*MiddlewareServerInfo, result.RowNumber())
	for i := range middlewareServerInfoList {
		middlewareServerInfoList[i] = NewEmptyMiddlewareServerInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(middlewareServerInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []dependency.Entity
	entityList := make([]dependency.Entity, result.RowNumber())
	for i := range entityList {
		entityList[i] = middlewareServerInfoList[i]
	}

	return entityList, nil
}

func (msr *MiddlewareServerRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
      select id, cluster_id_middleware, server_name, middleware_role, host_ip, port_num, del_flag, create_time, last_update_time
      from t_meta_middleware_server_info
      where del_flag = 0
      and id = ?;
   `
	log.Debugf("metadata MiddlewareServerRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := msr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata MiddlewareServerInfo.GetByID(): data does not exists, id: %s", id))
	case 1:
		middlewareServerInfo := NewEmptyMiddlewareServerInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(middlewareServerInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return middlewareServerInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata MiddlewareServerInfo.GetByID(): duplicate key exists, id: %s", id))
	}
}

// GetID checks identity of given entity from the middleware
func (msr *MiddlewareServerRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_middleware_server_info where del_flag = 0 and server_name = ?;`
	log.Debugf("metadata MiddlewareServerRepo.GetID() select sql: %s", sql)
	result, err := msr.Execute(sql, entity.(*MiddlewareServerInfo).ServerName)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (msr *MiddlewareServerRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `insert into t_meta_middleware_server_info(cluster_id_middleware, server_name, middleware_role, host_ip, port_num) values(?, ?, ?, ?, ?);`
	log.Debugf("metadata MiddlewareServerRepo.Create() insert sql: %s", sql)
	// execute
	_, err := msr.Execute(sql, entity.(*MiddlewareServerInfo).ClusterIDMiddleware, entity.(*MiddlewareServerInfo).ServerName, entity.(*MiddlewareServerInfo).MiddlewareRole, entity.(*MiddlewareServerInfo).HostIP, entity.(*MiddlewareServerInfo).PortNum)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := msr.GetID(entity)
	if err != nil {
		return nil, err
	}
	// get entity
	return msr.GetByID(id)
}

// Update updates data with given entity in the middleware
func (msr *MiddlewareServerRepo) Update(entity dependency.Entity) error {
	sql := `update t_meta_middleware_server_info set cluster_id_middleware = ?, server_name = ?, middleware_role = ?, host_ip = ?, port_num = ? where id = ?;`
	log.Debugf("metadata MiddlewareServerRepo.Update() update sql: %s", sql)
	middlewareServerInfo := entity.(*MiddlewareServerInfo)
	_, err := msr.Execute(sql, middlewareServerInfo.ClusterIDMiddleware, middlewareServerInfo.ServerName, middlewareServerInfo.MiddlewareRole, middlewareServerInfo.HostIP, middlewareServerInfo.PortNum, middlewareServerInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (msr *MiddlewareServerRepo) Delete(id string) error {
	sql := `update t_meta_middleware_server_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata MiddlewareServerRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = msr.Execute(sql, idInt)

	return err
}
