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
	middlewareClusterAddr   = "localhost:3306"
	middlewareClusterDBName = "das"
	middlewareClusterDBUser = "root"
	middlewareClusterDBPass = "rootroot"

	newMiddlewareClusterName    = "ttt"
	onlineMiddlewareClusterName = "test"
	onlineMiddlewareClusterID   = 8
)

var middlewareClusterRepo = initMiddlewareClusterRepo()

func initMiddlewareClusterRepo() *MiddlewareClusterRepo {
	pool, err := mysql.NewPoolWithDefault(middlewareClusterAddr, middlewareClusterDBName, middlewareClusterDBUser, middlewareClusterDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMiddlewareClusterRepo() failed", err))
		return nil
	}

	return NewMiddlewareClusterRepo(pool)
}

func createMiddlewareCluster() (metadata.MiddlewareCluster, error) {
	middlewareClusterInfo := NewMiddlewareClusterInfoWithDefault(
		defaultMiddlewareClusterInfoClusterName,
		defaultMiddlewareClusterInfoOwnerID,
		defaultMiddlewareClusterInfoEnvID,
	)
	entity, err := middlewareClusterRepo.Create(middlewareClusterInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteMiddlewareClusterByID(id int) error {
	sql := `delete from t_meta_middleware_cluster_info where id = ?`
	_, err := middlewareClusterRepo.Execute(sql, id)
	return err
}

func TestMiddlewareClusterRepoAll(t *testing.T) {
	TestMiddlewareClusterRepo_Execute(t)
	TestMiddlewareClusterRepo_Transaction(t)
	TestMiddlewareClusterRepo_GetAll(t)
	TestMiddlewareClusterRepo_GetByEnv(t)
	TestMiddlewareClusterRepo_GetByID(t)
	TestMiddlewareClusterRepo_GetByName(t)
	TestMiddlewareClusterRepo_GetID(t)
	TestMiddlewareClusterRepo_GetMiddlewareServerIDList(t)
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
	asst.Equal(1, r, "test Execute() failed")
}

func TestMiddlewareClusterRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_middleware_cluster_info(cluster_name, owner_id, env_id) values(?, ?, ?);`
	tx, err := middlewareClusterRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultMiddlewareClusterInfoClusterName, defaultMiddlewareClusterInfoOwnerID, defaultMiddlewareClusterInfoEnvID)
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
		middlewareClusterName := entity.GetClusterName()
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
	middlewareClusterName := entities[0].GetClusterName()
	asst.Equal(onlineMiddlewareClusterName, middlewareClusterName, "test GetAll() failed")
}

func TestMiddlewareClusterRepo_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	entities, err := middlewareClusterRepo.GetByEnv(1)
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
	middlewareClusterName := entities[0].GetClusterName()
	asst.Equal(onlineMiddlewareClusterName, middlewareClusterName, "test GetAll() failed")
}

func TestMiddlewareClusterRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := middlewareClusterRepo.GetByID(8)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	middlewareClusterName := entity.GetClusterName()
	asst.Equal(onlineMiddlewareClusterName, middlewareClusterName, "test GetByID() failed")
}

func TestMiddlewareClusterRepo_GetByName(t *testing.T) {
	asst := assert.New(t)

	entity, err := middlewareClusterRepo.GetByName("test")
	asst.Nil(err, common.CombineMessageWithError("test GetByName() failed", err))
	middlewareClusterName := entity.GetClusterName()
	asst.Equal(onlineMiddlewareClusterName, middlewareClusterName, "test GetByID() failed")
}

func TestMiddlewareClusterRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	id, err := middlewareClusterRepo.GetID("test", 1)
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.Equal(onlineMiddlewareClusterID, id, "test GetID() failed")
}

func TestMiddlewareClusterRepo_GetMiddlewareServerIDList(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareServerIDList() failed", err))
	middlewareServerIDList, err := middlewareClusterRepo.GetMiddlewareServerIDList(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareServerIDList() failed", err))
	asst.Equal(0, len(middlewareServerIDList), "test GetMiddlewareServerIDList() failed")
	// delete
	err = deleteMiddlewareClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareServerIDList() failed", err))
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
	middlewareClusterName := entity.GetClusterName()
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
