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
	defaultMSInfoID                   = 1
	defaultMSInfoDelFlag              = 0
	defaultMSInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultMSInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	mSNameJSON                        = "system_name"
)

func initNewMSInfo() *MSInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultMSInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultMSInfoLastUpdateTimeString)
	return NewMSInfoWithGlobal(defaultMSInfoID, defaultMSInfoMSName, defaultMSInfoSystemType, defaultMSInfoHostIp, defaultMSInfoPortNum, defaultMSInfoPortNumSlow, defaultMSInfoBaseUrl, defaultMSInfoDelFlag, createTime, lastUpdateTime)
}

func mSEqual(a, b *MSInfo) bool {
	return a.ID == b.ID && a.MSName == b.MSName && a.SystemType == b.SystemType && a.HostIp == b.HostIp && a.PortNum == b.PortNum && a.PortNumSlow == b.PortNumSlow && a.BaseUrl == b.BaseUrl && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime
}

func TestMSEntityAll(t *testing.T) {
	TestMSInfo_Identity(t)
	TestMSInfo_IsDeleted(t)
	TestMSInfo_GetCreateTime(t)
	TestMSInfo_GetLastUpdateTime(t)
	TestMSInfo_Get(t)
	TestMSInfo_Set(t)
	TestMSInfo_Delete(t)
	TestMSInfo_MarshalJSON(t)
	TestMSInfo_MarshalJSONWithFields(t)
}

func TestMSInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	mSInfo := initNewMSInfo()
	asst.Equal(strconv.Itoa(defaultMSInfoID), mSInfo.Identity(), "test Identity() failed")
}

func TestMSInfo_IsDeleted(t *testing.T) {
	asst := assert.New(t)

	mSInfo := initNewMSInfo()
	asst.False(mSInfo.IsDeleted(), "test IsDeleted() failed")
}

func TestMSInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	mSInfo := initNewMSInfo()
	asst.True(reflect.DeepEqual(mSInfo.CreateTime, mSInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestMSInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	mSInfo := initNewMSInfo()
	asst.True(reflect.DeepEqual(mSInfo.LastUpdateTime, mSInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMSInfo_Get(t *testing.T) {
	asst := assert.New(t)

	mSInfo := initNewMSInfo()
	mSName, err := mSInfo.Get(mSNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mSInfo.MSName, mSName, "test Get() failed")
}

func TestMSInfo_Set(t *testing.T) {
	asst := assert.New(t)

	mSInfo := initNewMSInfo()
	newMSName := "new_mS"
	err := mSInfo.Set(map[string]interface{}{"MSName": newMSName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(newMSName, mSInfo.MSName, "test Set() failed")
}

func TestMSInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	mSInfo := initNewMSInfo()
	mSInfo.Delete()
	asst.True(mSInfo.IsDeleted(), "test Delete() failed")
}

func TestMSInfo_MarshalJSON(t *testing.T) {
	var mSInfoUnmarshal *MSInfo

	asst := assert.New(t)

	mSInfo := initNewMSInfo()
	data, err := mSInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &mSInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(mSEqual(mSInfo, mSInfoUnmarshal), "test MarshalJSON() failed")
}

func TestMSInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	mSInfo := initNewMSInfo()
	data, err := mSInfo.MarshalJSONWithFields(mSNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{mSNameJSON: "ms"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
