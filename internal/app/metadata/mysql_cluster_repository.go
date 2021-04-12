package metadata

import (
	"fmt"
	"strconv"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"

	"github.com/romberli/log"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency"
)

const (
	addr   = "192.168.171.159:3306"
	dbName = "db_test"
	dbUser = "tester"
	dbPass = "mysql.1234"
)

var _ dependency.Repository = (*MySQLClusterRepo)(nil)

// MySQLClusterRepo implements Repository interface
type MySQLClusterRepo struct {
	Database middleware.Pool
}

// NewMySQLClusterRepo returns *MySQLClusterRepo with given middleware.Pool
func NewMySQLClusterRepo(db middleware.Pool) *MySQLClusterRepo {
	return &MySQLClusterRepo{db}
}

// NewMySQLClusterRepoWithGlobal returns *MySQLClusterRepo with global mysql pool
func NewMySQLClusterRepoWithGlobal() *MySQLClusterRepo {
	return NewMySQLClusterRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (mcr *MySQLClusterRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := mcr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata MySQLClusterRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns middleware.PoolConn
func (mcr *MySQLClusterRepo) Transaction() (middleware.Transaction, error) {
	return mcr.Database.Transaction()
}

// GetAll returns all available entities
func (mcr *MySQLClusterRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, cluster_name, middleware_cluster_id, monitor_system_id, 
			owner_id, owner_group, env_id, del_flag, create_time, last_update_time
		from t_meta_mysql_cluster_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata MySQLClusterRepo.GetAll() sql: \n%s", sql)

	result, err := mcr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*MySQLClusterInfo
	mysqlClusterInfoList := make([]*MySQLClusterInfo, result.RowNumber())
	for i := range mysqlClusterInfoList {
		mysqlClusterInfoList[i] = NewEmptyMySQLClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(mysqlClusterInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []dependency.Entity
	entityList := make([]dependency.Entity, result.RowNumber())
	for i := range entityList {
		entityList[i] = mysqlClusterInfoList[i]
	}

	return entityList, nil
}

// GetByID Select returns an available entity of the given id
func (mcr *MySQLClusterRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, cluster_name, middleware_cluster_id, monitor_system_id, 
			owner_id, owner_group, env_id, del_flag, create_time, last_update_time
		from t_meta_mysql_cluster_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := mcr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, fmt.Errorf("metadata MySQLClusterInfo.GetByID(): data does not exists, id: %s", id)
	case 1:
		mysqlClusterInfo := NewEmptyMySQLClusterInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return mysqlClusterInfo, nil
	default:
		return nil, fmt.Errorf("metadata MySQLClusterInfo.GetByID(): duplicate key exists, id: %s", id)
	}
}

// GetID checks identity of given entity from the middleware
func (mcr *MySQLClusterRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_mysql_cluster_info where del_flag = 0 and cluster_name = ?;`
	log.Debugf("metadata MySQLClusterRepo.GetID() select sql: %s", sql)
	result, err := mcr.Execute(sql, entity.(*MySQLClusterInfo).ClusterName)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (mcr *MySQLClusterRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `
		insert into t_meta_mysql_cluster_info(cluster_name,middleware_cluster_id,
			 monitor_system_id, owner_id, owner_group, env_id) 
		values(?,?,?,?,?,?);`
	log.Debugf("metadata MySQLClusterRepo.Create() insert sql: %s", sql)
	// execute
	_, err := mcr.Execute(sql,
		entity.(*MySQLClusterInfo).ClusterName,
		entity.(*MySQLClusterInfo).MiddlewareClusterID,
		entity.(*MySQLClusterInfo).MonitorSystemID,
		entity.(*MySQLClusterInfo).OwnerID,
		entity.(*MySQLClusterInfo).OwnerGroup,
		entity.(*MySQLClusterInfo).EnvID)
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
func (mcr *MySQLClusterRepo) Update(entity dependency.Entity) error {
	sql := `
		update t_meta_mysql_cluster_info set cluster_name = ?, middleware_cluster_id = ?, 
			monitor_system_id = ?, owner_id = ?, owner_group = ?, 
			env_id = ?, del_flag = ? 
		where id = ?;`
	log.Debugf("metadata MySQLClusterRepo.Update() update sql: %s", sql)
	mysqlClusterInfo := entity.(*MySQLClusterInfo)
	_, err := mcr.Execute(sql,
		mysqlClusterInfo.ClusterName,
		mysqlClusterInfo.MiddlewareClusterID,
		mysqlClusterInfo.MonitorSystemID,
		mysqlClusterInfo.OwnerID,
		mysqlClusterInfo.OwnerGroup,
		mysqlClusterInfo.EnvID,
		mysqlClusterInfo.DelFlag, mysqlClusterInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (mcr *MySQLClusterRepo) Delete(id string) error {
	sql := `update t_meta_mysql_cluster_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata MySQLClusterRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = mcr.Execute(sql, idInt)

	return err
}
