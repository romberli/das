package metadata

import (
	"errors"
	"fmt"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/log"

	"github.com/romberli/das/global"
)

var _ metadata.MiddlewareClusterRepo = (*MiddlewareClusterRepo)(nil)

type MiddlewareClusterRepo struct {
	Database middleware.Pool
}

// NewMiddlewareClusterRepo returns *MiddlewareClusterRepo with given middleware.Pool
func NewMiddlewareClusterRepo(db middleware.Pool) *MiddlewareClusterRepo {
	return &MiddlewareClusterRepo{db}
}

// NewMiddlewareClusterRepo returns *MiddlewareClusterRepo with global mysql pool
func NewMiddlewareClusterRepoWithGlobal() *MiddlewareClusterRepo {
	return NewMiddlewareClusterRepo(global.DASMySQLPool)
}

// Execute executes command with arguments on database
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

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (mcr *MiddlewareClusterRepo) Transaction() (middleware.Transaction, error) {
	return mcr.Database.Transaction()
}

// GetAll gets all middleware clusters from the middleware
func (mcr *MiddlewareClusterRepo) GetAll() ([]metadata.MiddlewareCluster, error) {
	sql := `
		select id, cluster_name, owner_id, env_id, del_flag, create_time, last_update_time
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
	entityList := make([]metadata.MiddlewareCluster, result.RowNumber())
	for i := range entityList {
		entityList[i] = middlewareClusterInfoList[i]
	}

	return entityList, nil
}

// GetByEnv gets middleware clusters of given env id from the middleware
func (mcr *MiddlewareClusterRepo) GetByEnv(envID int) ([]metadata.MiddlewareCluster, error) {
	sql := `
		select id, cluster_name, owner_id, env_id, del_flag, create_time, last_update_time
		from t_meta_middleware_cluster_info
		where del_flag = 0
		and env_id = ?
		order by id;
	`
	log.Debugf("metadata MiddlewareClusterRepo.GetByEnv() sql: \n%s", sql, envID)

	result, err := mcr.Execute(sql, envID)
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
	// init []metadata.MiddlewareCluster
	middlewareClusterList := make([]metadata.MiddlewareCluster, result.RowNumber())
	for i := range middlewareClusterList {
		middlewareClusterList[i] = middlewareClusterInfoList[i]
	}

	return middlewareClusterList, nil
}

// GetByID gets a middleware cluster by the identity from the middleware
func (mcr *MiddlewareClusterRepo) GetByID(id int) (metadata.MiddlewareCluster, error) {
	sql := `
		select id, cluster_name, owner_id, env_id, del_flag, create_time, last_update_time
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
		return nil, errors.New(fmt.Sprintf("metadata MiddlewareClusterInfo.GetByID(): data does not exists, id: %d", id))
	case 1:
		middlewareClusterInfo := NewEmptyMiddlewareClusterInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(middlewareClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return middlewareClusterInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata MiddlewareClusterInfo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetByName gets a middleware cluster of given cluster name from the middle ware
func (mcr *MiddlewareClusterRepo) GetByName(clusterName string) (metadata.MiddlewareCluster, error) {
	sql := `select id from t_meta_middleware_cluster_info where del_flag = 0 and cluster_name = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.GetByName() select sql: %s", sql)
	result, err := mcr.Execute(sql, clusterName)
	if err != nil {
		return nil, err
	}
	id, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	if err != nil {
		return nil, err
	}
	return mcr.GetByID(id)
}

// GetID gets the identity with given cluster name and env id from the middleware
func (mcr *MiddlewareClusterRepo) GetID(clusterName string, envID int) (int, error) {
	sql := `select id from t_meta_middleware_cluster_info where del_flag = 0 and cluster_name = ? and env_id = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.GetID() select sql: %s", sql)
	result, err := mcr.Execute(sql, clusterName, envID)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// GetMiddlewareServerIDList gets the middleware server id list of given cluster id from the middle ware
func (mcr *MiddlewareClusterRepo) GetMiddlewareServerIDList(clusterID int) ([]int, error) {
	sql := `select id from t_meta_middleware_server_info
            where del_flag = 0
            and cluster_id = ?
		    order by id;
	`
	log.Debugf("metadata MiddlewareCLusterRepo.GetMiddlewareServerIDList() select sql: %s", sql)
	result, err := mcr.Execute(sql, clusterID)
	if err != nil {
		return nil, err
	}
	resultNum := result.RowNumber()
	serverIDList := make([]int, resultNum)
	for row := 0; row < resultNum; row++ {
		serverID, err := result.GetInt(row, constant.ZeroInt)
		if err != nil {
			return nil, err
		}
		serverIDList[row] = serverID
	}
	return serverIDList, nil
}

// Create creates data with given entity in the middleware
func (mcr *MiddlewareClusterRepo) Create(middlewareCluster metadata.MiddlewareCluster) (metadata.MiddlewareCluster, error) {
	sql := `insert into t_meta_middleware_cluster_info(cluster_name, owner_id, env_id) values(?, ?, ?);`
	log.Debugf("metadata MiddlewareClusterRepo.Create() insert sql: %s", sql)
	// execute
	_, err := mcr.Execute(sql,
		middlewareCluster.(*MiddlewareClusterInfo).ClusterName,
		middlewareCluster.(*MiddlewareClusterInfo).OwnerID,
		middlewareCluster.(*MiddlewareClusterInfo).EnvID,
	)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := mcr.GetID(middlewareCluster.GetClusterName(), middlewareCluster.GetEnvID())
	if err != nil {
		return nil, err
	}
	// get entity
	return mcr.GetByID(id)
}

// Update updates data with given entity in the middleware
func (mcr *MiddlewareClusterRepo) Update(middlewareCluster metadata.MiddlewareCluster) error {
	sql := `update t_meta_middleware_cluster_info set cluster_name = ?, owner_id = ?, env_id = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.Update() update sql: %s", sql)
	_, err := mcr.Execute(sql,
		middlewareCluster.GetClusterName(),
		middlewareCluster.GetOwnerID(),
		middlewareCluster.GetEnvID(),
		middlewareCluster.GetDelFlag(),
		middlewareCluster.Identity(),
	)

	return err
}

// Delete deletes the middleware cluster in the middleware
func (mcr *MiddlewareClusterRepo) Delete(id int) error {
	tx, err := mcr.Transaction()
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Close()
		if err != nil {
			log.Errorf("metadata MiddlewareClusterRepo.Delete(): close database connection failed.\n%s", err.Error())
		}
	}()

	err = tx.Begin()
	if err != nil {
		return err
	}
	sql := `delete from t_meta_middleware_cluster_info where id = ?;`
	log.Debugf("metadata MiddlewareClusterRepo.Delete() update sql: %s", sql)
	_, err = mcr.Execute(sql, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
