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

var _ dependency.Repository = (*MYSQLServerRepo)(nil)

// MYSQLServerRepo implements Repository interface
type MYSQLServerRepo struct {
	Database middleware.Pool
}

// NewMYSQLServerRepo returns *MYSQLServerRepo with given middleware.Pool
func NewMYSQLServerRepo(db middleware.Pool) *MYSQLServerRepo {
	return &MYSQLServerRepo{db}
}

// NewMYSQLServerRepoWithGlobal returns *MYSQLServerRepo with global mysql pool
func NewMYSQLServerRepoWithGlobal() *MYSQLServerRepo {
	return NewMYSQLServerRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (er *MYSQLServerRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := er.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata MYSQLServerRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns middleware.PoolConn
func (er *MYSQLServerRepo) Transaction() (middleware.Transaction, error) {
	return er.Database.Transaction()
}

// GetAll returns all available entities
func (er *MYSQLServerRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, cluster_id, host_ip, port_num, deployment_type, version, del_flag, 
			create_time, last_update_time
		from t_meta_mysql_server_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata MYSQLServerRepo.GetAll() sql: \n%s", sql)

	result, err := er.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*MYSQLServerInfo
	mysqlClusterInfoList := make([]*MYSQLServerInfo, result.RowNumber())
	for i := range mysqlClusterInfoList {
		mysqlClusterInfoList[i] = NewEmptyMYSQLServerInfoWithGlobal()
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
func (er *MYSQLServerRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, cluster_id, host_ip, port_num, deployment_type, version, del_flag, 
			create_time, last_update_time
		from t_meta_mysql_server_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata MYSQLServerRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := er.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, fmt.Errorf("metadata MYSQLServerInfo.GetByID(): data does not exists, id: %s", id)
	case 1:
		mysqlClusterInfo := NewEmptyMYSQLServerInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return mysqlClusterInfo, nil
	default:
		return nil, fmt.Errorf("metadata MYSQLServerInfo.GetByID(): duplicate key exists, id: %s", id)
	}
}

// GetID checks identity of given entity from the middleware
func (er *MYSQLServerRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_mysql_server_info where del_flag = 0 and host_ip = ? and port_num = ?;`
	log.Debugf("metadata MYSQLServerRepo.GetID() select sql: %s", sql)
	result, err := er.Execute(sql, entity.(*MYSQLServerInfo).HostIP, entity.(*MYSQLServerInfo).PortNum)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (er *MYSQLServerRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `
		insert into t_meta_mysql_server_info(
			cluster_id, host_ip, port_num, deployment_type, version) 
		values(?, ?, ?, ?, ?);`
	log.Debugf("metadata MYSQLServerRepo.Create() insert sql: %s", sql)
	// execute
	_, err := er.Execute(sql,
		entity.(*MYSQLServerInfo).ClusterID,
		entity.(*MYSQLServerInfo).HostIP,
		entity.(*MYSQLServerInfo).PortNum,
		entity.(*MYSQLServerInfo).DeploymentType,
		entity.(*MYSQLServerInfo).Version,
	)
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
func (er *MYSQLServerRepo) Update(entity dependency.Entity) error {
	sql := `
		update t_meta_mysql_server_info set 
			cluster_id = ?, host_ip = ?, port_num = ?, deployment_type = ?, 
			version = ?, del_flag = ? 
		where id = ?;`
	log.Debugf("metadata MYSQLServerRepo.Update() update sql: %s", sql)
	mysqlServerInfo := entity.(*MYSQLServerInfo)
	_, err := er.Execute(sql,
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
func (er *MYSQLServerRepo) Delete(id string) error {
	sql := `update t_meta_mysql_server_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata MYSQLServerRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = er.Execute(sql, idInt)

	return err
}
