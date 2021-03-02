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
	newDbName    = "newDb"
	onlineDbName = "db1"
)

var dbRepo = initDbRepo()

func initDbRepo() *DbRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initEnvRepo() failed", err))
		return nil
	}

	return NewDbRepo(pool)
}

func createDb() (dependency.Entity, error) {
	dbInfo := NewDbInfoWithDefault(defaultDbInfoDbName, defaultDbInfoOwnerId, defaultDbInfoEnvId)
	entity, err := dbRepo.Create(dbInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteDbByID(id string) error {
	sql := `delete from t_meta_db_info where id = ?`
	_, err := dbRepo.Execute(sql, id)
	return err
}

func TestDbRepoAll(t *testing.T) {
	TestDbRepo_Execute(t)
	TestDbRepo_GetAll(t)
	TestDbRepo_GetByID(t)
	TestDbRepo_Create(t)
	TestDbRepo_Update(t)
	TestDbRepo_Delete(t)
}

func TestDbRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := dbRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestDbRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_db_info(db_name, owner_id, env_id) values(?,?,?);`
	tx, err := dbRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultDbInfoDbName, defaultDbInfoOwnerId, defaultDbInfoEnvId)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select db_name from t_meta_db_info where db_name=?`
	result, err := tx.Execute(sql, defaultDbInfoDbName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	dbName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if dbName != defaultDbInfoDbName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := dbRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		dbName, err := entity.Get(dbNameStruct)
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if dbName == defaultDbInfoDbName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestDbRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := dbRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	dbName, err := entities[0].Get("DbName")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineDbName, dbName.(string), "test GetAll() failed")
}

func TestDbRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := dbRepo.GetByID("1")
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	dbName, err := entity.Get(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineDbName, dbName.(string), "test GetByID() failed")
}

func TestDbRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDb()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteDbByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestDbRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDb()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{dbNameStruct: newDbName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = dbRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = dbRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	dbName, err := entity.Get(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newDbName, dbName, "test Update() failed")
	// delete
	err = deleteDbByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestDbRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDb()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteDbByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
