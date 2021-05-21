package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"

	"github.com/romberli/das/internal/dependency/metadata"
)

const (
	// modify these connection information
	monitorSystemAddr   = "127.0.0.1:3306"
	monitorSystemDBName = "das"
	monitorSystemDBUser = "root"
	monitorSystemDBPass = "rootroot"

	newMonitorSystemName = "newTest"
)

var monitorSystemRepo = initMonitorSystemRepo()

func initMonitorSystemRepo() *MonitorSystemRepo {
	pool, err := mysql.NewPoolWithDefault(monitorSystemAddr, monitorSystemDBName, monitorSystemDBUser, monitorSystemDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMonitorSystemRepo() failed", err))
		return nil
	}

	return NewMonitorSystemRepo(pool)
}

func createMonitorSystem() (metadata.MonitorSystem, error) {
	monitorSystemInfo := NewMonitorSystemInfoWithDefault(defaultMonitorSystemInfoSystemName,
		defaultMonitorSystemInfoSystemType, defaultMonitorSystemInfoHostIP, defaultMonitorSystemInfoPortNum,
		defaultMonitorSystemInfoPortNumSlow, defaultMonitorSystemInfoBaseUrl, defaultMonitorSystemInfoEnvID)
	entity, err := monitorSystemRepo.Create(monitorSystemInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteMonitorSystemByID(id int) error {
	sql := `delete from t_meta_monitor_system_info where id = ?`
	_, err := monitorSystemRepo.Execute(sql, id)
	return err
}

func TestMonitorSystemRepoAll(t *testing.T) {
	TestMonitorSystemRepo_Execute(t)
	TestMonitorSystemRepo_GetAll(t)
	TestMonitorSystemRepo_GetByEnv(t)
	TestMonitorSystemRepo_GetByID(t)
	TestMonitorSystemRepo_Create(t)
	TestMonitorSystemRepo_Update(t)
	TestMonitorSystemRepo_Delete(t)
}

func TestMonitorSystemRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := monitorSystemRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestMonitorSystemRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_monitor_system_info(system_name, system_type, host_ip, port_num, port_num_slow, base_url, env_id) values(?, ?, ?, ?, ?, ?, ?);`
	tx, err := monitorSystemRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultMonitorSystemInfoSystemName, defaultMonitorSystemInfoSystemType, defaultMonitorSystemInfoHostIP,
		defaultMonitorSystemInfoPortNum, defaultMonitorSystemInfoPortNumSlow, defaultMonitorSystemInfoBaseUrl, defaultMonitorSystemInfoEnvID)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select system_name from t_meta_monitor_system_info where system_name = ?`
	result, err := tx.Execute(sql, defaultMonitorSystemInfoSystemName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	monitorSystemName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if monitorSystemName != defaultMonitorSystemInfoSystemName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := monitorSystemRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		monitorSystemName := entity.GetSystemName()
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if monitorSystemName == defaultMonitorSystemInfoSystemName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMonitorSystemRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	entities, err := monitorSystemRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	monitorSystemName := entities[0].GetSystemName()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(defaultMonitorSystemInfoSystemName, monitorSystemName, "test GetAll() failed")
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
}

func TestMonitorSystemRepo_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
	entities, err := monitorSystemRepo.GetByEnv(defaultMonitorSystemInfoEnvID)
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
	asst.Equal(defaultMonitorSystemInfoEnvID, entities[0].GetEnvID(), common.CombineMessageWithError("test GetByEnv() failed", err))
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
}

func TestMonitorSystemRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	db, err := monitorSystemRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	monitorSystemName := db.GetSystemName()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(defaultMonitorSystemInfoSystemName, monitorSystemName, "test GetByID() failed")
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
}

func TestMonitorSystemRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMonitorSystemRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{monitorSystemNameStruct: newMonitorSystemName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = monitorSystemRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = monitorSystemRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	monitorSystemName := entity.GetSystemName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newMonitorSystemName, monitorSystemName, "test Update() failed")
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMonitorSystemRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
