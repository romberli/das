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
	defaultMiddlewareServerInfoID                   = 1
	defaultMiddlewareServerInfoClusterID            = 1
	defaultMiddlewareServerInfoServerName           = "test"
	defaultMiddlewareServerInfoMiddlewareRole       = 1
	defaultMiddlewareServerInfoSHostIP              = "xxxxx"
	defaultMiddlewareServerInfoPortNum              = 1
	defaultMiddlewareServerInfoDelFlag              = 0
	defaultMiddlewareServerInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultMiddlewareServerInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	middlewareServerNameJSON                        = "server_name"
)

func initNewMiddlewareServerInfo() *MiddlewareServerInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultMiddlewareServerInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultMiddlewareServerInfoLastUpdateTimeString)
	return NewMiddlewareServerInfoWithGlobal(
		defaultMiddlewareServerInfoID,
		defaultMiddlewareServerInfoClusterID,
		defaultMiddlewareServerInfoServerName,
		defaultMiddlewareServerInfoMiddlewareRole,
		defaultMiddlewareServerInfoSHostIP,
		defaultMiddlewareServerInfoPortNum,
		defaultMiddlewareServerInfoDelFlag,
		createTime,
		lastUpdateTime)
}

func middlewareServerStuctEqual(a, b *MiddlewareServerInfo) bool {
	return a.ID == b.ID && a.ClusterID == b.ClusterID && a.ServerName == b.ServerName && a.MiddlewareRole == b.MiddlewareRole && a.HostIP == b.HostIP && a.PortNum == b.PortNum && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime
}

func TestMiddlewareServerEntityAll(t *testing.T) {
	TestMiddlewareServerInfo_Identity(t)
	TestMiddlewareServerInfo_IsDeleted(t)
	TestMiddlewareServerInfo_GetCreateTime(t)
	TestMiddlewareServerInfo_GetLastUpdateTime(t)
	TestMiddlewareServerInfo_Get(t)
	TestMiddlewareServerInfo_Set(t)
	TestMiddlewareServerInfo_Delete(t)
	TestMiddlewareServerInfo_MarshalJSON(t)
	TestMiddlewareServerInfo_MarshalJSONWithFields(t)
}

func TestMiddlewareServerInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	asst.Equal(strconv.Itoa(defaultMiddlewareServerInfoID), middlewareServerInfo.Identity(), "test Identity() failed")
}

func TestMiddlewareServerInfo_IsDeleted(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	asst.False(middlewareServerInfo.IsDeleted(), "test IsDeleted() failed")
}

func TestMiddlewareServerInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	asst.True(reflect.DeepEqual(middlewareServerInfo.CreateTime, middlewareServerInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestMiddlewareServerInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	asst.True(reflect.DeepEqual(middlewareServerInfo.LastUpdateTime, middlewareServerInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMiddlewareServerInfo_Get(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	middlewareServerName, err := middlewareServerInfo.Get(middlewareServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(middlewareServerInfo.ServerName, middlewareServerName, "test Get() failed")
}

func TestMiddlewareServerInfo_Set(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	newServerName := "new_cluster"
	err := middlewareServerInfo.Set(map[string]interface{}{"ServerName": newServerName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(newServerName, middlewareServerInfo.ServerName, "test Set() failed")
}

func TestMiddlewareServerInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	middlewareServerInfo.Delete()
	asst.True(middlewareServerInfo.IsDeleted(), "test Delete() failed")
}

func TestMiddlewareServerInfo_MarshalJSON(t *testing.T) {
	var middlewareServerInfoUnmarshal *MiddlewareServerInfo

	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	data, err := middlewareServerInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &middlewareServerInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(middlewareServerStuctEqual(middlewareServerInfo, middlewareServerInfoUnmarshal), "test MarshalJSON() failed")
}

func TestMiddlewareServerInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	data, err := middlewareServerInfo.MarshalJSONWithFields(middlewareServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{middlewareServerNameJSON: "test"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
