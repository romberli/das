package metadata

import (
	"encoding/json"
	"github.com/jinzhu/now"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
)

const (
	defaultMiddlewareClusterInfoID                   = 8
	defaultMiddlewareClusterInfoClusterName          = "ttt"
	defaultMiddlewareClusterInfoEnvID                = 8
	defaultMiddlewareClusterInfoDelFlag              = 0
	defaultMiddlewareClusterInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultMiddlewareClusterInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	middlewareClusterNameJSON                        = "cluster_name"
)

func initNewMiddlewareClusterInfo() *MiddlewareClusterInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultMiddlewareClusterInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultMiddlewareClusterInfoLastUpdateTimeString)
	return NewMiddlewareClusterInfoWithGlobal(defaultMiddlewareClusterInfoID, defaultMiddlewareClusterInfoClusterName, defaultMiddlewareClusterInfoEnvID, defaultMiddlewareClusterInfoDelFlag, createTime, lastUpdateTime)
}

func middlewareClusterStuctEqual(a, b *MiddlewareClusterInfo) bool {
	return a.ID == b.ID && a.ClusterName == b.ClusterName && a.EnvID == b.EnvID && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime
}

func TestMiddlewareClusterEntityAll(t *testing.T) {
	TestMiddlewareClusterInfo_Identity(t)
	TestMiddlewareClusterInfo_IsDeleted(t)
	TestMiddlewareClusterInfo_GetCreateTime(t)
	TestMiddlewareClusterInfo_GetLastUpdateTime(t)
	TestMiddlewareClusterInfo_Get(t)
	TestMiddlewareClusterInfo_Set(t)
	TestMiddlewareClusterInfo_Delete(t)
	TestMiddlewareClusterInfo_MarshalJSON(t)
	TestMiddlewareClusterInfo_MarshalJSONWithFields(t)
}

func TestMiddlewareClusterInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	asst.Equal(strconv.Itoa(defaultMiddlewareClusterInfoID), middlewareClusterInfo.Identity(), "test Identity() failed")
}

func TestMiddlewareClusterInfo_IsDeleted(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	asst.False(middlewareClusterInfo.IsDeleted(), "test IsDeleted() failed")
}

func TestMiddlewareClusterInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	asst.True(reflect.DeepEqual(middlewareClusterInfo.CreateTime, middlewareClusterInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestMiddlewareClusterInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	asst.True(reflect.DeepEqual(middlewareClusterInfo.LastUpdateTime, middlewareClusterInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMiddlewareClusterInfo_Get(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	middlewareClusterName, err := middlewareClusterInfo.Get(middlewareClusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(middlewareClusterInfo.ClusterName, middlewareClusterName, "test Get() failed")
}

func TestMiddlewareClusterInfo_Set(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	newClusterName := "new_cluster"
	err := middlewareClusterInfo.Set(map[string]interface{}{"ClusterName": newClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(newClusterName, middlewareClusterInfo.ClusterName, "test Set() failed")
}

func TestMiddlewareClusterInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	middlewareClusterInfo.Delete()
	asst.True(middlewareClusterInfo.IsDeleted(), "test Delete() failed")
}

func TestMiddlewareClusterInfo_MarshalJSON(t *testing.T) {
	var middlewareClusterInfoUnmarshal *MiddlewareClusterInfo

	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	data, err := middlewareClusterInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &middlewareClusterInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(middlewareClusterStuctEqual(middlewareClusterInfo, middlewareClusterInfoUnmarshal), "test MarshalJSON() failed")
}

func TestMiddlewareClusterInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	data, err := middlewareClusterInfo.MarshalJSONWithFields(middlewareClusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{middlewareClusterNameJSON: "ttt"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
