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
	envAddr       = "127.0.0.1:3306"
	envDBName     = "das"
	envDBUser     = "root"
	envDBPass     = "rootroot"
	newEnvName    = "newTest"
	onlineEnvName = "rel"
	onlineID      = 2
)

var envRepo = initEnvRepo()

func initEnvRepo() *EnvRepo {
	pool, err := mysql.NewPoolWithDefault(envAddr, envDBName, envDBUser, envDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initEnvRepo() failed", err))
		return nil
	}

	return NewEnvRepo(pool)
}

func createEnv() (metadata.Env, error) {
	envInfo := NewEnvInfoWithDefault(defaultEnvInfoEnvName)
	entity, err := envRepo.Create(envInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteEnvByID(id int) error {
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
	TestEnvRepo_GetID(t)
	TestEnvRepo_GetEnvByName(t)

}

func TestEnvRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := envRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestEnvRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_env_info(env_name) values(?);`
	tx, err := envRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultEnvInfoEnvName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select env_name from t_meta_env_info where env_name=?`
	result, err := tx.Execute(sql, defaultEnvInfoEnvName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	envName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if envName != defaultEnvInfoEnvName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	envs, err := envRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, env := range envs {
		envName := env.GetEnvName()
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if envName == defaultEnvInfoEnvName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestEnvRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := envRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	envName := entities[0].GetEnvName()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineEnvName, envName, "test GetAll() failed")
}

func TestEnvRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := envRepo.GetByID(2)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	envName := entity.GetEnvName()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineEnvName, envName, "test GetByID() failed")
}

func TestEnvRepo_Create(t *testing.T) {
	asst := assert.New(t)

	env, err := createEnv()

	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteEnvByID(env.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestEnvRepo_Update(t *testing.T) {
	asst := assert.New(t)

	env, err := createEnv()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = env.Set(map[string]interface{}{envNameStruct: newEnvName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = envRepo.Update(env)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	env, err = envRepo.GetByID(env.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	envName := env.GetEnvName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newEnvName, envName, "test Update() failed")
	// delete
	err = deleteEnvByID(env.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestEnvRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	env, err := createEnv()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	ID, err := envRepo.GetID(env.GetEnvName())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = envRepo.Delete(ID)
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteEnvByID(env.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestEnvRepo_GetEnvByName(t *testing.T) {
	asst := assert.New(t)

	entity, err := envRepo.GetEnvByName("rel")
	asst.Nil(err, common.CombineMessageWithError("test GetEnvByName() failed", err))
	envName := entity.GetEnvName()
	asst.Nil(err, common.CombineMessageWithError("test GetEnvByName() failed", err))
	asst.Equal(onlineEnvName, envName, "test GetEnvByName() failed")
}

func TestEnvRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	env, err := envRepo.GetEnvByName("rel")
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	ID, err := envRepo.GetID(env.GetEnvName())
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.Equal(onlineID, ID, "test GetID() failed")
}
