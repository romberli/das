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
	onlineAppSystemName = "1"
	newAppSystemName    = "dfname"
)

var appSystemRepo = initAppSystemRepo()

func initAppSystemRepo() *AppSystemRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initAppSystemRepo() failed", err))
		return nil
	}

	return NewAppSystemRepo(pool)
}

func createAppSystem() (dependency.Entity, error) {
	appSystemInfo := NewAppSystemInfoWithDefault(
		defaultAppSystemInfoAppSystemName,
		defaultAppSystemInfoLevel,
		defaultAppSystemInfoOwnerID,
		defaultAppSystemInfoOwnerGroup)
	entity, err := appSystemRepo.Create(appSystemInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteAppSystemByID(id string) error {
	sql := `delete from t_meta_app_system_info where id = ?`
	_, err := appSystemRepo.Execute(sql, id)
	return err
}

func TestAppSystemRepoAll(t *testing.T) {
	TestAppSystemRepo_Execute(t)
	TestAppSystemRepo_GetAll(t)
	TestAppSystemRepo_GetByID(t)
	TestAppSystemRepo_Create(t)
	TestAppSystemRepo_Update(t)
	TestAppSystemRepo_Delete(t)
}

func TestAppSystemRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := appSystemRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestAppSystemRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_app_system_info(system_name,level,owner_id,owner_group) values(?,?,?,?);`
	tx, err := appSystemRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultAppSystemInfoAppSystemName, defaultAppSystemInfoLevel, defaultAppSystemInfoOwnerID, defaultAppSystemInfoOwnerGroup)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select system_name from t_meta_app_system_info where system_name=?`
	result, err := tx.Execute(sql,
		defaultAppSystemInfoAppSystemName,
	)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	appSystemName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if appSystemName != defaultAppSystemInfoAppSystemName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := appSystemRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		appSystemName, err := entity.Get(appSystemNameStruct)
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if appSystemName == defaultAppSystemInfoAppSystemName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestAppSystemRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := appSystemRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	appSystemName, err := entities[0].Get("AppSystemName")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineAppSystemName, appSystemName.(string), "test GetAll() failed")
}

func TestAppSystemRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := appSystemRepo.GetByID("66")
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	appSystemName, err := entity.Get(appSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineAppSystemName, appSystemName.(string), "test GetByID() failed")
}

func TestAppSystemRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createAppSystem()

	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteAppSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestAppSystemRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createAppSystem()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{appSystemNameStruct: newAppSystemName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = appSystemRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = appSystemRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	appSystemName, err := entity.Get(appSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newAppSystemName, appSystemName, "test Update() failed")
	// delete
	err = deleteAppSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestAppSystemRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createAppSystem()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteAppSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
