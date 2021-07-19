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
	// modify these connection information
	dbAddr    = "127.0.0.1:3306"
	dbDBName  = "das"
	dbDBUser  = "root"
	dbDBPass  = "rootroot"
	newDBName = "newTest"
)

var dbRepo = initDBRepo()

func initDBRepo() *DBRepo {
	pool, err := mysql.NewPoolWithDefault(dbAddr, dbDBName, dbDBUser, dbDBPass)
	log.Infof("pool: %v, error: %v", pool, err)
	if err != nil {
		log.Error(common.CombineMessageWithError("initDBRepo() failed", err))
		return nil
	}

	return NewDBRepo(pool)
}

func createDB() (metadata.DB, error) {
	dbInfo := NewDBInfoWithDefault(defaultDBInfoDBName, defaultDBInfoClusterID, defaultDBInfoClusterType, defaultDBInfoEnvID)
	entity, err := dbRepo.Create(dbInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteDBByID(id int) error {
	sql := `delete from t_meta_db_info where id = ?`
	_, err := dbRepo.Execute(sql, id)
	return err
}

func TestDBRepoAll(t *testing.T) {
	TestDBRepo_Execute(t)
	TestDBRepo_GetAll(t)
	TestDBRepo_GetByEnv(t)
	TestDBRepo_GetByID(t)
	TestDBRepo_GetByNameAndClusterInfo(t)
	TestDBRepo_GetAppIDList(t)
	TestDBRepo_Create(t)
	TestDBRepo_Update(t)
	TestDBRepo_Delete(t)
	TestDBRepo_AddDBApp(t)
	TestDBRepo_DeleteDBApp(t)
}

func TestDBRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := dbRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestDBRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_db_info(db_name, cluster_id, cluster_type, owner_id, env_id) values(?, ?, ?, ?, ?);`
	tx, err := dbRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, newDBName, defaultDBInfoClusterID, defaultDBInfoClusterType, defaultDBInfoOwnerID, defaultDBInfoEnvID)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select db_name from t_meta_db_info where db_name = ?`
	result, err := tx.Execute(sql, newDBName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	dbName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if dbName != newDBName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := dbRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		dbName := entity.GetDBName()
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if dbName == newDBName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestDBRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	entities, err := dbRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	dbName := entities[0].GetDBName()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(defaultDBInfoDBName, dbName, "test GetAll() failed")
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
}

func TestDBRepo_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
	entities, err := dbRepo.GetByEnv(defaultDBInfoEnvID)
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
	asst.Equal(defaultDBInfoEnvID, entities[0].GetEnvID(), common.CombineMessageWithError("test GetByEnv() failed", err))
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
}

func TestDBRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	db, err := dbRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	dbName := db.GetDBName()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(defaultDBInfoDBName, dbName, "test GetByID() failed")
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
}

func TestDBRepo_GetByNameAndClusterInfo(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	db, err := dbRepo.GetByNameAndClusterInfo(entity.GetDBName(), entity.GetClusterID(), entity.GetClusterType())
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	dbName := db.GetDBName()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(defaultDBInfoDBName, dbName, "test GetByID() failed")
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
}

func TestDBRepo_GetAppIDList(t *testing.T) {
	asst := assert.New(t)

	appIDList, err := dbRepo.GetAppIDList(1)
	asst.Nil(err, common.CombineMessageWithError("test GetAppIDList() failed", err))
	asst.Equal(2, len(appIDList), "test GetAppIDList() failed")
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
	err = entity.Set(map[string]interface{}{dbDBNameStruct: newDBName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = dbRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = dbRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	dbName := entity.GetDBName()
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

func TestDBRepo_AddDBApp(t *testing.T) {
	asst := assert.New(t)

	err := dbRepo.AddApp(1, 3)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	appIDList, err := dbRepo.GetAppIDList(1)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	asst.Equal(3, len(appIDList), "test AddApp() failed")
}

func TestDBRepo_DeleteDBApp(t *testing.T) {
	asst := assert.New(t)

	err := dbRepo.DeleteApp(1, 3)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	appIDList, err := dbRepo.GetAppIDList(1)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	asst.Equal(2, len(appIDList), "test DeleteApp() failed")
}
