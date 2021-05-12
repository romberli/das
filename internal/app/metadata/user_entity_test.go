package metadata

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	defaultUserInfoID                   = 1
	defaultUserInfoUserName             = "nun"
	defaultUserInfoDepartmentName       = "dn1"
	defaultUserInfoEmployeeID           = "11"
	defaultUserInfoAccountName          = "da"
	defaultUserInfoEmail                = "e1"
	defaultUserInfoTelephone            = "t1"
	defaultUserInfoMobile               = "m1"
	defaultUserInfoRole                 = 11
	defaultUserInfoDelFlag              = 0
	defaultUserInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultUserInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	userNameJSON                        = "user_name"
)

func initNewUserInfo() *UserInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultUserInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultUserInfoLastUpdateTimeString)
	return NewUserInfoWithGlobal(
		defaultUserInfoID,
		defaultUserInfoUserName,
		defaultUserInfoDepartmentName,
		defaultUserInfoEmployeeID,
		defaultUserInfoAccountName,
		defaultUserInfoEmail,
		defaultUserInfoTelephone,
		defaultUserInfoMobile,
		defaultUserInfoRole,
		defaultUserInfoDelFlag,
		createTime,
		lastUpdateTime,
	)
}

func userStructEqual(a, b *UserInfo) bool {
	return a.ID == b.ID && a.UserName == b.UserName && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime && a.DepartmentName == b.DepartmentName && a.EmployeeID == b.EmployeeID && a.AccountName == b.AccountName && a.Email == b.Email && a.Telephone == b.Telephone && a.Mobile == b.Mobile && a.Role == b.Role
}

func TestUserEntityAll(t *testing.T) {
	TestUserInfo_Identity(t)
	TestUserInfo_IsDeleted(t)
	TestUserInfo_GetCreateTime(t)
	TestUserInfo_GetLastUpdateTime(t)
	TestUserInfo_Get(t)
	TestUserInfo_Set(t)
	TestUserInfo_Delete(t)
	TestUserInfo_MarshalJSON(t)
	TestUserInfo_MarshalJSONWithFields(t)
	TestUserInfo_GetUserName(t)
	TestUserInfo_GetDepartmentName(t)
	TestUserInfo_GetEmployeeID(t)
	TestUserInfo_GetAccountName(t)
	TestUserInfo_GetEmail(t)
	TestUserInfo_GetTelephone(t)
	TestUserInfo_GetMobile(t)
	TestUserInfo_GetDelFlag(t)
}

func TestUserInfo_GetUserName(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.Equal(defaultUserInfoUserName, userInfo.GetUserName(), "test GetUserName() failed")
}

func TestUserInfo_GetDepartmentName(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.Equal(defaultUserInfoDepartmentName, userInfo.GetDepartmentName(), "test GetDepartmentName() failed")
}

func TestUserInfo_GetEmployeeID(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.Equal(defaultUserInfoEmployeeID, userInfo.GetEmployeeID(), "test GetEmployeeID() failed")
}

func TestUserInfo_GetAccountName(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.Equal(defaultUserInfoAccountName, userInfo.GetAccountName(), "test GetAccountName() failed")
}

func TestUserInfo_GetEmail(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.Equal(defaultUserInfoEmail, userInfo.GetEmail(), "test GetEmail() failed")
}

func TestUserInfo_GetTelephone(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.Equal(defaultUserInfoTelephone, userInfo.GetTelephone(), "test GetTelephone() failed")
}

func TestUserInfo_GetMobile(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.Equal(defaultUserInfoMobile, userInfo.GetMobile(), "test GetMobile() failed")
}

func TestUserInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.Equal(defaultUserInfoDelFlag, userInfo.GetDelFlag(), "test GetDelFlag() failed")
}

func TestUserInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.Equal(defaultUserInfoID, userInfo.Identity(), "test Identity() failed")
}

func TestUserInfo_IsDeleted(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.False(userInfo.GetDelFlag() != constant.ZeroInt, "test IsDeleted() failed")
}

func TestUserInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.True(reflect.DeepEqual(userInfo.CreateTime, userInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestUserInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.True(reflect.DeepEqual(userInfo.LastUpdateTime, userInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestUserInfo_Get(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	userName, err := userInfo.Get(userNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(userInfo.UserName, userName, "test Get() failed")
}

func TestUserInfo_Set(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	newUserName := "new_user"
	err := userInfo.Set(map[string]interface{}{"UserName": newUserName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(newUserName, userInfo.UserName, "test Set() failed")
}

func TestUserInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	userInfo.Delete()
	asst.True(userInfo.GetDelFlag() != constant.ZeroInt, "test Delete() failed")
}

func TestUserInfo_MarshalJSON(t *testing.T) {
	var userInfoUnmarshal *UserInfo

	asst := assert.New(t)

	userInfo := initNewUserInfo()
	data, err := userInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &userInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(userStructEqual(userInfo, userInfoUnmarshal), "test MarshalJSON() failed")
}

func TestUserInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	data, err := userInfo.MarshalJSONWithFields(userNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{userNameJSON: "nun"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
