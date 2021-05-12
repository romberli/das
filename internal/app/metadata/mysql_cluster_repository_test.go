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
	testInitClusterID          = 1
	testInitClusterName        = "cluster_name_init"
	testTransactionClusterID   = 2
	testTransactionClusterName = "cluster_name_need_rollback"
	testInsertClusterName      = "cluster_name_insert"
	testUpdateClusterName      = "cluster_name_update"
)

var mysqlClusterRepo = initMySQLClusterRepo()

func initMySQLClusterRepo() *MySQLClusterRepo {
	pool, err := mysql.NewPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMySQLClusterRepo() failed", err))
		return nil
	}

	return NewMySQLClusterRepo(pool)
}

func createMySQLCluster() (metadata.MySQLCluster, error) {
	mysqlClusterInfo := NewMySQLClusterInfoWithDefault(
		defaultMySQLClusterInfoClusterName,
		defaultMySQLClusterInfoEnvID)
	entity, err := mysqlClusterRepo.Create(mysqlClusterInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteMySQLClusterByID(id int) error {
	sql := `delete from t_meta_mysql_cluster_info where id = ?`
	_, err := mysqlClusterRepo.Execute(sql, id)
	return err
}

func TestMySQLClusterRepoAll(t *testing.T) {
	TestMySQLClusterRepo_Execute(t)
	TestMySQLClusterRepo_Transaction(t)
	TestMySQLClusterRepo_Create(t)
	TestMySQLClusterRepo_GetAll(t)
	TestMySQLClusterRepo_GetByEnv(t)
	TestMySQLClusterRepo_GetByID(t)
	TestMySQLClusterRepo_GetByName(t)
	TestMySQLClusterRepo_GetID(t)
	TestMySQLClusterRepo_GetMySQLServerIDList(t)
	TestMySQLClusterRepo_Update(t)
	TestMySQLClusterRepo_Delete(t)
}

func TestMySQLClusterRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := mysqlClusterRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestMySQLClusterRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_cluster_info(
		id, cluster_name, middleware_cluster_id, monitor_system_id, 
		owner_id, env_id) 
	values(?,?,?,?,?,?);`

	tx, err := mysqlClusterRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql,
		testTransactionClusterID,
		testTransactionClusterName,
		defaultMySQLClusterInfoMiddlewareClusterID,
		defaultMySQLClusterInfoMonitorSystemID,
		defaultMySQLClusterInfoOwnerID,
		defaultMySQLClusterInfoEnvID)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select cluster_name from t_meta_mysql_cluster_info where cluster_name=?`
	result, err := tx.Execute(sql, testTransactionClusterName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	mysqlClusterName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if mysqlClusterName != testTransactionClusterName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := mysqlClusterRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		mysqlClusterName := entity.GetClusterName()
		if mysqlClusterName == testTransactionClusterName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMySQLClusterRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_cluster_info(
		id, cluster_name, middleware_cluster_id, monitor_system_id, 
		owner_id, env_id) 
	values(?,?,?,?,?,?);`

	// init data avoid empty set
	_, err := mysqlClusterRepo.Execute(sql,
		testInitClusterID,
		testInitClusterName,
		defaultMySQLClusterInfoMiddlewareClusterID,
		defaultMySQLClusterInfoMonitorSystemID,
		defaultMySQLClusterInfoOwnerID,
		defaultMySQLClusterInfoEnvID)
	// asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))

	entities, err := mysqlClusterRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	mysqlClusterName := entities[0].GetClusterName()
	asst.Equal(testInitClusterName, mysqlClusterName, "test GetAll() failed")
}

func TestMySQLClusterRepo_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	entities, err := mysqlClusterRepo.GetByEnv(testInitClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))

	for _, entity := range entities {
		clusterName := entity.GetClusterName()
		asst.Equal(testInitClusterName, clusterName, "test GetByEnv() failed")
	}
}

func TestMySQLClusterRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := mysqlClusterRepo.GetByID(testInitClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	mysqlClusterName := entity.GetClusterName()
	asst.Equal(testInitClusterName, mysqlClusterName, "test GetByID() failed")
}

func TestMySQLClusterRepo_GetByName(t *testing.T) {
	asst := assert.New(t)

	entity, err := mysqlClusterRepo.GetByName(testInitClusterName)
	asst.Nil(err, common.CombineMessageWithError("test GetByName() failed", err))
	clusterName := entity.GetClusterName()
	asst.Equal(testInitClusterName, clusterName, "test GetByName() failed")
}

func TestMySQLClusterRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	id, err := mysqlClusterRepo.GetID(testInitClusterName)
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.NotEqual(0, id, "test GetID() failed")
}

func TestMySQLClusterRepo_GetMySQLServerIDList(t *testing.T) {
	asst := assert.New(t)
	mysqlServerIDList, err := mysqlClusterRepo.GetMySQLServerIDList(testInitClusterID)
	// mysqlServerIDList, err := mysqlClusterRepo.GetMySQLServerIDList(97)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServerIDList() failed", err))
	for _, mysqlServerID := range mysqlServerIDList {
		mysqlServer, err := mysqlServerRepo.GetByID(mysqlServerID)
		asst.Nil(err, common.CombineMessageWithError("test GetMySQLServerIDList() failed", err))
		asst.Equal(mysqlServer.GetClusterID(), testInitClusterID, "test GetMySQLServerIDList() failed", err)
	}
}

func TestMySQLClusterRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMySQLClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMySQLClusterRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{clusterNameStruct: testUpdateClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = mysqlClusterRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = mysqlClusterRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	mysqlClusterName := entity.GetClusterName()
	asst.Equal(testUpdateClusterName, mysqlClusterName, "test Update() failed")
	// delete
	err = deleteMySQLClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMySQLClusterRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMySQLClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
