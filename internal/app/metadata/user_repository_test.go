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
	onlineUserName        = "online"
	newUserName           = "nun"
	OnlineUserAccountName = "1"
	OnlineUserEmployeeID  = "1"
	OnlineUserEmail       = "1"
	OnlineUserTelephone   = "1"
	OnlineUserMobile      = "1"
)

var userRepo = initUserRepo()

func initUserRepo() *UserRepo {
	pool, err := mysql.NewPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initUserRepo() failed", err))
		return nil
	}

	return NewUserRepo(pool)
}

func createUser() (metadata.User, error) {
	userInfo := NewUserInfoWithDefault(
		defaultUserInfoUserName,
		defaultUserInfoDepartmentName,
		defaultUserInfoAccountName,
		defaultUserInfoEmail,
		defaultUserInfoRole,
	)
	entity, err := userRepo.Create(userInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteUserByID(id int) error {
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
	TestUserRepo_GetByName(t)
	TestUserRepo_GetByAccountName(t)
	TestUserRepo_GetByEmail(t)
	TestUserRepo_GetByTelephone(t)
	TestUserRepo_GetByMobile(t)
	TestUserRepo_GetID(t)
	TestUserRepo_GetByEmployeeID(t)
}

func TestUserRepo_GetByName(t *testing.T) {
	asst := assert.New(t)

	entity, err := userRepo.GetByName(onlineUserName)
	asst.Nil(err, common.CombineMessageWithError("test GetByName() failed", err))
	userName := entity[0].GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test GetByName() failed", err))
	asst.Equal(onlineUserName, userName, "test GetByName() failed")
}

func TestUserRepo_GetByAccountName(t *testing.T) {
	asst := assert.New(t)

	entity, err := userRepo.GetByAccountName(OnlineUserAccountName)
	asst.Nil(err, common.CombineMessageWithError("test GetByAccountName failed", err))
	userName := entity.GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test GetByAccountName failed", err))
	asst.Equal(onlineUserName, userName, "test GetByAccountName failed")
}

func TestUserRepo_GetByEmail(t *testing.T) {
	asst := assert.New(t)

	entity, err := userRepo.GetByEmail(OnlineUserEmail)
	asst.Nil(err, common.CombineMessageWithError("test GetByEmail failed", err))
	userName := entity.GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test GetByEmail failed", err))
	asst.Equal(onlineUserName, userName, "test GetByEmail failed")

}

func TestUserRepo_GetByTelephone(t *testing.T) {
	asst := assert.New(t)

	entity, err := userRepo.GetByTelephone(OnlineUserTelephone)
	asst.Nil(err, common.CombineMessageWithError("test GetByTelephone failed", err))
	userName := entity.GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test GetByTelephone failed", err))
	asst.Equal(onlineUserName, userName, "test GetByTelephone failed")
}

func TestUserRepo_GetByMobile(t *testing.T) {
	asst := assert.New(t)

	entity, err := userRepo.GetByMobile(OnlineUserMobile)
	asst.Nil(err, common.CombineMessageWithError("test GetByMobile failed", err))
	userName := entity.GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test GetByMobile failed", err))
	asst.Equal(onlineUserName, userName, "test GetByMobile failed")
}

func TestUserRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	UserID, err := userRepo.GetID(OnlineUserAccountName)
	asst.Nil(err, common.CombineMessageWithError("test GetID failed", err))
	entity, err := userRepo.GetByID(UserID)
	asst.Nil(err, common.CombineMessageWithError("test GetID failed", err))
	userAccountName := entity.GetAccountName()
	asst.Equal(OnlineUserAccountName, userAccountName, "test GetID failed")

}

func TestUserRepo_GetByEmployeeID(t *testing.T) {
	asst := assert.New(t)

	entity, err := userRepo.GetByEmployeeID(OnlineUserEmployeeID)
	asst.Nil(err, common.CombineMessageWithError("test GetByEmployeeID failed", err))
	userName := entity.GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test GetByEmployeeID failed", err))
	asst.Equal(onlineUserName, userName, "test GetByEmployeeID failed")
}

func TestUserRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := userRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
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
		userName := entity.GetUserName()
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
	userName := entities[0].GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineUserName, userName, "test GetAll() failed")
}

func TestUserRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := userRepo.GetByID(2)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	userName := entity.GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineUserName, userName, "test GetByID() failed")
}

func TestUserRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createUser()

	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	id := entity.Identity()
	err = deleteUserByID(id)
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
	userName := entity.GetUserName()
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
