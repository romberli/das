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
	testInitClusterID          = 1
	testInitClusterName        = "cluster_name_init"
	testTransactionClusterID   = 2
	testTransactionClusterName = "cluster_name_need_rollback"
	testInsertClusterName      = "cluster_name_insert"
	testUpdateClusterName      = "cluster_name_update"
)

var mysqlClusterRepo = initMySQLClusterRepo()

func initMySQLClusterRepo() *MySQLClusterRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMySQLClusterRepo() failed", err))
		return nil
	}

	return NewMySQLClusterRepo(pool)
}

func createMySQLCluster() (dependency.Entity, error) {
	mysqlClusterInfo := NewMySQLClusterInfoWithDefault(
		defaultMySQLClusterInfoClusterName,
		defaultMySQLClusterInfoEnvID)
	entity, err := mysqlClusterRepo.Create(mysqlClusterInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteMySQLClusterByID(id string) error {
	sql := `delete from t_meta_mysql_cluster_info where id = ?`
	_, err := mysqlClusterRepo.Execute(sql, id)
	return err
}

func TestMySQLClusterRepoAll(t *testing.T) {
	TestMySQLClusterRepo_Execute(t)
	TestMySQLClusterRepo_Create(t)
	TestMySQLClusterRepo_GetAll(t)
	TestMySQLClusterRepo_GetByID(t)
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
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestMySQLClusterRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_cluster_info(
		id, cluster_name, middleware_cluster_id, monitor_system_id, 
		owner_id, owner_group, env_id) 
	values(?,?,?,?,?,?,?);`

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
		defaultMySQLClusterInfoOwnerGroup,
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
		mysqlClusterName, err := entity.Get(clusterNameStruct)
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
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
		owner_id, owner_group, env_id) 
	values(?,?,?,?,?,?,?);`

	// init data avoid empty set
	_, err := mysqlClusterRepo.Execute(sql,
		testInitClusterID,
		testInitClusterName,
		defaultMySQLClusterInfoMiddlewareClusterID,
		defaultMySQLClusterInfoMonitorSystemID,
		defaultMySQLClusterInfoOwnerID,
		defaultMySQLClusterInfoOwnerGroup,
		defaultMySQLClusterInfoEnvID)
	// asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))

	entities, err := mysqlClusterRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	mysqlClusterName, err := entities[0].Get("ClusterName")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(testInitClusterName, mysqlClusterName.(string), "test GetAll() failed")
}

func TestMySQLClusterRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := mysqlClusterRepo.GetByID("1")
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	mysqlClusterName, err := entity.Get(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testInitClusterName, mysqlClusterName.(string), "test GetByID() failed")
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
	mysqlClusterName, err := entity.Get(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
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
