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
	newMiddlewareClusterName    = "ttt"
	onlineMiddlewareClusterName = "test"
)

var middlewareClusterRepo = initMiddlewareClusterRepo()

func initMiddlewareClusterRepo() *MiddlewareClusterRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMiddlewareClusterRepo() failed", err))
		return nil
	}

	return NewMiddlewareClusterRepo(pool)
}

func createMiddlewareCluster() (dependency.Entity, error) {
	middlewareClusterInfo := NewMiddlewareClusterInfoWithDefault(defaultMiddlewareClusterInfoClusterName, defaultMiddlewareClusterInfoEnvID)
	entity, err := middlewareClusterRepo.Create(middlewareClusterInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteMiddlewareClusterByID(id string) error {
	sql := `delete from t_meta_middleware_cluster_info where id = ?`
	_, err := middlewareClusterRepo.Execute(sql, id)
	return err
}

func TestMiddlewareClusterRepoAll(t *testing.T) {
	TestMiddlewareClusterRepo_Execute(t)
	TestMiddlewareClusterRepo_GetAll(t)
	TestMiddlewareClusterRepo_GetByID(t)
	TestMiddlewareClusterRepo_Create(t)
	TestMiddlewareClusterRepo_Update(t)
	TestMiddlewareClusterRepo_Delete(t)
}
func TestMiddlewareClusterRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := middlewareClusterRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestMiddlewareClusterRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_middleware_cluster_info(cluster_name, env_id) values(?, ?);`
	tx, err := middlewareClusterRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultMiddlewareClusterInfoClusterName, defaultMiddlewareClusterInfoEnvID)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select cluster_name from t_meta_middleware_cluster_info where cluster_name=?`
	result, err := tx.Execute(sql, defaultMiddlewareClusterInfoClusterName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	middlewareClusterName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if middlewareClusterName != defaultMiddlewareClusterInfoClusterName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := middlewareClusterRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		middlewareClusterName, err := entity.Get(middlewareClusterNameStruct)
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if middlewareClusterName == defaultMiddlewareClusterInfoClusterName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMiddlewareClusterRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := middlewareClusterRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	middlewareClusterName, err := entities[0].Get("ClusterName")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineMiddlewareClusterName, middlewareClusterName.(string), "test GetAll() failed")
}

func TestMiddlewareClusterRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := middlewareClusterRepo.GetByID("3")
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	middlewareClusterName, err := entity.Get(middlewareClusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineMiddlewareClusterName, middlewareClusterName.(string), "test GetByID() failed")
}

func TestMiddlewareClusterRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMiddlewareClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMiddlewareClusterRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{middlewareClusterNameStruct: newMiddlewareClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = middlewareClusterRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = middlewareClusterRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	middlewareClusterName, err := entity.Get(middlewareClusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newMiddlewareClusterName, middlewareClusterName, "test Update() failed")
	// delete
	err = deleteMiddlewareClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMiddlewareClusterRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMiddlewareClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
