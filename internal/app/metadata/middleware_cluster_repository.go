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

var _ dependency.Repository = (*MiddlewareClusterRepo)(nil)

type MiddlewareClusterRepo struct {
	Database middleware.Pool
}

// NewMiddlewareClusterRepo returns *MiddlewareClusterRepo with given middleware.Pool
func NewMiddlewareClusterRepo(db middleware.Pool) *MiddlewareClusterRepo {
	return &MiddlewareClusterRepo{db}
}

// NewMiddlewareClusterRepo returns *MiddlewareClusterRepo with global mysql pool
func NewMiddlewareClusterRepoWithGlobal() *MiddlewareClusterRepo {
	return NewMiddlewareClusterRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (mcr *MiddlewareClusterRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := mcr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata MiddlewareClusterRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

func (mcr *MiddlewareClusterRepo) Transaction() (middleware.Transaction, error) {
	return mcr.Database.Transaction()
}

// GetAll returns all available entities
func (mcr *MiddlewareClusterRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, cluster_name, env_id, del_flag, create_time, last_update_time
		from t_meta_middleware_cluster_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata MiddlewareClusterRepo.GetAll() sql: \n%s", sql)

	result, err := mcr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*MiddlewareClusterInfo
	middlewareClusterInfoList := make([]*MiddlewareClusterInfo, result.RowNumber())
	for i := range middlewareClusterInfoList {
		middlewareClusterInfoList[i] = NewEmptyMiddlewareClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(middlewareClusterInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []dependency.Entity
	entityList := make([]dependency.Entity, result.RowNumber())
	for i := range entityList {
		entityList[i] = middlewareClusterInfoList[i]
	}

	return entityList, nil
}

func (mcr *MiddlewareClusterRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, cluster_name, env_id, del_flag, create_time, last_update_time
		from t_meta_middleware_cluster_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata MiddlewareClusterRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := mcr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata MiddlewareClusterInfo.GetByID(): data does not exists, id: %s", id))
	case 1:
		middlewareClusterInfo := NewEmptyMiddlewareClusterInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(middlewareClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return middlewareClusterInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata MiddlewareClusterInfo.GetByID(): duplicate key exists, id: %s", id))
	}
}

// GetID checks identity of given entity from the middleware
func (mcr *MiddlewareClusterRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_middleware_cluster_info where del_flag = 0 and cluster_name = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.GetID() select sql: %s", sql)
	result, err := mcr.Execute(sql, entity.(*MiddlewareClusterInfo).ClusterName)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (mcr *MiddlewareClusterRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `insert into t_meta_middleware_cluster_info(cluster_name, env_id) values(?, ?);`
	log.Debugf("metadata MiddlewareClusterRepo.Create() insert sql: %s", sql)
	// execute
	_, err := mcr.Execute(sql, entity.(*MiddlewareClusterInfo).ClusterName, entity.(*MiddlewareClusterInfo).EnvID)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := mcr.GetID(entity)
	if err != nil {
		return nil, err
	}
	// get entity
	return mcr.GetByID(id)
}

// Update updates data with given entity in the middleware
func (mcr *MiddlewareClusterRepo) Update(entity dependency.Entity) error {
	sql := `update t_meta_middleware_cluster_info set cluster_name = ?, env_id = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.Update() update sql: %s", sql)
	middlewareClusterInfo := entity.(*MiddlewareClusterInfo)
	_, err := mcr.Execute(sql, middlewareClusterInfo.ClusterName, middlewareClusterInfo.EnvID, middlewareClusterInfo.DelFlag, middlewareClusterInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (mcr *MiddlewareClusterRepo) Delete(id string) error {
	sql := `update t_meta_middleware_cluster_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = mcr.Execute(sql, idInt)

	return err
}
