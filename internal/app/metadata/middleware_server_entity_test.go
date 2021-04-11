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
	defaultMiddlewareServerInfoID                   = 1
	defaultMiddlewareServerInfoClusterID            = 1
	defaultMiddlewareServerInfoServerName           = "test"
	defaultMiddlewareServerInfoMiddlewareRole       = 1
	defaultMiddlewareServerInfoHostIP               = "xxxxx"
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
	return NewMiddlewareServerInfo(
		middlewareServerRepo,
		defaultMiddlewareServerInfoID,
		defaultMiddlewareServerInfoClusterID,
		defaultMiddlewareServerInfoServerName,
		defaultMiddlewareServerInfoMiddlewareRole,
		defaultMiddlewareServerInfoHostIP,
		defaultMiddlewareServerInfoPortNum,
		defaultMiddlewareServerInfoDelFlag,
		createTime,
		lastUpdateTime,
	)
}

func middlewareServerStructEqual(a, b *MiddlewareServerInfo) bool {
	return a.ID == b.ID && a.ClusterID == b.ClusterID && a.ServerName == b.ServerName && a.MiddlewareRole == b.MiddlewareRole && a.HostIP == b.HostIP && a.PortNum == b.PortNum && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime
}

func TestMiddlewareServerEntityAll(t *testing.T) {
	TestMiddlewareServerInfo_Identity(t)
	TestMiddlewareServerInfo_GetClusterID(t)
	TestMiddlewareServerInfo_GetServerName(t)
	TestMiddlewareServerInfo_GetMiddlewareRole(t)
	TestMiddlewareServerInfo_GetHostIP(t)
	TestMiddlewareServerInfo_GetPortNum(t)
	TestMiddlewareServerInfo_GetDelFlag(t)
	TestMiddlewareServerInfo_GetCreateTime(t)
	TestMiddlewareServerInfo_GetLastUpdateTime(t)
	TestMiddlewareServerInfo_Set(t)
	TestMiddlewareServerInfo_Delete(t)
	TestMiddlewareServerInfo_MarshalJSON(t)
	TestMiddlewareServerInfo_MarshalJSONWithFields(t)
}

func TestMiddlewareServerInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	asst.Equal(defaultMiddlewareServerInfoID, middlewareServerInfo.Identity(), "test Identity() failed")
}

func TestMiddlewareServerInfo_GetClusterID(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	asst.Equal(defaultMiddlewareServerInfoClusterID, middlewareServerInfo.GetClusterID(), "test GetClusterID() failed")
}

func TestMiddlewareServerInfo_GetServerName(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	asst.Equal(defaultMiddlewareServerInfoServerName, middlewareServerInfo.GetServerName(), "test GetServerName() failed")
}

func TestMiddlewareServerInfo_GetMiddlewareRole(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	asst.Equal(defaultMiddlewareServerInfoMiddlewareRole, middlewareServerInfo.GetMiddlewareRole(), "test GetServerName() failed")
}

func TestMiddlewareServerInfo_GetHostIP(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	asst.Equal(defaultMiddlewareServerInfoHostIP, middlewareServerInfo.GetHostIP(), "test GetServerName() failed")
}

func TestMiddlewareServerInfo_GetPortNum(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	asst.Equal(defaultMiddlewareServerInfoPortNum, middlewareServerInfo.GetPortNum(), "test GetServerName() failed")
}

func TestMiddlewareServerInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	asst.Equal(defaultMiddlewareServerInfoDelFlag, middlewareServerInfo.GetDelFlag(), "test GetServerName() failed")
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
	asst.Equal(1, middlewareServerInfo.GetDelFlag(), "test Delete() failed")
}

func TestMiddlewareServerInfo_MarshalJSON(t *testing.T) {
	var middlewareServerInfoUnmarshal *MiddlewareServerInfo

	asst := assert.New(t)

	middlewareServerInfo := initNewMiddlewareServerInfo()
	data, err := middlewareServerInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &middlewareServerInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(middlewareServerStructEqual(middlewareServerInfo, middlewareServerInfoUnmarshal), "test MarshalJSON() failed")
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
