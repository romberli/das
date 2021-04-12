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

var _ metadata.MiddlewareServerRepo = (*MiddlewareServerRepo)(nil)

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
func (msr *MiddlewareServerRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := msr.Database.Get()
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

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (msr *MiddlewareServerRepo) Transaction() (middleware.Transaction, error) {
	return msr.Database.Transaction()
}

// GetAll returns all available entities
func (msr *MiddlewareServerRepo) GetAll() ([]metadata.MiddlewareServer, error) {
	sql := `
      select id, cluster_id, server_name, middleware_role, host_ip, port_num, del_flag, create_time, last_update_time
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
	entityList := make([]metadata.MiddlewareServer, result.RowNumber())
	for i := range entityList {
		entityList[i] = middlewareServerInfoList[i]
	}

	return entityList, nil
}

// GetByClusterID gets middleware servers with given cluster id
func (msr *MiddlewareServerRepo) GetByClusterID(clusterID int) ([]metadata.MiddlewareServer, error) {
	sql := `
		select id, cluster_id, server_name, middleware_role, host_ip, port_num, del_flag, create_time, last_update_time
		from t_meta_middleware_server_info
		where del_flag = 0
		and cluster_id = ?
		order by id;
	`
	log.Debugf("metadata MiddlewareServerRepo.GetByClusterID() sql: \n%s", sql, clusterID)

	result, err := msr.Execute(sql, clusterID)
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
	// init []metadata.MiddlewareServer
	middlewareServerList := make([]metadata.MiddlewareServer, result.RowNumber())
	for i := range middlewareServerList {
		middlewareServerList[i] = middlewareServerInfoList[i]
	}

	return middlewareServerList, nil
}

func (msr *MiddlewareServerRepo) GetByID(id int) (metadata.MiddlewareServer, error) {
	sql := `
      select id, cluster_id, server_name, middleware_role, host_ip, port_num, del_flag, create_time, last_update_time
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
		return nil, errors.New(fmt.Sprintf("metadata MiddlewareServerInfo.GetByID(): data does not exists, id: %d", id))
	case 1:
		middlewareServerInfo := NewEmptyMiddlewareServerInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(middlewareServerInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return middlewareServerInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata MiddlewareServerInfo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetByHostInfo gets a middleware server with given host ip and port number
func (msr *MiddlewareServerRepo) GetByHostInfo(hostIP string, portNum int) (metadata.MiddlewareServer, error) {
	sql := `
      select id from t_meta_middleware_server_info
      where del_flag = 0
      and host_ip = ? and port_num = ?;
   `
	log.Debugf("metadata MiddlewareServerRepo.GetByHostInfo() sql: %s", sql)

	result, err := msr.Execute(sql, hostIP, portNum)
	if err != nil {
		return nil, err
	}
	id, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	if err != nil {
		return nil, err
	}
	return msr.GetByID(id)
}

// GetID checks identity of given entity from the middleware
func (msr *MiddlewareServerRepo) GetID(hostIP string, portNum int) (int, error) {
	sql := `select id from t_meta_middleware_server_info where del_flag = 0 and host_ip = ? and port_num = ?;`
	log.Debugf("metadata MiddlewareServerRepo.GetID() select sql: %s", sql)
	result, err := msr.Execute(sql, hostIP, portNum)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (msr *MiddlewareServerRepo) Create(middlewareServer metadata.MiddlewareServer) (metadata.MiddlewareServer, error) {
	sql := `insert into t_meta_middleware_server_info(cluster_id, server_name, middleware_role, host_ip, port_num) values(?, ?, ?, ?, ?);`
	log.Debugf("metadata MiddlewareServerRepo.Create() insert sql: %s", sql)
	// execute
	_, err := msr.Execute(sql,
		middlewareServer.(*MiddlewareServerInfo).ClusterID,
		middlewareServer.(*MiddlewareServerInfo).ServerName,
		middlewareServer.(*MiddlewareServerInfo).MiddlewareRole,
		middlewareServer.(*MiddlewareServerInfo).HostIP,
		middlewareServer.(*MiddlewareServerInfo).PortNum,
	)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := msr.GetID(middlewareServer.GetHostIP(), middlewareServer.GetPortNum())
	if err != nil {
		return nil, err
	}
	// get entity
	return msr.GetByID(id)
}

// Update updates data with given entity in the middleware
func (msr *MiddlewareServerRepo) Update(middlewareServer metadata.MiddlewareServer) error {
	sql := `update t_meta_middleware_server_info set cluster_id = ?, server_name = ?, middleware_role = ?, host_ip = ?, port_num = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata MiddlewareServerRepo.Update() update sql: %s", sql)
	_, err := msr.Execute(sql,
		middlewareServer.GetClusterID(),
		middlewareServer.GetServerName(),
		middlewareServer.GetMiddlewareRole(),
		middlewareServer.GetHostIP(),
		middlewareServer.GetPortNum(),
		middlewareServer.GetDelFlag(),
		middlewareServer.Identity(),
	)

	return err
}

// Delete deletes the middleware server in the middleware
func (msr *MiddlewareServerRepo) Delete(id int) error {
	tx, err := msr.Transaction()
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Close()
		if err != nil {
			log.Errorf("metadata MiddlewareServerRepo.Delete(): close database connection failed.\n%s", err.Error())
		}
	}()

	err = tx.Begin()
	if err != nil {
		return err
	}
	sql := `delete from t_meta_middleware_server_info where id = ?;`
	log.Debugf("metadata MiddlewareServerRepo.Delete() update sql: %s", sql)
	_, err = msr.Execute(sql, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
