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

var _ dependency.Repository = (*MSRepo)(nil)

type MSRepo struct {
	Database middleware.Pool
}

// NewMSRepo returns *MSRepo with given middleware.Pool
func NewMSRepo(db middleware.Pool) *MSRepo {
	return &MSRepo{db}
}

// NewMSRepo returns *MSRepo with global mysql pool
func NewMSRepoWithGlobal() *MSRepo {
	return NewMSRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (er *MSRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := er.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata MSRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

func (er *MSRepo) Transaction() (middleware.Transaction, error) {
	return er.Database.Transaction()
}

// GetAll returns all available entities
func (er *MSRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, system_name, system_type, host_ip, port_num, port_num_slow, base_url, del_flag, create_time, last_update_time
		from t_meta_monitor_system_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata MSRepo.GetAll() sql: \n%s", sql)

	result, err := er.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*MSInfo
	mSInfoList := make([]*MSInfo, result.RowNumber())
	for i := range mSInfoList {
		mSInfoList[i] = NewEmptyMSInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(mSInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []dependency.Entity
	entityList := make([]dependency.Entity, result.RowNumber())
	for i := range entityList {
		entityList[i] = mSInfoList[i]
	}

	return entityList, nil
}

func (er *MSRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, system_name, system_type, host_ip, port_num, port_num_slow, base_url, del_flag, create_time, last_update_time
		from t_meta_monitor_system_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata MSRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := er.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata MSInfo.GetByID(): data does not exists, id: %s", id))
	case 1:
		dbInfo := NewEmptyMSInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(dbInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return dbInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata MSInfo.GetByID(): duplicate key exists, id: %s", id))
	}
}

// GetID checks identity of given entity from the middleware
func (er *MSRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_monitor_system_info where del_flag = 0 and system_name = ?;`
	log.Debugf("metadata MSRepo.GetID() select sql: %s", sql)
	result, err := er.Execute(sql, entity.(*MSInfo).MSName)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (er *MSRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `insert into t_meta_monitor_system_info(system_name, system_type, host_ip, port_num, port_num_slow, base_url) values(?,?,?,?,?,?);`
	log.Debugf("metadata MSRepo.Create() insert sql: %s", sql)
	// execute
	_, err := er.Execute(sql, entity.(*MSInfo).MSName, entity.(*MSInfo).SystemType, entity.(*MSInfo).HostIp, entity.(*MSInfo).PortNum, entity.(*MSInfo).PortNumSlow, entity.(*MSInfo).BaseUrl)
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
func (er *MSRepo) Update(entity dependency.Entity) error {
	sql := `update t_meta_monitor_system_info set system_name = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata MSRepo.Update() update sql: %s", sql)
	mSInfo := entity.(*MSInfo)
	_, err := er.Execute(sql, mSInfo.MSName, mSInfo.DelFlag, mSInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (er *MSRepo) Delete(id string) error {
	sql := `update t_meta_monitor_system_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata MSRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = er.Execute(sql, idInt)

	return err
}
