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
	defaultEnvInfoID                   = 2
	defaultEnvInfoEnvName              = "test2"
	defaultEnvInfoDelFlag              = 0
	defaultEnvInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultEnvInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	envNameJSON                        = "env_name"
)

func initNewEnvInfo() *EnvInfo {
	// now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultEnvInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultEnvInfoLastUpdateTimeString)
	return NewEnvInfoWithGlobal(defaultEnvInfoID, defaultEnvInfoEnvName, defaultEnvInfoDelFlag, createTime, lastUpdateTime)
}

func equal(a, b *EnvInfo) bool {
	return a.ID == b.ID && a.EnvName == b.EnvName && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime
}

func TestEnvEntityAll(t *testing.T) {
	TestEnvInfo_Identity(t)
	// TestEnvInfo_IsDeleted(t)
	TestEnvInfo_GetCreateTime(t)
	TestEnvInfo_GetLastUpdateTime(t)
	TestEnvInfo_Get(t)
	TestEnvInfo_Set(t)
	TestEnvInfo_Delete(t)
	TestEnvInfo_MarshalJSON(t)
	TestEnvInfo_MarshalJSONWithFields(t)
	TestEnvInfo_GetEnvName(t)
	TestEnvInfo_GetDelFlag(t)

}

func TestEnvInfo_GetEnvName(t *testing.T) {
	asst := assert.New(t)

	envInfo := initNewEnvInfo()
	asst.Equal(defaultEnvInfoEnvName, envInfo.GetEnvName(), "test GetEnvName failed")
}

func TestEnvInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	envInfo := initNewEnvInfo()
	asst.Equal(defaultEnvInfoDelFlag, envInfo.GetDelFlag(), "test GetEnvName failed")
}
func TestEnvInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	envInfo := initNewEnvInfo()
	asst.Equal(defaultEnvInfoID, envInfo.Identity(), "test Identity() failed")
}

// func TestEnvInfo_IsDeleted(t *testing.T) {
// 	asst := assert.New(t)

// 	envInfo := initNewEnvInfo()
// 	asst.False(envInfo.IsDeleted(), "test IsDeleted() failed")
// }

func TestEnvInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	envInfo := initNewEnvInfo()
	asst.True(reflect.DeepEqual(envInfo.CreateTime, envInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestEnvInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	envInfo := initNewEnvInfo()
	asst.True(reflect.DeepEqual(envInfo.LastUpdateTime, envInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestEnvInfo_Get(t *testing.T) {
	asst := assert.New(t)

	envInfo := initNewEnvInfo()
	envName, err := envInfo.Get(envNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(envInfo.EnvName, envName, "test Get() failed")
}

func TestEnvInfo_Set(t *testing.T) {
	asst := assert.New(t)

	envInfo := initNewEnvInfo()
	newEnvName := "new_env"
	err := envInfo.Set(map[string]interface{}{"EnvName": newEnvName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(newEnvName, envInfo.EnvName, "test Set() failed")
}

func TestEnvInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	envInfo := initNewEnvInfo()
	envInfo.Delete()
	DF := envInfo.GetDelFlag()
	asst.True(DF != constant.ZeroInt, "test Delete() failed")
}

func TestEnvInfo_MarshalJSON(t *testing.T) {
	var envInfoUnmarshal *EnvInfo

	asst := assert.New(t)

	envInfo := initNewEnvInfo()
	data, err := envInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &envInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(equal(envInfo, envInfoUnmarshal), "test MarshalJSON() failed")
}

func TestEnvInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	envInfo := initNewEnvInfo()
	data, err := envInfo.MarshalJSONWithFields(envNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{envNameJSON: "test2"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
