package metadata

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	defaultUserInfoID                   = 11
	defaultUserInfoUserName             = "nun"
	defaultUserInfoDepartmentName       = "dn1"
	defaultUserInfoEmployeeID           = 11
	defaultUserInfoAccountName          = "da1"
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
		defaultUserInfoDelFlag,
		createTime,
		lastUpdateTime,
		defaultUserInfoDepartmentName,
		defaultUserInfoEmployeeID,
		defaultUserInfoAccountName,
		defaultUserInfoEmail,
		defaultUserInfoTelephone,
		defaultUserInfoMobile,
		defaultUserInfoRole,
	)
}

func userEqual(a, b *UserInfo) bool {
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
}

func TestUserInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.Equal(strconv.Itoa(defaultUserInfoID), userInfo.Identity(), "test Identity() failed")
}

func TestUserInfo_IsDeleted(t *testing.T) {
	asst := assert.New(t)

	userInfo := initNewUserInfo()
	asst.False(userInfo.IsDeleted(), "test IsDeleted() failed")
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
	asst.True(userInfo.IsDeleted(), "test Delete() failed")
}

func TestUserInfo_MarshalJSON(t *testing.T) {
	var userInfoUnmarshal *UserInfo

	asst := assert.New(t)

	userInfo := initNewUserInfo()
	data, err := userInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &userInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(userEqual(userInfo, userInfoUnmarshal), "test MarshalJSON() failed")
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
