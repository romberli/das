package metadata

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"

	"github.com/romberli/log"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Repository = (*MonitorSystemRepo)(nil)

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

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
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

// Transaction implements dependency.Repository interface
func (msr *MonitorSystemRepo) Transaction() (middleware.Transaction, error) {
	return msr.Database.Transaction()
}

// GetAll returns all available entities
func (msr *MonitorSystemRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, system_name, system_type, host_ip, port_num, port_num_slow, base_url, del_flag, create_time, last_update_time
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
	// init []dependency.Entity
	entityList := make([]dependency.Entity, result.RowNumber())
	for i := range entityList {
		entityList[i] = monitorSystemInfoList[i]
	}

	return entityList, nil
}

func (msr *MonitorSystemRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, system_name, system_type, host_ip, port_num, port_num_slow, base_url, del_flag, create_time, last_update_time
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
		return nil, errors.New(fmt.Sprintf("metadata MonitorSystemInfo.GetByID(): data does not exists, id: %s", id))
	case 1:
		monitorSystemInfo := NewEmptyMonitorSystemInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(monitorSystemInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return monitorSystemInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata MonitorSystemInfo.GetByID(): duplicate key exists, id: %s", id))
	}
}

// GetID checks identity of given entity from the middleware
func (msr *MonitorSystemRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_monitor_system_info where del_flag = 0 and host_ip = ? and port_num = ?;`
	log.Debugf("metadata MonitorSystemRepo.GetID() select sql: %s", sql)
	result, err := msr.Execute(sql, entity.(*MonitorSystemInfo).MonitorSystemHostIP, entity.(*MonitorSystemInfo).MonitorSystemPortNum)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (msr *MonitorSystemRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `insert into t_meta_monitor_system_info(system_name, system_type, host_ip, port_num, port_num_slow, base_url) values(?, ?, ?, ?, ?, ?);`
	log.Debugf("metadata MonitorSystemRepo.Create() insert sql: %s", sql)
	// execute
	_, err := msr.Execute(sql, entity.(*MonitorSystemInfo).MonitorSystemName, entity.(*MonitorSystemInfo).MonitorSystemType, entity.(*MonitorSystemInfo).MonitorSystemHostIP, entity.(*MonitorSystemInfo).MonitorSystemPortNum, entity.(*MonitorSystemInfo).MonitorSystemPortNumSlow, entity.(*MonitorSystemInfo).BaseUrl)
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
func (msr *MonitorSystemRepo) Update(entity dependency.Entity) error {
	sql := `update t_meta_monitor_system_info set system_name = ?, system_type = ?, host_ip = ?, port_num = ?, port_num_slow = ?, base_url = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata MonitorSystemRepo.Update() update sql: %s", sql)
	monitorSystemInfo := entity.(*MonitorSystemInfo)
	_, err := msr.Execute(sql, monitorSystemInfo.MonitorSystemName, monitorSystemInfo.MonitorSystemType, monitorSystemInfo.MonitorSystemHostIP, monitorSystemInfo.MonitorSystemPortNum, monitorSystemInfo.MonitorSystemPortNumSlow, monitorSystemInfo.BaseUrl, monitorSystemInfo.DelFlag, monitorSystemInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (msr *MonitorSystemRepo) Delete(id string) error {
	sql := `update t_meta_monitor_system_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata MonitorSystemRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = msr.Execute(sql, idInt)

	return err
}
