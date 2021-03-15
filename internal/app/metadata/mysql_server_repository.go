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

var _ dependency.Repository = (*MySQLServerRepo)(nil)

// MySQLServerRepo implements Repository interface
type MySQLServerRepo struct {
	Database middleware.Pool
}

// NewMySQLServerRepo returns *MySQLServerRepo with given middleware.Pool
func NewMySQLServerRepo(db middleware.Pool) *MySQLServerRepo {
	return &MySQLServerRepo{db}
}

// NewMySQLServerRepoWithGlobal returns *MySQLServerRepo with global mysql pool
func NewMySQLServerRepoWithGlobal() *MySQLServerRepo {
	return NewMySQLServerRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (msr *MySQLServerRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := msr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata MySQLServerRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns middleware.PoolConn
func (msr *MySQLServerRepo) Transaction() (middleware.Transaction, error) {
	return msr.Database.Transaction()
}

// GetAll returns all available entities
func (msr *MySQLServerRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, cluster_id, host_ip, port_num, deployment_type, version, del_flag, 
			create_time, last_update_time
		from t_meta_mysql_server_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata MySQLServerRepo.GetAll() sql: \n%s", sql)

	result, err := msr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*MySQLServerInfo
	mysqlClusterInfoList := make([]*MySQLServerInfo, result.RowNumber())
	for i := range mysqlClusterInfoList {
		mysqlClusterInfoList[i] = NewEmptyMySQLServerInfoWithGlobal()
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
func (msr *MySQLServerRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, cluster_id, host_ip, port_num, deployment_type, version, del_flag, 
			create_time, last_update_time
		from t_meta_mysql_server_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata MySQLServerRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := msr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, fmt.Errorf("metadata MySQLServerInfo.GetByID(): data does not exists, id: %s", id)
	case 1:
		mysqlClusterInfo := NewEmptyMySQLServerInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return mysqlClusterInfo, nil
	default:
		return nil, fmt.Errorf("metadata MySQLServerInfo.GetByID(): duplicate key exists, id: %s", id)
	}
}

// GetID checks identity of given entity from the middleware
func (msr *MySQLServerRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_mysql_server_info where del_flag = 0 and host_ip = ? and port_num = ?;`
	log.Debugf("metadata MySQLServerRepo.GetID() select sql: %s", sql)
	result, err := msr.Execute(sql, entity.(*MySQLServerInfo).HostIP, entity.(*MySQLServerInfo).PortNum)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (msr *MySQLServerRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `
		insert into t_meta_mysql_server_info(
			cluster_id, host_ip, port_num, deployment_type, version) 
		values(?, ?, ?, ?, ?);`
	log.Debugf("metadata MySQLServerRepo.Create() insert sql: %s", sql)
	// execute
	_, err := msr.Execute(sql,
		entity.(*MySQLServerInfo).ClusterID,
		entity.(*MySQLServerInfo).HostIP,
		entity.(*MySQLServerInfo).PortNum,
		entity.(*MySQLServerInfo).DeploymentType,
		entity.(*MySQLServerInfo).Version,
	)
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
func (msr *MySQLServerRepo) Update(entity dependency.Entity) error {
	sql := `
		update t_meta_mysql_server_info set 
			cluster_id = ?, host_ip = ?, port_num = ?, deployment_type = ?, 
			version = ?, del_flag = ? 
		where id = ?;`
	log.Debugf("metadata MySQLServerRepo.Update() update sql: %s", sql)
	mysqlServerInfo := entity.(*MySQLServerInfo)
	_, err := msr.Execute(sql,
		mysqlServerInfo.ClusterID,
		mysqlServerInfo.HostIP,
		mysqlServerInfo.PortNum,
		mysqlServerInfo.DeploymentType,
		mysqlServerInfo.Version,
		mysqlServerInfo.DelFlag,
		mysqlServerInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (msr *MySQLServerRepo) Delete(id string) error {
	sql := `update t_meta_mysql_server_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata MySQLServerRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = msr.Execute(sql, idInt)

	return err
}
