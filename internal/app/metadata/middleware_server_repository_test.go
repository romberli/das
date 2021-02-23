package metadata

import (
	"github.com/romberli/das/internal/dependency"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	newMiddlewareServerName    = "test"
	onlineMiddlewareServerName = "online"
)

var middlewareServerRepo = initMiddlewareServerRepo()

func initMiddlewareServerRepo() *MiddlewareServerRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMiddlewareServerRepo() failed", err))
		return nil
	}

	return NewMiddlewareServerRepo(pool)
}

func createMiddlewareServer() (dependency.Entity, error) {
	middlewareServerInfo := NewMiddlewareServerInfoWithDefault(
		defaultMiddlewareServerInfoClusterID,
		defaultMiddlewareServerInfoServerName,
		defaultMiddlewareServerInfoMiddlewareRole,
		defaultMiddlewareServerInfoSHostIP,
		defaultMiddlewareServerInfoPortNum,
	)
	entity, err := middlewareServerRepo.Create(middlewareServerInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteMiddlewareServerByID(id string) error {
	sql := `delete from t_meta_middleware_server_info where id = ?`
	_, err := middlewareServerRepo.Execute(sql, id)
	return err
}

func TestMiddlewareServerRepoAll(t *testing.T) {
	TestMiddlewareServerRepo_Execute(t)
	TestMiddlewareServerRepo_GetAll(t)
	TestMiddlewareServerRepo_GetByID(t)
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
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestMiddlewareServerRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_middleware_server_info(cluster_id_middleware, server_name, middleware_role, host_ip, port_num) values(?, ?, ?, ?, ?);`
	tx, err := middlewareServerRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultMiddlewareServerInfoClusterID, defaultMiddlewareServerInfoServerName, defaultMiddlewareServerInfoMiddlewareRole, defaultMiddlewareServerInfoSHostIP, defaultMiddlewareServerInfoPortNum)
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
		middlewareServerName, err := entity.Get(middlewareServerNameStruct)
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
	middlewareServerName, err := entities[0].Get("ServerName")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineMiddlewareServerName, middlewareServerName.(string), "test GetAll() failed")
}

func TestMiddlewareServerRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := middlewareServerRepo.GetByID("1")
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	middlewareServerName, err := entity.Get(middlewareServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineMiddlewareServerName, middlewareServerName.(string), "test GetByID() failed")
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
	middlewareServerName, err := entity.Get(middlewareServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
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
