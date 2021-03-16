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
	defaultDBInfoID                   = 1
	defaultDBInfoDBName               = "das1"
	defaultDBInfoClusterID            = "1"
	defaultDBInfoClusterType          = "1"
	defaultDBInfoOwnerID              = "1"
	defaultDBInfoOwnerGroup           = "2,5,6"
	defaultDBInfoEnvID                = "2"
	defaultDBInfoDelFlag              = 0
	defaultDBInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultDBInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	dbNameJSON                        = "db_name"
)

func initNewDBInfo() *DBInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultDBInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultDBInfoLastUpdateTimeString)
	return NewDBInfoWithGlobal(defaultDBInfoID, defaultDBInfoDBName, defaultDBInfoClusterID, defaultDBInfoClusterType, defaultDBInfoOwnerID, defaultDBInfoOwnerGroup, defaultDBInfoEnvID, defaultDBInfoDelFlag, createTime, lastUpdateTime)
}

func dbEqual(a, b *DBInfo) bool {
	return a.ID == b.ID && a.DBName == b.DBName && a.ClusterID == b.ClusterID && a.ClusterType == b.ClusterType && a.OwnerID == b.OwnerID && a.OwnerGroup == b.OwnerGroup && a.EnvID == b.EnvID && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime
}

func TestDBEntityAll(t *testing.T) {
	TestDBInfo_Identity(t)
	TestDBInfo_IsDeleted(t)
	TestDBInfo_GetCreateTime(t)
	TestDBInfo_GetLastUpdateTime(t)
	TestDBInfo_Get(t)
	TestDBInfo_Set(t)
	TestDBInfo_Delete(t)
	TestDBInfo_MarshalJSON(t)
	TestDBInfo_MarshalJSONWithFields(t)
}

func TestDBInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	asst.Equal(strconv.Itoa(defaultDBInfoID), dbInfo.Identity(), "test Identity() failed")
}

func TestDBInfo_IsDeleted(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	asst.False(dbInfo.IsDeleted(), "test IsDeleted() failed")
}

func TestDBInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	asst.True(reflect.DeepEqual(dbInfo.CreateTime, dbInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestDBInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	asst.True(reflect.DeepEqual(dbInfo.LastUpdateTime, dbInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestDBInfo_Get(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	dbName, err := dbInfo.Get(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(dbInfo.DBName, dbName, "test Get() failed")
}

func TestDBInfo_Set(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	newDBName := "new_db"
	err := dbInfo.Set(map[string]interface{}{"DBName": newDBName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(newDBName, dbInfo.DBName, "test Set() failed")
}

func TestDBInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	dbInfo.Delete()
	asst.True(dbInfo.IsDeleted(), "test Delete() failed")
}

func TestDBInfo_MarshalJSON(t *testing.T) {
	var dbInfoUnmarshal *DBInfo

	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	data, err := dbInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &dbInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(dbEqual(dbInfo, dbInfoUnmarshal), "test MarshalJSON() failed")
}

func TestDBInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	data, err := dbInfo.MarshalJSONWithFields(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{dbNameJSON: "das1"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
