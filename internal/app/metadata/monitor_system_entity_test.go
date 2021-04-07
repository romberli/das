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
	defaultMonitorSystemInfoID                   = 1
	defaultMonitorSystemInfoSystemName           = "monitor_system"
	defaultMonitorSystemInfoSystemType           = 1
	defaultMonitorSystemInfoHostIP               = "0.0.0.0"
	defaultMonitorSystemInfoPortNum              = 3306
	defaultMonitorSystemInfoPortNumSlow          = 3307
	defaultMonitorSystemInfoBaseUrl              = "http://127.0.0.1/prometheus/api/v1/"
	defaultMonitorSystemInfoEnvID                = 1
	defaultMonitorSystemInfoDelFlag              = 0
	defaultMonitorSystemInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultMonitorSystemInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	monitorSystemNameJSON                        = "system_name"
)

func initNewMonitorSystemInfo() *MonitorSystemInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultMonitorSystemInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultMonitorSystemInfoLastUpdateTimeString)
	return NewMonitorSystemInfo(monitorSystemRepo, defaultMonitorSystemInfoID, defaultMonitorSystemInfoSystemName,
		defaultMonitorSystemInfoSystemType, defaultMonitorSystemInfoHostIP, defaultMonitorSystemInfoPortNum,
		defaultMonitorSystemInfoPortNumSlow, defaultMonitorSystemInfoBaseUrl, defaultMonitorSystemInfoEnvID,
		defaultMonitorSystemInfoDelFlag, createTime, lastUpdateTime)
}

func monitorSystemEqual(a, b *MonitorSystemInfo) bool {
	return a.ID == b.ID && a.MonitorSystemName == b.MonitorSystemName && a.MonitorSystemType == b.MonitorSystemType &&
		a.MonitorSystemHostIP == b.MonitorSystemHostIP && a.MonitorSystemPortNum == b.MonitorSystemPortNum &&
		a.MonitorSystemPortNumSlow == b.MonitorSystemPortNumSlow && a.BaseURL == b.BaseURL && a.EnvID == b.EnvID &&
		a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime
}

func TestMonitorSystemEntityAll(t *testing.T) {
	TestMonitorSystemInfo_Identity(t)
	TestMonitorSystemInfo_GetMonitorSystemName(t)
	TestMonitorSystemInfo_GetMonitorSystemType(t)
	TestMonitorSystemInfo_GetMonitorSystemHostIP(t)
	TestMonitorSystemInfo_GetMonitorSystemPortNum(t)
	TestMonitorSystemInfo_GetMonitorSystemPortNumSlow(t)
	TestMonitorSystemInfo_GetBaseURL(t)
	TestMonitorSystemInfo_GetEnvID(t)
	TestMonitorSystemInfo_GetDelFlag(t)
	TestMonitorSystemInfo_GetCreateTime(t)
	TestMonitorSystemInfo_GetLastUpdateTime(t)
	TestMonitorSystemInfo_Set(t)
	TestMonitorSystemInfo_Delete(t)
	TestMonitorSystemInfo_MarshalJSON(t)
	TestMonitorSystemInfo_MarshalJSONWithFields(t)
}

func TestMonitorSystemInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	asst.Equal(defaultMonitorSystemInfoID, monitorSystemInfo.Identity(), "test Identity() failed")
}

func TestMonitorSystemInfo_GetMonitorSystemName(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	asst.Equal(defaultMonitorSystemInfoSystemName, monitorSystemInfo.GetSystemName(), "test GetSystemName() failed")
}

func TestMonitorSystemInfo_GetMonitorSystemType(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	asst.Equal(defaultMonitorSystemInfoSystemType, monitorSystemInfo.GetSystemType(), "test GetSystemType() failed")
}

func TestMonitorSystemInfo_GetMonitorSystemHostIP(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	asst.Equal(defaultMonitorSystemInfoHostIP, monitorSystemInfo.GetHostIP(), "test GetHostIP() failed")
}

func TestMonitorSystemInfo_GetMonitorSystemPortNum(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	asst.Equal(defaultMonitorSystemInfoPortNum, monitorSystemInfo.GetPortNum(), "test GetPortNum() failed")
}

func TestMonitorSystemInfo_GetMonitorSystemPortNumSlow(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	asst.Equal(defaultMonitorSystemInfoPortNumSlow, monitorSystemInfo.GetPortNumSlow(), "test GetPortNumSlow() failed")
}

func TestMonitorSystemInfo_GetBaseURL(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	asst.Equal(defaultMonitorSystemInfoBaseUrl, monitorSystemInfo.GetBaseURL(), "test GetBaseURL() failed")
}

func TestMonitorSystemInfo_GetEnvID(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	asst.Equal(defaultMonitorSystemInfoEnvID, monitorSystemInfo.GetEnvID(), "test GetEnvID() failed")
}

func TestMonitorSystemInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	asst.Equal(defaultMonitorSystemInfoDelFlag, monitorSystemInfo.GetDelFlag(), "test GetDelFlag() failed")
}

func TestMonitorSystemInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	asst.True(reflect.DeepEqual(monitorSystemInfo.CreateTime, monitorSystemInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestMonitorSystemInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	asst.True(reflect.DeepEqual(monitorSystemInfo.LastUpdateTime, monitorSystemInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMonitorSystemInfo_Set(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	newMonitorSystemName := "new_monitor_system"
	err := monitorSystemInfo.Set(map[string]interface{}{"MonitorSystemName": newMonitorSystemName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(newMonitorSystemName, monitorSystemInfo.MonitorSystemName, "test Set() failed")
}

func TestMonitorSystemInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	monitorSystemInfo.Delete()
	asst.Equal(1, monitorSystemInfo.DelFlag, "test Delete() failed")
}

func TestMonitorSystemInfo_MarshalJSON(t *testing.T) {
	var monitorSystemInfoUnmarshal *MonitorSystemInfo

	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	data, err := monitorSystemInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &monitorSystemInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(monitorSystemEqual(monitorSystemInfo, monitorSystemInfoUnmarshal), "test MarshalJSON() failed")
}

func TestMonitorSystemInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := initNewMonitorSystemInfo()
	data, err := monitorSystemInfo.MarshalJSONWithFields(monitorSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{monitorSystemNameJSON: "monitor_system"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
