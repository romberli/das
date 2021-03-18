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
	// modify these connection information
	// addr          = "192.168.137.11:3306"
	// dbName        = "das"
	// dbUser        = "root"
	// dbPass        = "root"
	newEnvName    = "newTest"
	onlineEnvName = "test"
)

var envRepo = initEnvRepo()

func initEnvRepo() *EnvRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initEnvRepo() failed", err))
		return nil
	}

	return NewEnvRepo(pool)
}

func createEnv() (dependency.Entity, error) {
	envInfo := NewEnvInfoWithDefault(defaultMSInfoMSName)
	entity, err := envRepo.Create(envInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteEnvByID(id string) error {
	sql := `delete from t_meta_env_info where id = ?`
	_, err := envRepo.Execute(sql, id)
	return err
}

func TestEnvRepoAll(t *testing.T) {
	TestEnvRepo_Execute(t)
	TestEnvRepo_GetAll(t)
	TestEnvRepo_GetByID(t)
	TestEnvRepo_Create(t)
	TestEnvRepo_Update(t)
	TestEnvRepo_Delete(t)
}

func TestEnvRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := envRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestEnvRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_env_info(env_name) values(?);`
	tx, err := envRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultMSInfoMSName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select env_name from t_meta_env_info where env_name=?`
	result, err := tx.Execute(sql, defaultMSInfoMSName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	envName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if envName != defaultMSInfoMSName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := envRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		envName, err := entity.Get(mSNameStruct)
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if envName == defaultMSInfoMSName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestEnvRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := envRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	envName, err := entities[0].Get("EnvName")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineEnvName, envName.(string), "test GetAll() failed")
}

func TestEnvRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := envRepo.GetByID("1")
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	envName, err := entity.Get(mSNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineEnvName, envName.(string), "test GetByID() failed")
}

func TestEnvRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createEnv()

	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteEnvByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestEnvRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createEnv()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{mSNameStruct: newEnvName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = envRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = envRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	envName, err := entity.Get(mSNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newEnvName, envName, "test Update() failed")
	// delete
	err = deleteEnvByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestEnvRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createEnv()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteEnvByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
