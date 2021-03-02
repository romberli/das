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

var _ dependency.Repository = (*MYSQLClusterRepo)(nil)

// MYSQLClusterRepo implements Repository interface
type MYSQLClusterRepo struct {
	Database middleware.Pool
}

// NewMYSQLClusterRepo returns *MYSQLClusterRepo with given middleware.Pool
func NewMYSQLClusterRepo(db middleware.Pool) *MYSQLClusterRepo {
	return &MYSQLClusterRepo{db}
}

// NewMYSQLClusterRepoWithGlobal returns *MYSQLClusterRepo with global mysql pool
func NewMYSQLClusterRepoWithGlobal() *MYSQLClusterRepo {
	return NewMYSQLClusterRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (er *MYSQLClusterRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := er.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata MYSQLClusterRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns middleware.PoolConn
func (er *MYSQLClusterRepo) Transaction() (middleware.Transaction, error) {
	return er.Database.Transaction()
}

// GetAll returns all available entities
func (er *MYSQLClusterRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, cluster_name, middleware_cluster_id, monitor_system_id, 
			owner_id, owner_group, env_id, del_flag, create_time, last_update_time
		from t_meta_mysql_cluster_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata MYSQLClusterRepo.GetAll() sql: \n%s", sql)

	result, err := er.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*MYSQLClusterInfo
	mysqlClusterInfoList := make([]*MYSQLClusterInfo, result.RowNumber())
	for i := range mysqlClusterInfoList {
		mysqlClusterInfoList[i] = NewEmptyMYSQLClusterInfoWithGlobal()
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
func (er *MYSQLClusterRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, cluster_name, middleware_cluster_id, monitor_system_id, 
			owner_id, owner_group, env_id, del_flag, create_time, last_update_time
		from t_meta_mysql_cluster_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata MYSQLClusterRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := er.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, fmt.Errorf("metadata MYSQLClusterInfo.GetByID(): data does not exists, id: %s", id)
	case 1:
		mysqlClusterInfo := NewEmptyMYSQLClusterInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return mysqlClusterInfo, nil
	default:
		return nil, fmt.Errorf("metadata MYSQLClusterInfo.GetByID(): duplicate key exists, id: %s", id)
	}
}

// GetID checks identity of given entity from the middleware
func (er *MYSQLClusterRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_mysql_cluster_info where del_flag = 0 and cluster_name = ?;`
	log.Debugf("metadata MYSQLClusterRepo.GetID() select sql: %s", sql)
	result, err := er.Execute(sql, entity.(*MYSQLClusterInfo).ClusterName)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (er *MYSQLClusterRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `
		insert into t_meta_mysql_cluster_info(cluster_name,middleware_cluster_id,
			 monitor_system_id, owner_id, owner_group, env_id) 
		values(?,?,?,?,?,?);`
	log.Debugf("metadata MYSQLClusterRepo.Create() insert sql: %s", sql)
	// execute
	_, err := er.Execute(sql,
		entity.(*MYSQLClusterInfo).ClusterName,
		entity.(*MYSQLClusterInfo).MiddlewareClusterID,
		entity.(*MYSQLClusterInfo).MonitorSystemID,
		entity.(*MYSQLClusterInfo).OwnerID,
		entity.(*MYSQLClusterInfo).OwnerGroup,
		entity.(*MYSQLClusterInfo).EnvID)
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
func (er *MYSQLClusterRepo) Update(entity dependency.Entity) error {
	sql := `
		update t_meta_mysql_cluster_info set cluster_name = ?, middleware_cluster_id = ?, 
			monitor_system_id = ?, owner_id = ?, owner_group = ?, 
			env_id = ?, del_flag = ? 
		where id = ?;`
	log.Debugf("metadata MYSQLClusterRepo.Update() update sql: %s", sql)
	mysqlClusterInfo := entity.(*MYSQLClusterInfo)
	_, err := er.Execute(sql,
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
func (er *MYSQLClusterRepo) Delete(id string) error {
	sql := `update t_meta_mysql_cluster_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata MYSQLClusterRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = er.Execute(sql, idInt)

	return err
}
