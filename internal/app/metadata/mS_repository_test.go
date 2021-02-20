package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"

	"github.com/romberli/das/internal/dependency"
)

const (
	defaultMSInfoMSName  = "ms"
	defaultMSInfoHostIp  = "0.0.0.0"
	defaultMSInfoPortNum = "3306"
	defaultMSInfoBaseUrl = "http://127.0.0.1/prometheus/api/v1/"
	newMSName            = "newMS"
	onlineMSName         = "pmm"
)

var mSRepo = initMSRepo()

func initMSRepo() *MSRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMSRepo() failed", err))
		return nil
	}

	return NewMSRepo(pool)
}

func createMS() (dependency.Entity, error) {
	mSInfo := NewMSInfoWithDefault(defaultMSInfoMSName, defaultMSInfoHostIp, defaultMSInfoBaseUrl, defaultMSInfoPortNum)
	entity, err := mSRepo.Create(mSInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteMSByID(id string) error {
	sql := `delete from t_meta_monitor_system_info where id = ?`
	_, err := mSRepo.Execute(sql, id)
	return err
}

func TestMSRepoAll(t *testing.T) {
	TestMSRepo_Execute(t)
	TestMSRepo_GetAll(t)
	TestMSRepo_GetByID(t)
	TestMSRepo_Create(t)
	TestMSRepo_Update(t)
	TestMSRepo_Delete(t)
}

func TestMSRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := mSRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestMSRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_monitor_system_info(system_name, host_ip, port_num, base_url) values(?,?,?,?);`
	tx, err := mSRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultMSInfoMSName, defaultMSInfoHostIp, defaultMSInfoPortNum, defaultMSInfoBaseUrl)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select system_name from t_meta_monitor_system_info where system_name=?`
	result, err := tx.Execute(sql, defaultMSInfoMSName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	mSName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if mSName != defaultMSInfoMSName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := mSRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		mSName, err := entity.Get(mSNameStruct)
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if mSName == defaultMSInfoMSName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMSRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := mSRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	mSName, err := entities[0].Get("MSName")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineMSName, mSName.(string), "test GetAll() failed")
}

func TestMSRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := mSRepo.GetByID("1")
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	mSName, err := entity.Get(mSNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineMSName, mSName.(string), "test GetByID() failed")
}

func TestMSRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMS()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMSByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMSRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMS()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{mSNameStruct: newMSName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = mSRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = mSRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	mSName, err := entity.Get(mSNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newMSName, mSName, "test Update() failed")
	// delete
	err = deleteMSByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMSRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMS()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMSByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
