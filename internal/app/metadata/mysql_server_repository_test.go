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
	testInitServerID        = 1
	testInitHostIP          = "host_ip_init"
	testInitPortNum         = 3306
	testTransactionServerID = 2
	testTransactionHostIP   = "host_ip_need_rollback"
	testTransactionPortNum  = 3308
	testInsertHostIP        = "host_ip_insert"
	testUpdateHostIP        = "host_ip_update"
	testUpdatePortNum       = 3309
)

var mysqlServerRepo = initMySQLServerRepo()

func initMySQLServerRepo() *MySQLServerRepo {
	pool, err := mysql.NewPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMySQLServerRepo() failed", err))
		return nil
	}

	return NewMySQLServerRepo(pool)
}

func createMySQLServer() (metadata.MySQLServer, error) {
	mysqlServerInfo := NewMySQLServerInfoWithDefault(
		defaultMySQLServerInfoClusterID,
		defaultMySQLServerInfoServerName,
		defaultMySQLServerInfoHostIP,
		defaultMySQLServerInfoPortNum,
		defaultMySQLServerInfoDeploymentType)
	entity, err := mysqlServerRepo.Create(mysqlServerInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteMySQLServerByID(id int) error {
	sql := `delete from t_meta_mysql_server_info where id = ?`
	_, err := mysqlServerRepo.Execute(sql, id)
	return err
}

func TestMySQLServerRepoAll(t *testing.T) {
	TestMySQLServerRepo_Execute(t)
	TestMySQLServerRepo_Transaction(t)
	TestMySQLServerRepo_Create(t)
	TestMySQLServerRepo_GetAll(t)
	TestMySQLServerRepo_GetByClusterID(t)
	TestMySQLServerRepo_GetByID(t)
	TestMySQLServerRepo_GetByHostInfo(t)
	TestMySQLServerRepo_GetID(t)
	TestMySQLServerRepo_Update(t)
	TestMySQLServerRepo_Delete(t)
}

func TestMySQLServerRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := mysqlServerRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestMySQLServerRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_server_info(
		id, cluster_id, server_name, service_name, host_ip, port_num, deployment_type, version) 
	values(?,?,?,?,?,?,?);`

	tx, err := mysqlServerRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql,
		testTransactionServerID,
		defaultMySQLServerInfoClusterID,
		defaultMySQLServerInfoServerName,
		defaultMySQLServerInfoServiceName,
		testTransactionHostIP,
		testTransactionPortNum,
		defaultMySQLServerInfoDeploymentType,
		defaultMySQLServerInfoVersion)
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
		hostIP := entity.GetHostIP()
		portNum := entity.GetPortNum()
		if hostIP == testTransactionHostIP && portNum == testTransactionPortNum {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMySQLServerRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_server_info(
		id, cluster_id, server_name, service_name, host_ip, port_num, deployment_type, version) 
	values(?,?,?,?,?,?,?,?);`

	// init data avoid empty set
	_, err := mysqlServerRepo.Execute(sql,
		testInitServerID,
		defaultMySQLServerInfoClusterID,
		defaultMySQLServerInfoServerName,
		testInitHostIP,
		testInitPortNum,
		defaultMySQLServerInfoDeploymentType,
		defaultMySQLServerInfoVersion)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))

	entities, err := mysqlServerRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	hostIP := entities[0].GetHostIP()
	asst.Equal(testInitHostIP, hostIP, "test GetAll() failed")
	portNum := entities[0].GetPortNum()
	asst.Equal(testInitPortNum, portNum, "test GetAll() failed")
}

func TestMySQLServerRepo_GetByClusterID(t *testing.T) {
	asst := assert.New(t)

	entities, err := mysqlServerRepo.GetByClusterID(testInitClusterID)

	for _, entity := range entities {
		asst.Nil(err, common.CombineMessageWithError("test GetByClusterID() failed", err))
		hostIP := entity.GetHostIP()
		asst.Equal(testInitHostIP, hostIP, "test GetByClusterID() failed")
		portNum := entity.GetPortNum()
		asst.Equal(testInitPortNum, portNum, "test GetByClusterID() failed")
	}
}

func TestMySQLServerRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := mysqlServerRepo.GetByID(1)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	hostIP := entity.GetHostIP()
	asst.Equal(testInitHostIP, hostIP, "test GetByID() failed")
	portNum := entity.GetPortNum()
	asst.Equal(testInitPortNum, portNum, "test GetByID() failed")
}

func TestMySQLServerRepo_GetByHostInfo(t *testing.T) {
	asst := assert.New(t)

	entity, err := mysqlServerRepo.GetByHostInfo(testInitHostIP, testInitPortNum)
	asst.Nil(err, common.CombineMessageWithError("test GetByHostInfo() failed", err))
	hostIP := entity.GetHostIP()
	asst.Equal(testInitHostIP, hostIP, "test GetByHostInfo() failed")
	portNum := entity.GetPortNum()
	asst.Equal(testInitPortNum, portNum, "test GetByHostInfo() failed")
}

func TestMySQLServerRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	id, err := mysqlServerRepo.GetID(testInitHostIP, testInitPortNum)
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.NotEqual(0, id, "test GetID() failed")
}

func TestMySQLServerRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMySQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMySQLServerRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{
		hostIPStruct:  testUpdateHostIP,
		portNumStruct: testUpdatePortNum})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = mysqlServerRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = mysqlServerRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	hostIP := entity.GetHostIP()
	asst.Equal(testUpdateHostIP, hostIP, "test Update() failed")
	portNum := entity.GetPortNum()
	asst.Equal(testUpdatePortNum, portNum, "test Update() failed")
	// delete
	err = deleteMySQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMySQLServerRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMySQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
