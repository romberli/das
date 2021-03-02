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
	defaultAppSystemInfoID                   = 1
	defaultAppSystemInfoAppSystemName        = "kkkk"
	defaultAppSystemInfoLevel                = 8
	defaultAppSystemInfoOwnerID              = 8
	defaultAppSystemInfoOwnerGroup           = "k"
	defaultAppSystemInfoDelFlag              = 0
	defaultAppSystemInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultAppSystemInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	appSystemNameJSON                        = "system_name"
)

func initNewAppSystemInfo() *AppSystemInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultAppSystemInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultAppSystemInfoLastUpdateTimeString)
	return NewAppSystemInfoWithGlobal(
		defaultAppSystemInfoID,
		defaultAppSystemInfoAppSystemName,
		defaultAppSystemInfoDelFlag, createTime,
		lastUpdateTime,
		defaultAppSystemInfoLevel,
		defaultAppSystemInfoOwnerID,
		defaultAppSystemInfoOwnerGroup)
}

func appSystemEqual(a, b *AppSystemInfo) bool {
	return a.ID == b.ID && a.AppSystemName == b.AppSystemName && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime && a.Level == b.Level && a.OwnerID == b.OwnerID && a.OwnerGroup == b.OwnerGroup
}

func TestAppSystemEntityAll(t *testing.T) {
	TestAppSystemInfo_Identity(t)
	TestAppSystemInfo_IsDeleted(t)
	TestAppSystemInfo_GetCreateTime(t)
	TestAppSystemInfo_GetLastUpdateTime(t)
	TestAppSystemInfo_Get(t)
	TestAppSystemInfo_Set(t)
	TestAppSystemInfo_Delete(t)
	TestAppSystemInfo_MarshalJSON(t)
	TestAppSystemInfo_MarshalJSONWithFields(t)
}

func TestAppSystemInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppSystemInfo()
	asst.Equal(strconv.Itoa(defaultAppSystemInfoID), appSystemInfo.Identity(), "test Identity() failed")
}

func TestAppSystemInfo_IsDeleted(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppSystemInfo()
	asst.False(appSystemInfo.IsDeleted(), "test IsDeleted() failed")
}

func TestAppSystemInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppSystemInfo()
	asst.True(reflect.DeepEqual(appSystemInfo.CreateTime, appSystemInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestAppSystemInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppSystemInfo()
	asst.True(reflect.DeepEqual(appSystemInfo.LastUpdateTime, appSystemInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestAppSystemInfo_Get(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppSystemInfo()
	appSystemName, err := appSystemInfo.Get(appSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(appSystemInfo.AppSystemName, appSystemName, "test Get() failed")
}

func TestAppSystemInfo_Set(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppSystemInfo()
	newAppSystemName := "new_appSystem"
	err := appSystemInfo.Set(map[string]interface{}{"AppSystemName": newAppSystemName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(newAppSystemName, appSystemInfo.AppSystemName, "test Set() failed")
}

func TestAppSystemInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppSystemInfo()
	appSystemInfo.Delete()
	asst.True(appSystemInfo.IsDeleted(), "test Delete() failed")
}

func TestAppSystemInfo_MarshalJSON(t *testing.T) {
	var appSystemInfoUnmarshal *AppSystemInfo

	asst := assert.New(t)

	appSystemInfo := initNewAppSystemInfo()
	data, err := appSystemInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &appSystemInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(appSystemEqual(appSystemInfo, appSystemInfoUnmarshal), "test MarshalJSON() failed")
}

func TestAppSystemInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppSystemInfo()
	data, err := appSystemInfo.MarshalJSONWithFields(appSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{appSystemNameJSON: "kkkk"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
