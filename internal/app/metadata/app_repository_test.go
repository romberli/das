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
	appAddr       = "127.0.0.1:3306"
	appDBName     = "das"
	appDBUser     = "root"
	appDBPass     = "rootroot"
	onlineAppName = "test2"
	newAppName    = "testApp"
)

var appRepo = initAppRepo()

func initAppRepo() *AppRepo {
	pool, err := mysql.NewPoolWithDefault(appAddr, appDBName, appDBUser, appDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initAppRepo() failed", err))
		return nil
	}

	return NewAppRepo(pool)
}

func createApp() (metadata.App, error) {
	appSystemInfo := NewAppInfoWithDefault(
		defaultAppInfoAppName,
		defaultAppInfoLevel,
	)
	entity, err := appRepo.Create(appSystemInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteAppByID(id int) error {
	sql := `delete from t_meta_app_system_info where id = ?`
	_, err := appRepo.Execute(sql, id)
	return err
}

func TestAppRepoAll(t *testing.T) {
	TestAppRepo_Execute(t)
	TestAppRepo_GetAll(t)
	TestAppRepo_GetByID(t)
	TestAppRepo_GetAppByName(t)
	TestAppRepo_GetDBIDList(t)
	TestAppRepo_Create(t)
	TestAppRepo_Update(t)
	TestAppRepo_Delete(t)
	TestAppRepo_AddAppDB(t)
	TestAppRepo_DeleteAppDB(t)

}

func TestAppRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := appRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestAppRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_app_system_info(system_name,level,owner_id,owner_group) values(?,?,?,?);`
	tx, err := appRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultAppInfoAppName, defaultAppInfoLevel, defaultAppInfoOwnerID, defaultAppInfoOwnerGroup)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select system_name from t_meta_app_system_info where system_name=?`
	result, err := tx.Execute(sql,
		defaultAppInfoAppName,
	)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	appSystemName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if appSystemName != defaultAppInfoAppName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := appRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		systemName := entity.GetAppName()
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if systemName == defaultAppInfoAppName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestAppRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := appRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	systemName := entities[0].GetAppName()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineAppName, systemName, "test GetAll() failed")
}

func TestAppRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := appRepo.GetByID(2)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	systemName := entity.GetAppName()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineAppName, systemName, "test GetByID() failed")
}

func TestAppRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createApp()

	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestAppRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{appAppNameStruct: newAppName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = appRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = appRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	systemName := entity.GetAppName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newAppName, systemName, "test Update() failed")
	// delete
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestAppRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestAppRepo_GetAppByName(t *testing.T) {
	asst := assert.New(t)

	_, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test GetAppByName() failed", err))
	entity, err := appRepo.GetAppByName(defaultAppInfoAppName)
	asst.Nil(err, common.CombineMessageWithError("test GetAppByName() failed", err))
	asst.Equal(defaultAppInfoAppName, entity.GetAppName(), common.CombineMessageWithError("test GetAppByName() failed", err))
	// delete
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetAppByName() failed", err))
}

func TestAppRepo_GetDBIDList(t *testing.T) {
	asst := assert.New(t)

	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test GetDBIDList() failed", err))
	dbIDList, err := appRepo.GetDBIDList(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetDBIDList() failed", err))
	asst.Equal(2, len(dbIDList), "test GetDBIDList() failed")
	// delete
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetAppByName() failed", err))
}

func TestAppRepo_AddAppDB(t *testing.T) {
	asst := assert.New(t)

	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	err = appRepo.AddDB(entity.Identity(), 3)
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	dbIDList, err := appRepo.GetDBIDList(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	asst.Equal(3, len(dbIDList), "test AddDB() failed")
	// delete
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
}

func TestAppRepo_DeleteAppDB(t *testing.T) {
	asst := assert.New(t)

	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	err = appRepo.DeleteDB(entity.Identity(), 3)
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	dbIDList, err := appRepo.GetDBIDList(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	asst.Equal(2, len(dbIDList), "test DeleteDB() failed")
	// delete
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
}
