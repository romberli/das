package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceAll(t *testing.T) {
	TestUserService_GetUsers(t)
	TestUserService_GetAll(t)
	TestUserService_GetByID(t)
	TestUserService_Create(t)
	TestUserService_Update(t)
	TestUserService_Delete(t)
	TestUserService_Marshal(t)
	TestUserService_MarshalWithFields(t)
}

func TestUserService_GetUsers(t *testing.T) {
	asst := assert.New(t)

	s := NewUserService(userRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	Users := s.GetUsers()
	asst.Greater(len(Users), constant.ZeroInt, "test GetEnvs() failed")
}

func TestUserService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewUserService(userRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	Users := s.GetUsers()
	asst.Greater(len(Users), constant.ZeroInt, "test GetEnvs() failed")
}

func TestUserService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewUserService(userRepo)
	err := s.GetByID(66)
	asst.Nil(err, "test GetByID() failed")
	id := s.Users[constant.ZeroInt].Identity()
	asst.Equal(66, id, "test GetByID() failed")
}

func TestUserService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewUserService(userRepo)
	err := s.Create(map[string]interface{}{
		userNameStruct:       defaultUserInfoUserName,
		departmentNameStruct: defaultUserInfoDepartmentName,
		employeeIDStruct:     defaultUserInfoEmployeeID,
		//	accountNameStruct:    defaultUserInfoAccountName,
		emailStruct:     defaultUserInfoEmail,
		telephoneStruct: defaultUserInfoTelephone,
		mobileStruct:    defaultUserInfoMobile,
		roleStruct:      defaultUserInfoRole,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteUserByID(s.Users[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestUserService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createUser()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewUserService(userRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{userNameStruct: newUserName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	userName := s.GetUsers()[constant.ZeroInt].GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newUserName, userName)
	// delete
	err = deleteUserByID(s.Users[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestUserService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createUser()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewUserService(userRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteUserByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestUserService_Marshal(t *testing.T) {
	var UsersUnmarshal []*UserInfo

	asst := assert.New(t)

	s := NewUserService(userRepo)
	err := s.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	data, err := s.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	err = json.Unmarshal(data, &UsersUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	Users := s.GetUsers()
	for i := 0; i < len(Users); i++ {
		entity := Users[i]
		entityUnmarshal := UsersUnmarshal[i]
		asst.True(userStructEqual(entity.(*UserInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestUserService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createUser()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewUserService(userRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(userNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(userNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteUserByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
