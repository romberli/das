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
	testInitHostIP         = "host_ip_init"
	testInitPortNum        = 3306
	testInsertHostIP       = "host_ip_insert"
	testInsertPortNum      = 3307
	testTransactionHostIP  = "host_ip_need_rollback"
	testTransactionPortNum = 3308
	testUpdateHostIP       = "host_ip_update"
	testUpdatePortNum      = 3309
)

var mysqlServerRepo = initMYSQLServerRepo()

func initMYSQLServerRepo() *MYSQLServerRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMYSQLServerRepo() failed", err))
		return nil
	}

	return NewMYSQLServerRepo(pool)
}

func createMYSQLServer() (dependency.Entity, error) {
	mysqlServerInfo := NewMYSQLServerInfoWithDefault(
		defaultMYSQLServerInfoHostIP, defaultMYSQLServerInfoPortNum)
	entity, err := mysqlServerRepo.Create(mysqlServerInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteMYSQLServerByID(id string) error {
	sql := `delete from t_meta_mysql_server_info where id = ?`
	_, err := mysqlServerRepo.Execute(sql, id)
	return err
}

func TestMYSQLServerRepoAll(t *testing.T) {
	TestMYSQLServerRepo_Execute(t)
	TestMYSQLServerRepo_Create(t)
	TestMYSQLServerRepo_GetAll(t)
	TestMYSQLServerRepo_GetByID(t)
	TestMYSQLServerRepo_Update(t)
	TestMYSQLServerRepo_Delete(t)
}

func TestMYSQLServerRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := mysqlServerRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestMYSQLServerRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_server_info(
		cluster_id, host_ip, port_num, deployment_type, version) 
	values(?,?,?,?,?);`

	tx, err := mysqlServerRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql,
		defaultMYSQLServerInfoClusterID,
		testTransactionHostIP,
		testTransactionPortNum,
		defaultMYSQLServerInfoDeploymentType,
		defaultMYSQLServerInfoVersion)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select host_ip, port_num from t_meta_mysql_server_info where host_ip=? and port_num=?`
	result, err := tx.Execute(sql, testTransactionHostIP, testTransactionPortNum)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	hostIP, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if hostIP != testTransactionHostIP {
		asst.Fail("test Transaction() failed")
	}
	portNum, err := result.GetInt(0, 1)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if portNum != testTransactionPortNum {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := mysqlServerRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		hostIP, err := entity.Get(hostIPStruct)
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		portNum, err := entity.Get(portNumStruct)
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if hostIP == testTransactionHostIP && portNum == testTransactionPortNum {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMYSQLServerRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_server_info(
		id, cluster_id, host_ip, port_num, deployment_type, version) 
	values(1,?,?,?,?,?);`

	// init data avoid empty set
	_, err := mysqlServerRepo.Execute(sql,
		defaultMYSQLServerInfoClusterID,
		testInitHostIP,
		testInitPortNum,
		defaultMYSQLServerInfoDeploymentType,
		defaultMYSQLServerInfoVersion)
	// asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))

	entities, err := mysqlServerRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	hostIP, err := entities[0].Get("HostIP")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(testInitHostIP, hostIP.(string), "test GetAll() failed")
	portNum, err := entities[0].Get("PortNum")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(testInitPortNum, portNum.(int), "test GetAll() failed")
}

func TestMYSQLServerRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := mysqlServerRepo.GetByID("1")
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	hostIP, err := entity.Get(hostIPStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testInitHostIP, hostIP.(string), "test GetByID() failed")
	portNum, err := entity.Get(portNumStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testInitPortNum, portNum.(int), "test GetByID() failed")
}

func TestMYSQLServerRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMYSQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMYSQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMYSQLServerRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMYSQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{
		hostIPStruct:  testUpdateHostIP,
		portNumStruct: testUpdatePortNum})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = mysqlServerRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = mysqlServerRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	hostIP, err := entity.Get(hostIPStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testUpdateHostIP, hostIP, "test Update() failed")
	portNum, err := entity.Get(portNumStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testUpdatePortNum, portNum, "test Update() failed")
	// delete
	err = deleteMYSQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMYSQLServerRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMYSQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMYSQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
