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
	defaultDBInfoID                   = 1
	defaultDBInfoDBName               = "test"
	defaultDBInfoClusterID            = 1
	defaultDBInfoClusterType          = 1
	defaultDBInfoOwnerID              = 1
	defaultDBInfoEnvID                = 2
	defaultDBInfoDelFlag              = 0
	defaultDBInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultDBInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	dbNameJSON                        = "db_name"
)

func initNewDBInfo() *DBInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultDBInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultDBInfoLastUpdateTimeString)
	return NewDBInfo(dbRepo, defaultDBInfoID, defaultDBInfoDBName, defaultDBInfoClusterID,
		defaultDBInfoClusterType, defaultDBInfoOwnerID, defaultDBInfoEnvID, defaultDBInfoDelFlag,
		createTime, lastUpdateTime)
}

func dbEqual(a, b *DBInfo) bool {
	return a.ID == b.ID && a.DBName == b.DBName && a.ClusterID == b.ClusterID && a.ClusterType == b.ClusterType &&
		a.OwnerID == b.OwnerID && a.EnvID == b.EnvID && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime &&
		a.LastUpdateTime == b.LastUpdateTime
}

func TestDBEntityAll(t *testing.T) {
	TestDBInfo_Identity(t)
	TestDBInfo_GetDBName(t)
	TestDBInfo_GetClusterID(t)
	TestDBInfo_GetClusterType(t)
	TestDBInfo_GetOwnerID(t)
	TestDBInfo_GetEnvID(t)
	TestDBInfo_GetDelFlag(t)
	TestDBInfo_GetCreateTime(t)
	TestDBInfo_GetLastUpdateTime(t)
	TestDBInfo_GetAppIDList(t)
	TestDBInfo_Set(t)
	TestDBInfo_Delete(t)
	TestDBInfo_AddDBApp(t)
	TestDBInfo_DeleteDBApp(t)
	TestDBInfo_MarshalJSON(t)
	TestDBInfo_MarshalJSONWithFields(t)
}

func TestDBInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	asst.Equal(defaultDBInfoID, dbInfo.Identity(), "test Identity() failed")
}

func TestDBInfo_GetDBName(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	asst.Equal(defaultDBInfoDBName, dbInfo.GetDBName(), "test GetDBName() failed")
}

func TestDBInfo_GetClusterID(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	asst.Equal(defaultDBInfoClusterID, dbInfo.GetClusterID(), "test GetClusterID() failed")
}

func TestDBInfo_GetClusterType(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	asst.Equal(defaultDBInfoClusterType, dbInfo.GetClusterType(), "test GetClusterType() failed")
}

func TestDBInfo_GetOwnerID(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	asst.Equal(defaultDBInfoOwnerID, dbInfo.GetOwnerID(), "test GetOwnerID() failed")
}

func TestDBInfo_GetEnvID(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	asst.Equal(defaultDBInfoEnvID, dbInfo.GetEnvID(), "test GetEnvID() failed")
}

func TestDBInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	asst.Equal(defaultDBInfoDelFlag, dbInfo.GetDelFlag(), "test GetDelFlag() failed")
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

func TestDBInfo_GetAppIDList(t *testing.T) {
	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	appIDList, err := dbInfo.GetAppIDList()
	asst.Nil(err, common.CombineMessageWithError("test GetAppIDList() failed", err))
	defaultAppIDList := []int{1, 2}
	for i := 0; i < len(appIDList); i++ {
		asst.Equal(defaultAppIDList[i], appIDList[i], "test GetAppIDList() failed")
	}
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
	asst.Equal(1, dbInfo.DelFlag, "test Delete() failed")
}

func TestDBInfo_AddDBApp(t *testing.T) {
	var appIDList []int

	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	err := dbInfo.AddApp(3)
	appIDList, err = dbInfo.GetAppIDList()
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	asst.Equal(3, len(appIDList), "test AddApp() failed")
	// delete
	err = dbInfo.DeleteApp(3)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
}

func TestDBInfo_DeleteDBApp(t *testing.T) {
	var appIDList []int

	asst := assert.New(t)

	dbInfo := initNewDBInfo()
	err := dbInfo.DeleteApp(2)
	appIDList, err = dbInfo.GetAppIDList()
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	asst.Equal(1, len(appIDList), "test DeleteApp() failed")
	// add
	err = dbInfo.AddApp(2)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
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
	data, err := dbInfo.MarshalJSONWithFields(dbDBNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{dbNameJSON: "test"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
