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

var _ metadata.MonitorSystemRepo = (*MonitorSystemRepo)(nil)

type MonitorSystemRepo struct {
	Database middleware.Pool
}

// NewMonitorSystemRepo returns *MonitorSystemRepo with given middleware.Pool
func NewMonitorSystemRepo(db middleware.Pool) *MonitorSystemRepo {
	return &MonitorSystemRepo{db}
}

// NewMonitorSystemRepo returns *MonitorSystemRepo with global mysql pool
func NewMonitorSystemRepoWithGlobal() *MonitorSystemRepo {
	return NewMonitorSystemRepo(global.MySQLPool)
}

// Execute executes given command and placeholders on the middleware
func (msr *MonitorSystemRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := msr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata MonitorSystemRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (msr *MonitorSystemRepo) Transaction() (middleware.Transaction, error) {
	return msr.Database.Transaction()
}

// GetAll gets all monitor systems from the middleware
func (msr *MonitorSystemRepo) GetAll() ([]metadata.MonitorSystem, error) {
	sql := `
		select id, system_name, system_type, host_ip, port_num, port_num_slow, base_url, env_id, del_flag, create_time, last_update_time
		from t_meta_monitor_system_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata MonitorSystemRepo.GetAll() sql: \n%s", sql)

	result, err := msr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*MonitorSystemInfo
	monitorSystemInfoList := make([]*MonitorSystemInfo, result.RowNumber())
	for i := range monitorSystemInfoList {
		monitorSystemInfoList[i] = NewEmptyMonitorSystemInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(monitorSystemInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.MonitorSystem
	monitorSystemList := make([]metadata.MonitorSystem, result.RowNumber())
	for i := range monitorSystemList {
		monitorSystemList[i] = monitorSystemInfoList[i]
	}

	return monitorSystemList, nil
}

// GetByEnv gets monitor systems of given env id from the middleware
func (msr *MonitorSystemRepo) GetByEnv(envID int) ([]metadata.MonitorSystem, error) {
	sql := `
		select id, system_name, system_type, host_ip, port_num, port_num_slow, base_url, env_id, del_flag, create_time, last_update_time
		from t_meta_monitor_system_info
		where del_flag = 0
		and env_id = ?
		order by id;
	`
	log.Debugf("metadata MonitorSystemRepo.GetByEnv sql: \n%s", sql)

	result, err := msr.Execute(sql, envID)
	if err != nil {
		return nil, err
	}
	// init []*MonitorSystemInfo
	monitorSystemInfoList := make([]*MonitorSystemInfo, result.RowNumber())
	for i := range monitorSystemInfoList {
		monitorSystemInfoList[i] = NewEmptyMonitorSystemInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(monitorSystemInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.MonitorSystem
	monitorSystemList := make([]metadata.MonitorSystem, result.RowNumber())
	for i := range monitorSystemList {
		monitorSystemList[i] = monitorSystemInfoList[i]
	}

	return monitorSystemList, nil
}

// GetByID gets a monitor system by the identity from the middleware
func (msr *MonitorSystemRepo) GetByID(id int) (metadata.MonitorSystem, error) {
	sql := `
		select id, system_name, system_type, host_ip, port_num, port_num_slow, base_url, env_id, del_flag, create_time, last_update_time
		from t_meta_monitor_system_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata MonitorSystemRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := msr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata MonitorSystemInfo.GetByID(): data does not exists, id: %d", id))
	case 1:
		monitorSystemInfo := NewEmptyMonitorSystemInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(monitorSystemInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return monitorSystemInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata MonitorSystemInfo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetByID gets a monitor system by the identity from the middleware
func (msr *MonitorSystemRepo) GetByHostInfo(hostIP string, portNum int) (metadata.MonitorSystem, error) {
	sql := `select id from t_meta_monitor_system_info where del_flag = 0 and host_ip = ? and port_num = ?;`
	log.Debugf("metadata MonitorSystemRepo.GetByHostInfo() sql: \n%s\nplaceholders: %s", sql)

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

// GetID gets the identity with given host ip and port num from the middleware
func (msr *MonitorSystemRepo) GetID(monitorSystemHostIP string, monitorSystemPortNum int) (int, error) {
	sql := `select id from t_meta_monitor_system_info where del_flag = 0 and host_ip = ? and port_num = ?;`
	log.Debugf("metadata MonitorSystemRepo.GetID() select sql: %s", sql)
	result, err := msr.Execute(sql, monitorSystemHostIP, monitorSystemPortNum)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// Create creates a monitor system in the middleware
func (msr *MonitorSystemRepo) Create(monitorSystem metadata.MonitorSystem) (metadata.MonitorSystem, error) {
	sql := `insert into t_meta_monitor_system_info(system_name, system_type, host_ip, port_num, port_num_slow, base_url, env_id) values(?, ?, ?, ?, ?, ?, ?);`
	log.Debugf("metadata MonitorSystemRepo.Create() insert sql: %s", sql)
	// execute
	_, err := msr.Execute(sql, monitorSystem.GetSystemName(), monitorSystem.GetSystemType(), monitorSystem.GetHostIP(),
		monitorSystem.GetPortNum(), monitorSystem.GetPortNumSlow(), monitorSystem.GetBaseURL(), monitorSystem.GetEnvID())
	if err != nil {
		return nil, err
	}
	// get id
	id, err := msr.GetID(monitorSystem.GetHostIP(), monitorSystem.GetPortNum())
	if err != nil {
		return nil, err
	}
	// get entity
	return msr.GetByID(id)
}

// Update updates the monitor system in the middleware
func (msr *MonitorSystemRepo) Update(monitorSystem metadata.MonitorSystem) error {
	sql := `update t_meta_monitor_system_info set system_name = ?, system_type = ?, host_ip = ?, port_num = ?, port_num_slow = ?, base_url = ?, env_id = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata MonitorSystemRepo.Update() update sql: %s", sql)
	_, err := msr.Execute(sql, monitorSystem.GetSystemName(), monitorSystem.GetSystemType(), monitorSystem.GetHostIP(),
		monitorSystem.GetPortNum(), monitorSystem.GetPortNumSlow(), monitorSystem.GetBaseURL(), monitorSystem.GetEnvID(),
		monitorSystem.GetDelFlag(), monitorSystem.Identity())

	return err
}

// Delete deletes the monitor system in the middleware
func (msr *MonitorSystemRepo) Delete(id int) error {
	tx, err := msr.Transaction()
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Close()
		if err != nil {
			log.Errorf("metadata MonitorSystemRepo.Delete(): close database connection failed.\n%s", err.Error())
		}
	}()

	err = tx.Begin()
	if err != nil {
		return err
	}
	sql := `delete from t_meta_monitor_system_info where id = ?;`
	log.Debugf("metadata MonitorSystemRepo.Delete() delete sql(t_meta_monitor_system_info): %s", sql)
	_, err = msr.Execute(sql, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
