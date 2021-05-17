package metadata

import (
	"testing"

	"github.com/romberli/das/internal/dependency/metadata"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	middlewareServerAddr   = "localhost:3306"
	middlewareServerDBName = "das"
	middlewareServerDBUser = "root"
	middlewareServerDBPass = "rootroot"

	newMiddlewareServerName    = "newTest"
	onlineMiddlewareServerName = "test001"
	onlineMiddlewareServerID   = 1
)

var middlewareServerRepo = initMiddlewareServerRepo()

func initMiddlewareServerRepo() *MiddlewareServerRepo {
	pool, err := mysql.NewPoolWithDefault(middlewareServerAddr, middlewareServerDBName, middlewareServerDBUser, middlewareServerDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMiddlewareServerRepo() failed", err))
		return nil
	}

	return NewMiddlewareServerRepo(pool)
}

func createMiddlewareServer() (metadata.MiddlewareServer, error) {
	middlewareServerInfo := NewMiddlewareServerInfoWithDefault(
		defaultMiddlewareServerInfoClusterID,
		defaultMiddlewareServerInfoServerName,
		defaultMiddlewareServerInfoMiddlewareRole,
		defaultMiddlewareServerInfoHostIP,
		defaultMiddlewareServerInfoPortNum,
	)
	entity, err := middlewareServerRepo.Create(middlewareServerInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteMiddlewareServerByID(id int) error {
	sql := `delete from t_meta_middleware_server_info where id = ?`
	_, err := middlewareServerRepo.Execute(sql, id)
	return err
}

func TestMiddlewareServerRepoAll(t *testing.T) {
	TestMiddlewareServerRepo_Execute(t)
	TestMiddlewareClusterRepo_Transaction(t)
	TestMiddlewareServerRepo_GetAll(t)
	TestMiddlewareServerRepo_GetByClusterID(t)
	TestMiddlewareServerRepo_GetByID(t)
	TestMiddlewareServerRepo_GetByHostInfo(t)
	TestMiddlewareServerRepo_GetID(t)
	TestMiddlewareServerRepo_Create(t)
	TestMiddlewareServerRepo_Update(t)
	TestMiddlewareServerRepo_Delete(t)
}
func TestMiddlewareServerRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := middlewareServerRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestMiddlewareServerRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_middleware_server_info(cluster_id, server_name, middleware_role, host_ip, port_num) values(?, ?, ?, ?, ?);`
	tx, err := middlewareServerRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultMiddlewareServerInfoClusterID, defaultMiddlewareServerInfoServerName, defaultMiddlewareServerInfoMiddlewareRole, defaultMiddlewareServerInfoHostIP, defaultMiddlewareServerInfoPortNum)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select server_name from t_meta_middleware_server_info where server_name=?`
	result, err := tx.Execute(sql, defaultMiddlewareServerInfoServerName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	middlewareServerName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if middlewareServerName != defaultMiddlewareServerInfoServerName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := middlewareServerRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		middlewareServerName := entity.GetServerName()
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if middlewareServerName == defaultMiddlewareServerInfoServerName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMiddlewareServerRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := middlewareServerRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	middlewareServerName := entities[0].GetServerName()
	asst.Equal(onlineMiddlewareServerName, middlewareServerName, "test GetAll() failed")
}

func TestMiddlewareServerRepo_GetByClusterID(t *testing.T) {
	asst := assert.New(t)

	entities, err := middlewareServerRepo.GetByClusterID(13)
	asst.Nil(err, common.CombineMessageWithError("test GetByClusterID failed", err))
	middlewareServerName := entities[0].GetServerName()
	asst.Equal(onlineMiddlewareServerName, middlewareServerName, "test GetByClusterID failed")
}

func TestMiddlewareServerRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := middlewareServerRepo.GetByID(1)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	middlewareServerName := entity.GetServerName()
	asst.Equal(onlineMiddlewareServerName, middlewareServerName, "test GetByID() failed")
}

func TestMiddlewareServerRepo_GetByHostInfo(t *testing.T) {
	asst := assert.New(t)

	entity, err := middlewareServerRepo.GetByHostInfo("1", 1)
	asst.Nil(err, common.CombineMessageWithError("test GetByHostInfo() failed", err))
	middlewareServerName := entity.GetServerName()
	asst.Equal(onlineMiddlewareServerName, middlewareServerName, "test GetByHostInfo() failed")
}

func TestMiddlewareServerRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	id, err := middlewareServerRepo.GetID("1", 1)
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.Equal(onlineMiddlewareServerID, id, "")
}

func TestMiddlewareServerRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMiddlewareServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMiddlewareServerRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{middlewareServerNameStruct: newMiddlewareServerName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = middlewareServerRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = middlewareServerRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	middlewareServerName := entity.GetServerName()

	asst.Equal(newMiddlewareServerName, middlewareServerName, "test Update() failed")
	// delete
	err = deleteMiddlewareServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMiddlewareServerRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMiddlewareServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
