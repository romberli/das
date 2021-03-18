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
	defaultMySQLClusterInfoID                   = 1
	defaultMySQLClusterInfoClusterName          = "test"
	defaultMySQLClusterInfoMiddlewareClusterID  = 1
	defaultMySQLClusterInfoMonitorSystemID      = 1
	defaultMySQLClusterInfoOwnerID              = 1
	defaultMySQLClusterInfoOwnerGroup           = "2,3"
	defaultMySQLClusterInfoEnvID                = 1
	defaultMySQLClusterInfoDelFlag              = 0
	defaultMySQLClusterInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultMySQLClusterInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	clusterNameJSON                             = "cluster_name"
)

func initNewMySQLClusterInfo() *MySQLClusterInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultMySQLClusterInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultMySQLClusterInfoLastUpdateTimeString)
	return NewMySQLClusterInfoWithGlobal(
		defaultMySQLClusterInfoID,
		defaultMySQLClusterInfoClusterName,
		defaultMySQLClusterInfoMiddlewareClusterID,
		defaultMySQLClusterInfoMonitorSystemID,
		defaultMySQLClusterInfoOwnerID,
		defaultMySQLClusterInfoOwnerGroup,
		defaultMySQLClusterInfoEnvID,
		defaultMySQLClusterInfoDelFlag,
		createTime,
		lastUpdateTime)
}

func equalMySQLClusterInfo(a, b *MySQLClusterInfo) bool {
	return a.ID == b.ID &&
		a.ClusterName == b.ClusterName &&
		a.MiddlewareClusterID == b.MiddlewareClusterID &&
		a.MonitorSystemID == b.MonitorSystemID &&
		a.OwnerID == b.OwnerID &&
		a.OwnerGroup == b.OwnerGroup &&
		a.EnvID == b.EnvID &&
		a.DelFlag == b.DelFlag &&
		a.CreateTime == b.CreateTime &&
		a.LastUpdateTime == b.LastUpdateTime
}

func TestMySQLClusterEntityAll(t *testing.T) {
	TestMySQLClusterInfo_Identity(t)
	TestMySQLClusterInfo_IsDeleted(t)
	TestMySQLClusterInfo_GetCreateTime(t)
	TestMySQLClusterInfo_GetLastUpdateTime(t)
	TestMySQLClusterInfo_Get(t)
	TestMySQLClusterInfo_Set(t)
	TestMySQLClusterInfo_Delete(t)
	TestMySQLClusterInfo_MarshalJSON(t)
	TestMySQLClusterInfo_MarshalJSONWithFields(t)
}

func TestMySQLClusterInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	asst.Equal(strconv.Itoa(defaultMySQLClusterInfoID), mysqlClusterInfo.Identity(), "test Identity() failed")
}

func TestMySQLClusterInfo_IsDeleted(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	asst.False(mysqlClusterInfo.IsDeleted(), "test IsDeleted() failed")
}

func TestMySQLClusterInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	asst.True(reflect.DeepEqual(mysqlClusterInfo.CreateTime, mysqlClusterInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestMySQLClusterInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	asst.True(reflect.DeepEqual(mysqlClusterInfo.LastUpdateTime, mysqlClusterInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMySQLClusterInfo_Get(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	clusterName, err := mysqlClusterInfo.Get(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mysqlClusterInfo.ClusterName, clusterName, "test Get() failed")

	middlewareClusterID, err := mysqlClusterInfo.Get(middlewareClusterIDStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mysqlClusterInfo.MiddlewareClusterID, middlewareClusterID, "test Get() failed")

	monitorSystemID, err := mysqlClusterInfo.Get(monitorSystemIDStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mysqlClusterInfo.MonitorSystemID, monitorSystemID, "test Get() failed")

	ownerID, err := mysqlClusterInfo.Get(ownerIDStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mysqlClusterInfo.OwnerID, ownerID, "test Get() failed")

	ownerGroup, err := mysqlClusterInfo.Get(ownerGroupStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mysqlClusterInfo.OwnerGroup, ownerGroup, "test Get() failed")

	envID, err := mysqlClusterInfo.Get(envIDStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mysqlClusterInfo.EnvID, envID, "test Get() failed")
}

func TestMySQLClusterInfo_Set(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()

	newClusterName := defaultMySQLClusterInfoClusterName
	err := mysqlClusterInfo.Set(map[string]interface{}{"ClusterName": newClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newClusterName, mysqlClusterInfo.ClusterName, "test Set() failed")

	newMiddlewareClusterID := defaultMySQLClusterInfoMiddlewareClusterID
	err = mysqlClusterInfo.Set(map[string]interface{}{"MiddlewareClusterID": newMiddlewareClusterID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newMiddlewareClusterID, mysqlClusterInfo.MiddlewareClusterID, "test Set() failed")

	newMonitorSystemID := defaultMySQLClusterInfoMonitorSystemID
	err = mysqlClusterInfo.Set(map[string]interface{}{"MonitorSystemID": newMonitorSystemID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newMonitorSystemID, mysqlClusterInfo.MonitorSystemID, "test Set() failed")

	newOwnerID := defaultMySQLClusterInfoOwnerID
	err = mysqlClusterInfo.Set(map[string]interface{}{"OwnerID": newOwnerID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newOwnerID, mysqlClusterInfo.OwnerID, "test Set() failed")

	newOwnerGroup := defaultMySQLClusterInfoOwnerGroup
	err = mysqlClusterInfo.Set(map[string]interface{}{"OwnerGroup": newOwnerGroup})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newOwnerGroup, mysqlClusterInfo.OwnerGroup, "test Set() failed")

	newEnvID := defaultMySQLClusterInfoEnvID
	err = mysqlClusterInfo.Set(map[string]interface{}{"EnvID": newEnvID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newEnvID, mysqlClusterInfo.EnvID, "test Set() failed")
}

func TestMySQLClusterInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	mysqlClusterInfo.Delete()
	asst.True(mysqlClusterInfo.IsDeleted(), "test Delete() failed")
}

func TestMySQLClusterInfo_MarshalJSON(t *testing.T) {
	var mysqlClusterInfoUnmarshal *MySQLClusterInfo

	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	data, err := mysqlClusterInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &mysqlClusterInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(equalMySQLClusterInfo(mysqlClusterInfo, mysqlClusterInfoUnmarshal), "test MarshalJSON() failed")
}

func TestMySQLClusterInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	data, err := mysqlClusterInfo.MarshalJSONWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{clusterNameJSON: "test"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
