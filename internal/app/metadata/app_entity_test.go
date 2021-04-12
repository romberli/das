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
	defaultAppInfoID                   = 1
	defaultAppInfoAppName              = "dfname"
	defaultAppInfoLevel                = 1
	defaultAppInfoOwnerID              = 8
	defaultAppInfoOwnerGroup           = "k"
	defaultAppInfoDelFlag              = 0
	defaultAppInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultAppInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	appSystemNameJSON                  = "app_name"
)

func initNewAppInfo() *AppInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultAppInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultAppInfoLastUpdateTimeString)
	return NewAppInfo(
		appRepo,
		defaultAppInfoID,
		defaultAppInfoAppName,
		defaultAppInfoLevel,
		defaultAppInfoOwnerID,
		defaultAppInfoDelFlag, createTime,
		lastUpdateTime,
	)
}

func appSystemStructEqual(a, b *AppInfo) bool {
	return a.ID == b.ID && a.AppName == b.AppName && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime && a.Level == b.Level && a.OwnerID == b.OwnerID
}

func TestAppEntityAll(t *testing.T) {
	TestAppInfo_Identity(t)
	TestAppInfo_GetAppName(t)
	TestAppInfo_GetLevel(t)
	TestAppInfo_GetOwnerID(t)
	TestAppInfo_GetDelFlag(t)
	TestAppInfo_GetCreateTime(t)
	TestAppInfo_GetLastUpdateTime(t)
	TestAppInfo_Set(t)
	TestAppInfo_Delete(t)
	TestAppInfo_MarshalJSON(t)
	TestAppInfo_MarshalJSONWithFields(t)
	TestAppInfo_GetDBIDList(t)
	TestAppInfo_AddAppDB(t)
	TestAppInfo_DeleteAppDB(t)
}

func TestAppInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.Equal(defaultAppInfoID, appSystemInfo.Identity(), "test Identity() failed")
}

func TestAppInfo_GetAppName(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.Equal(defaultAppInfoAppName, appSystemInfo.GetAppName(), "test GetAppName() failed")
}

func TestAppInfo_GetLevel(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.Equal(defaultAppInfoLevel, appSystemInfo.GetLevel(), "test GetLevel() failed")
}

func TestAppInfo_GetOwnerID(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.Equal(defaultAppInfoOwnerID, appSystemInfo.GetOwnerID(), "test GetLevel() failed")
}

func TestAppInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.Equal(constant.ZeroInt, appSystemInfo.GetDelFlag(), "test GetDelFlag() failed")
}

func TestAppInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.True(reflect.DeepEqual(appSystemInfo.CreateTime, appSystemInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestAppInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.True(reflect.DeepEqual(appSystemInfo.LastUpdateTime, appSystemInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestAppInfo_Set(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	newAppName := "new_appSystem"
	err := appSystemInfo.Set(map[string]interface{}{"AppName": newAppName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(newAppName, appSystemInfo.AppName, "test Set() failed")
}

func TestAppInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	appSystemInfo.Delete()
	asst.Equal(1, appSystemInfo.GetDelFlag(), "test Delete() failed")
}

func TestAppInfo_MarshalJSON(t *testing.T) {
	var appSystemInfoUnmarshal *AppInfo

	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	data, err := appSystemInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &appSystemInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(appSystemStructEqual(appSystemInfo, appSystemInfoUnmarshal), "test MarshalJSON() failed")
}

func TestAppInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	data, err := appSystemInfo.MarshalJSONWithFields(appAppNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{appSystemNameJSON: defaultAppInfoAppName})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}

func TestAppInfo_GetDBIDList(t *testing.T) {
	var dbIDList []int

	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	dbIDList, err := appSystemInfo.GetDBIDList()
	asst.Nil(err, common.CombineMessageWithError("test GetDBIDList() failed", err))
	defaultDBIDList := []int{1, 2}
	for i := 0; i < len(dbIDList); i++ {
		asst.Equal(dbIDList[i], defaultDBIDList, "test GetDBIDList() failed")
	}
}

func TestAppInfo_AddAppDB(t *testing.T) {
	var dbIDList []int

	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	err := appSystemInfo.AddDB(3)
	dbIDList, err = appSystemInfo.GetDBIDList()
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	asst.Equal(0, len(dbIDList))
	// delete
	err = appSystemInfo.DeleteDB(3)
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
}

func TestAppInfo_DeleteAppDB(t *testing.T) {
	var dbIDList []int

	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	err := appSystemInfo.DeleteDB(2)
	dbIDList, err = appSystemInfo.GetDBIDList()
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	asst.Equal(0, len(dbIDList))
	// add
	err = appSystemInfo.AddDB(2)
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
}
