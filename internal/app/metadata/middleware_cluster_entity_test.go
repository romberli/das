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
	defaultMiddlewareClusterInfoID                   = 1
	defaultMiddlewareClusterInfoClusterName          = "ttt"
	defaultMiddlewareClusterInfoOwnerID              = 1
	defaultMiddlewareClusterInfoEnvID                = 1
	defaultMiddlewareClusterInfoDelFlag              = 0
	defaultMiddlewareClusterInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultMiddlewareClusterInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"

	middlewareServerIDListCount = 0
	middlewareClusterNameJSON   = "cluster_name"
)

func initNewMiddlewareClusterInfo() *MiddlewareClusterInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultMiddlewareClusterInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultMiddlewareClusterInfoLastUpdateTimeString)
	return NewMiddlewareClusterInfo(
		middlewareClusterRepo,
		defaultMiddlewareClusterInfoID,
		defaultMiddlewareClusterInfoClusterName,
		defaultMiddlewareClusterInfoOwnerID,
		defaultMiddlewareClusterInfoEnvID,
		defaultMiddlewareClusterInfoDelFlag,
		createTime,
		lastUpdateTime,
	)
}

func middlewareClusterStructEqual(a, b *MiddlewareClusterInfo) bool {
	return a.ID == b.ID && a.ClusterName == b.ClusterName && a.OwnerID == b.OwnerID && a.EnvID == b.EnvID && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime
}

func TestMiddlewareClusterEntityAll(t *testing.T) {
	TestMiddlewareClusterInfo_Identity(t)
	TestMiddlewareClusterInfo_GetClusterName(t)
	TestMiddlewareClusterInfo_GetOwnerID(t)
	TestMiddlewareClusterInfo_GetEnvID(t)
	TestMiddlewareClusterInfo_GetDelFlag(t)
	TestMiddlewareClusterInfo_GetCreateTime(t)
	TestMiddlewareClusterInfo_GetLastUpdateTime(t)
	TestMiddlewareClusterInfo_GetCreateTime(t)
	TestMiddlewareClusterInfo_GetMiddlewareServerIDList(t)
	TestMiddlewareClusterInfo_Set(t)
	TestMiddlewareClusterInfo_Delete(t)
	TestMiddlewareClusterInfo_MarshalJSON(t)
	TestMiddlewareClusterInfo_MarshalJSONWithFields(t)
}

func TestMiddlewareClusterInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	asst.Equal(defaultMiddlewareClusterInfoID, middlewareClusterInfo.Identity(), "test Identity() failed")
}

func TestMiddlewareClusterInfo_GetClusterName(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	asst.Equal(defaultMiddlewareClusterInfoClusterName, middlewareClusterInfo.GetClusterName(), "test GetClusterName() failed")
}

func TestMiddlewareClusterInfo_GetOwnerID(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	asst.Equal(defaultMiddlewareClusterInfoOwnerID, middlewareClusterInfo.GetOwnerID(), "test GetOwnerID() failed")
}

func TestMiddlewareClusterInfo_GetEnvID(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	asst.Equal(defaultMiddlewareClusterInfoEnvID, middlewareClusterInfo.GetEnvID(), "test GetEnvID() failed")
}
func TestMiddlewareClusterInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	asst.Equal(constant.ZeroInt, middlewareClusterInfo.GetDelFlag(), "test GetDelFlag() failed")
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

func TestMiddlewareClusterInfo_GetMiddlewareServerIDList(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	list, err := middlewareClusterInfo.GetMiddlewareServerIDList()
	count := len(list)
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareServerIDList() failed", err))
	asst.Equal(middlewareServerIDListCount, count, "test GetMiddlewareServerIDList() failed")
}

func TestMiddlewareClusterInfo_Set(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	newClusterName := "new_cluster"
	err := middlewareClusterInfo.Set(map[string]interface{}{"ClusterName": newClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newClusterName, middlewareClusterInfo.ClusterName, "test Set() failed")
}

func TestMiddlewareClusterInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	middlewareClusterInfo.Delete()
	asst.Equal(1, middlewareClusterInfo.GetDelFlag(), "test Delete() failed")
}

func TestMiddlewareClusterInfo_MarshalJSON(t *testing.T) {
	var middlewareClusterInfoUnmarshal *MiddlewareClusterInfo

	asst := assert.New(t)

	middlewareClusterInfo := initNewMiddlewareClusterInfo()
	data, err := middlewareClusterInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &middlewareClusterInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(middlewareClusterStructEqual(middlewareClusterInfo, middlewareClusterInfoUnmarshal), "test MarshalJSON() failed")
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
