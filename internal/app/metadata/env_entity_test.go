package metadata

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	defaultEnvInfoID                   = 1
	defaultEnvInfoEnvName              = "test"
	defaultEnvInfoDelFlag              = 0
	defaultEnvInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultEnvInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
)

var asst = assert.New(&testing.T{})

func InitNewEnvInfo() *EnvInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultEnvInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultEnvInfoLastUpdateTimeString)
	return NewEnvInfoWithGlobal(defaultEnvInfoID, defaultEnvInfoEnvName, defaultEnvInfoDelFlag, createTime, lastUpdateTime)
}

func TestAll(t *testing.T) {
	TestEnvInfo_Identity(t)
	TestEnvInfo_IsDeleted(t)
	TestEnvInfo_GetCreateTime(t)
	TestEnvInfo_GetLastUpdateTime(t)
	TestEnvInfo_Get(t)
	TestEnvInfo_Set(t)
	TestEnvInfo_Delete(t)
	TestEnvInfo_MarshalJSON(t)
	TestEnvInfo_MarshalJSONWithFields(t)
}

func TestEnvInfo_Identity(t *testing.T) {
	envInfo := InitNewEnvInfo()
	asst.Equal(strconv.Itoa(defaultEnvInfoID), envInfo.Identity(), "test Identity() failed")
}

func TestEnvInfo_IsDeleted(t *testing.T) {
	envInfo := InitNewEnvInfo()
	asst.False(envInfo.IsDeleted(), "test IsDeleted() failed")
}

func TestEnvInfo_GetCreateTime(t *testing.T) {
	envInfo := InitNewEnvInfo()
	asst.True(reflect.DeepEqual(envInfo.CreateTime, envInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestEnvInfo_GetLastUpdateTime(t *testing.T) {
	envInfo := InitNewEnvInfo()
	asst.True(reflect.DeepEqual(envInfo.LastUpdateTime, envInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestEnvInfo_Get(t *testing.T) {
	envInfo := InitNewEnvInfo()
	envName, err := envInfo.Get("env_name")
	asst.Nil(err, "test Get() failed")
	asst.Equal(envInfo.EnvName, envName, "test Get() failed")
}

func TestEnvInfo_Set(t *testing.T) {
	envInfo := InitNewEnvInfo()
	newEnvName := "new_env"
	err := envInfo.Set(map[string]interface{}{"EnvName": newEnvName})
	asst.Nil(err, "test Get() failed")
	asst.Equal(newEnvName, envInfo.EnvName, "test Set() failed")
}

func TestEnvInfo_Delete(t *testing.T) {
	envInfo := InitNewEnvInfo()
	envInfo.Delete()
	asst.True(envInfo.IsDeleted(), "test Delete() failed")
}

func TestEnvInfo_MarshalJSON(t *testing.T) {
	var envInfoUnmarshal *EnvInfo
	envInfo := InitNewEnvInfo()
	data, err := envInfo.MarshalJSON()
	asst.Nil(err, "test MarshalJSON() failed")
	err = json.Unmarshal(data, envInfoUnmarshal)
	asst.Nil(err, "test MarshalJSON() failed")
	asst.True(reflect.DeepEqual(envInfo, envInfoUnmarshal))
}

func TestEnvInfo_MarshalJSONWithFields(t *testing.T) {
	envInfo := InitNewEnvInfo()
	data, err := envInfo.MarshalJSONWithFields("EnvName")
	asst.Nil(err, "test MarshalJSONWithFields() failed")
	expect, err := json.Marshal(map[string]interface{}{"EnvName": "test"})
	asst.Nil(err, "test MarshalJSONWithFields() failed")
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
