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
	defaultDbInfoID                   = 1
	defaultDbInfoDbName               = "das1"
	defaultDbInfoClusterId            = 1
	defaultDbInfoClusterType          = 1
	defaultDbInfoOwnerId              = "1"
	defaultDbInfoOwnerGroup           = "2,5,6"
	defaultDbInfoEnvId                = "2"
	defaultDbInfoDelFlag              = 0
	defaultDbInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultDbInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	dbNameJSON                        = "db_name"
)

func initNewDbInfo() *DbInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultDbInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultDbInfoLastUpdateTimeString)
	return NewDbInfoWithGlobal(defaultDbInfoID, defaultDbInfoDbName, defaultDbInfoClusterId, defaultDbInfoClusterType, defaultDbInfoOwnerId, defaultDbInfoOwnerGroup, defaultDbInfoEnvId, defaultDbInfoDelFlag, createTime, lastUpdateTime)
}

func dbEqual(a, b *DbInfo) bool {
	return a.ID == b.ID && a.DbName == b.DbName && a.ClusterId == b.ClusterId && a.ClusterType == b.ClusterType && a.OwnerId == b.OwnerId && a.OwnerGroup == b.OwnerGroup && a.EnvId == b.EnvId && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime
}

func TestDbEntityAll(t *testing.T) {
	TestDbInfo_Identity(t)
	TestDbInfo_IsDeleted(t)
	TestDbInfo_GetCreateTime(t)
	TestDbInfo_GetLastUpdateTime(t)
	TestDbInfo_Get(t)
	TestDbInfo_Set(t)
	TestDbInfo_Delete(t)
	TestDbInfo_MarshalJSON(t)
	TestDbInfo_MarshalJSONWithFields(t)
}

func TestDbInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDbInfo()
	asst.Equal(strconv.Itoa(defaultDbInfoID), dbInfo.Identity(), "test Identity() failed")
}

func TestDbInfo_IsDeleted(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDbInfo()
	asst.False(dbInfo.IsDeleted(), "test IsDeleted() failed")
}

func TestDbInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDbInfo()
	asst.True(reflect.DeepEqual(dbInfo.CreateTime, dbInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestDbInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDbInfo()
	asst.True(reflect.DeepEqual(dbInfo.LastUpdateTime, dbInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestDbInfo_Get(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDbInfo()
	dbName, err := dbInfo.Get(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(dbInfo.DbName, dbName, "test Get() failed")
}

func TestDbInfo_Set(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDbInfo()
	newDbName := "new_db"
	err := dbInfo.Set(map[string]interface{}{"DbName": newDbName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(newDbName, dbInfo.DbName, "test Set() failed")
}

func TestDbInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDbInfo()
	dbInfo.Delete()
	asst.True(dbInfo.IsDeleted(), "test Delete() failed")
}

func TestDbInfo_MarshalJSON(t *testing.T) {
	var dbInfoUnmarshal *DbInfo

	asst := assert.New(t)

	dbInfo := initNewDbInfo()
	data, err := dbInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &dbInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(dbEqual(dbInfo, dbInfoUnmarshal), "test MarshalJSON() failed")
}

func TestDbInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDbInfo()
	data, err := dbInfo.MarshalJSONWithFields(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{dbNameJSON: "das1"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
