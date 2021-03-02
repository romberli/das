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
	defaultMYSQLClusterInfoID                   = 1
	defaultMYSQLClusterInfoClusterName          = "test"
	defaultMYSQLClusterInfoMiddlewareClusterID  = 1
	defaultMYSQLClusterInfoMonitorSystemID      = 1
	defaultMYSQLClusterInfoOwnerID              = 1
	defaultMYSQLClusterInfoOwnerGroup           = "2,3"
	defaultMYSQLClusterInfoEnvID                = 1
	defaultMYSQLClusterInfoDelFlag              = 0
	defaultMYSQLClusterInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultMYSQLClusterInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	clusterNameJSON                             = "cluster_name"
)

func initNewMYSQLClusterInfo() *MYSQLClusterInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultMYSQLClusterInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultMYSQLClusterInfoLastUpdateTimeString)
	return NewMYSQLClusterInfoWithGlobal(
		defaultMYSQLClusterInfoID,
		defaultMYSQLClusterInfoClusterName,
		defaultMYSQLClusterInfoMiddlewareClusterID,
		defaultMYSQLClusterInfoMonitorSystemID,
		defaultMYSQLClusterInfoOwnerID,
		defaultMYSQLClusterInfoOwnerGroup,
		defaultMYSQLClusterInfoEnvID,
		defaultMYSQLClusterInfoDelFlag,
		createTime,
		lastUpdateTime)
}

func equalMYSQLClusterInfo(a, b *MYSQLClusterInfo) bool {
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

func TestMYSQLClusterEntityAll(t *testing.T) {
	TestMYSQLClusterInfo_Identity(t)
	TestMYSQLClusterInfo_IsDeleted(t)
	TestMYSQLClusterInfo_GetCreateTime(t)
	TestMYSQLClusterInfo_GetLastUpdateTime(t)
	TestMYSQLClusterInfo_Get(t)
	TestMYSQLClusterInfo_Set(t)
	TestMYSQLClusterInfo_Delete(t)
	TestMYSQLClusterInfo_MarshalJSON(t)
	TestMYSQLClusterInfo_MarshalJSONWithFields(t)
}

func TestMYSQLClusterInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMYSQLClusterInfo()
	asst.Equal(strconv.Itoa(defaultMYSQLClusterInfoID), mysqlClusterInfo.Identity(), "test Identity() failed")
}

func TestMYSQLClusterInfo_IsDeleted(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMYSQLClusterInfo()
	asst.False(mysqlClusterInfo.IsDeleted(), "test IsDeleted() failed")
}

func TestMYSQLClusterInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMYSQLClusterInfo()
	asst.True(reflect.DeepEqual(mysqlClusterInfo.CreateTime, mysqlClusterInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestMYSQLClusterInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMYSQLClusterInfo()
	asst.True(reflect.DeepEqual(mysqlClusterInfo.LastUpdateTime, mysqlClusterInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMYSQLClusterInfo_Get(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMYSQLClusterInfo()
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

func TestMYSQLClusterInfo_Set(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMYSQLClusterInfo()

	newClusterName := defaultMYSQLClusterInfoClusterName
	err := mysqlClusterInfo.Set(map[string]interface{}{"ClusterName": newClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newClusterName, mysqlClusterInfo.ClusterName, "test Set() failed")

	newMiddlewareClusterID := defaultMYSQLClusterInfoMiddlewareClusterID
	err = mysqlClusterInfo.Set(map[string]interface{}{"MiddlewareClusterID": newMiddlewareClusterID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newMiddlewareClusterID, mysqlClusterInfo.MiddlewareClusterID, "test Set() failed")

	newMonitorSystemID := defaultMYSQLClusterInfoMonitorSystemID
	err = mysqlClusterInfo.Set(map[string]interface{}{"MonitorSystemID": newMonitorSystemID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newMonitorSystemID, mysqlClusterInfo.MonitorSystemID, "test Set() failed")

	newOwnerID := defaultMYSQLClusterInfoOwnerID
	err = mysqlClusterInfo.Set(map[string]interface{}{"OwnerID": newOwnerID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newOwnerID, mysqlClusterInfo.OwnerID, "test Set() failed")

	newOwnerGroup := defaultMYSQLClusterInfoOwnerGroup
	err = mysqlClusterInfo.Set(map[string]interface{}{"OwnerGroup": newOwnerGroup})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newOwnerGroup, mysqlClusterInfo.OwnerGroup, "test Set() failed")

	newEnvID := defaultMYSQLClusterInfoEnvID
	err = mysqlClusterInfo.Set(map[string]interface{}{"EnvID": newEnvID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newEnvID, mysqlClusterInfo.EnvID, "test Set() failed")
}

func TestMYSQLClusterInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMYSQLClusterInfo()
	mysqlClusterInfo.Delete()
	asst.True(mysqlClusterInfo.IsDeleted(), "test Delete() failed")
}

func TestMYSQLClusterInfo_MarshalJSON(t *testing.T) {
	var mysqlClusterInfoUnmarshal *MYSQLClusterInfo

	asst := assert.New(t)

	mysqlClusterInfo := initNewMYSQLClusterInfo()
	data, err := mysqlClusterInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &mysqlClusterInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(equalMYSQLClusterInfo(mysqlClusterInfo, mysqlClusterInfoUnmarshal), "test MarshalJSON() failed")
}

func TestMYSQLClusterInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMYSQLClusterInfo()
	data, err := mysqlClusterInfo.MarshalJSONWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{clusterNameJSON: "test"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
