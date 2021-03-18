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
	newDBName    = "newTest"
	onlineDBName = "test"
)

var dbRepo = initDBRepo()

func initDBRepo() *DBRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initDBRepo() failed", err))
		return nil
	}

	return NewDBRepo(pool)
}

func createDB() (dependency.Entity, error) {
	dbInfo := NewDBInfoWithDefault(defaultDBInfoDBName, defaultDBInfoClusterID, defaultDBInfoClusterType, defaultDBInfoEnvID)
	entity, err := dbRepo.Create(dbInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteDBByID(id string) error {
	sql := `delete from t_meta_db_info where id = ?`
	_, err := dbRepo.Execute(sql, id)
	return err
}

func TestDBRepoAll(t *testing.T) {
	TestDBRepo_Execute(t)
	TestDBRepo_GetAll(t)
	TestDBRepo_GetByID(t)
	TestDBRepo_Create(t)
	TestDBRepo_Update(t)
	TestDBRepo_Delete(t)
}

func TestDBRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := dbRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestDBRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_db_info(db_name, cluster_id, cluster_type, owner_id, owner_group, env_id) values(?, ?, ?, ?, ?, ?);`
	tx, err := dbRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultDBInfoDBName, defaultDBInfoClusterID, defaultDBInfoClusterType, defaultDBInfoOwnerID, defaultDBInfoOwnerGroup, defaultDBInfoEnvID)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select db_name from t_meta_db_info where db_name = ?`
	result, err := tx.Execute(sql, defaultDBInfoDBName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	dbName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if dbName != defaultDBInfoDBName {
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
		if dbName == defaultDBInfoDBName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestDBRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := dbRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	dbName, err := entities[0].Get("DBName")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineDBName, dbName.(string), "test GetAll() failed")
}

func TestDBRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := dbRepo.GetByID("1")
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	dbName, err := entity.Get(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineDBName, dbName.(string), "test GetByID() failed")
}

func TestDBRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestDBRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{dbNameStruct: newDBName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = dbRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = dbRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	dbName, err := entity.Get(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newDBName, dbName, "test Update() failed")
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestDBRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
