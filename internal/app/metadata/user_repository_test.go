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
	onlineUserName = "1"
	newUserName    = "nun"
)

var userRepo = initUserRepo()

func initUserRepo() *UserRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initUserRepo() failed", err))
		return nil
	}

	return NewUserRepo(pool)
}

func createUser() (dependency.Entity, error) {
	userInfo := NewUserInfoWithDefault(
		defaultUserInfoUserName,
		defaultUserInfoDepartmentName,
		defaultUserInfoEmployeeID,
		defaultUserInfoAccountName,
		defaultUserInfoEmail,
		defaultUserInfoTelephone,
		defaultUserInfoMobile,
		defaultUserInfoRole,
	)
	entity, err := userRepo.Create(userInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteUserByID(id string) error {
	sql := `delete from t_meta_user_info where id = ?`
	_, err := userRepo.Execute(sql, id)
	return err
}

func TestUserRepoAll(t *testing.T) {
	TestUserRepo_Execute(t)
	TestUserRepo_GetAll(t)
	TestUserRepo_GetByID(t)
	TestUserRepo_Create(t)
	TestUserRepo_Update(t)
	TestUserRepo_Delete(t)
}

func TestUserRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := userRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestUserRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_user_info(user_name,department_name,employee_id,account_name,email,telephone,mobile,role) values(?,?,?,?,?,?,?,?);`
	tx, err := userRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql,
		defaultUserInfoUserName,
		defaultUserInfoDepartmentName,
		defaultUserInfoEmployeeID,
		defaultUserInfoAccountName,
		defaultUserInfoEmail,
		defaultUserInfoTelephone,
		defaultUserInfoMobile,
		defaultUserInfoRole)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select user_name from t_meta_user_info where user_name=?`
	result, err := tx.Execute(sql,
		defaultUserInfoUserName,
	)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	userName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if userName != defaultUserInfoUserName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := userRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		userName, err := entity.Get(userNameStruct)
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if userName == defaultUserInfoUserName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestUserRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := userRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	userName, err := entities[0].Get("UserName")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineUserName, userName.(string), "test GetAll() failed")
}

func TestUserRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := userRepo.GetByID("66")
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	userName, err := entity.Get(userNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineUserName, userName.(string), "test GetByID() failed")
}

func TestUserRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createUser()

	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteUserByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestUserRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createUser()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{userNameStruct: newUserName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = userRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = userRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	userName, err := entity.Get(userNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newUserName, userName, "test Update() failed")
	// delete
	err = deleteUserByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestUserRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createUser()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteUserByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
