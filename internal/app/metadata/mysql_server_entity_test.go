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
	defaultMYSQLServerInfoID                   = 1
	defaultMYSQLServerInfoClusterID            = 1
	defaultMYSQLServerInfoHostIP               = "127.0.01"
	defaultMYSQLServerInfoPortNum              = 3306
	defaultMYSQLServerInfoDeploymentType       = 1
	defaultMYSQLServerInfoVersion              = "1.1.1"
	defaultMYSQLServerInfoDelFlag              = 0
	defaultMYSQLServerInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultMYSQLServerInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	hostIPJSON                                 = "host_ip"
	portNumJSON                                = "port_num"
)

func initNewMYSQLServerInfo() *MYSQLServerInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultMYSQLServerInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultMYSQLServerInfoLastUpdateTimeString)
	return NewMYSQLServerInfoWithGlobal(
		defaultMYSQLServerInfoID,
		defaultMYSQLServerInfoClusterID,
		defaultMYSQLServerInfoHostIP,
		defaultMYSQLServerInfoPortNum,
		defaultMYSQLServerInfoDeploymentType,
		defaultMYSQLServerInfoVersion,
		defaultMYSQLServerInfoDelFlag,
		createTime,
		lastUpdateTime)
}

func equalMYSQLServerInfo(a, b *MYSQLServerInfo) bool {
	return a.ID == b.ID &&
		a.ClusterID == b.ClusterID &&
		a.HostIP == b.HostIP &&
		a.PortNum == b.PortNum &&
		a.DeploymentType == b.DeploymentType &&
		a.Version == b.Version &&
		a.DelFlag == b.DelFlag &&
		a.CreateTime == b.CreateTime &&
		a.LastUpdateTime == b.LastUpdateTime
}

func TestMYSQLServerEntityAll(t *testing.T) {
	TestMYSQLServerInfo_Identity(t)
	TestMYSQLServerInfo_IsDeleted(t)
	TestMYSQLServerInfo_GetCreateTime(t)
	TestMYSQLServerInfo_GetLastUpdateTime(t)
	TestMYSQLServerInfo_Get(t)
	TestMYSQLServerInfo_Set(t)
	TestMYSQLServerInfo_Delete(t)
	TestMYSQLServerInfo_MarshalJSON(t)
	TestMYSQLServerInfo_MarshalJSONWithFields(t)
}

func TestMYSQLServerInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMYSQLServerInfo()
	asst.Equal(strconv.Itoa(defaultMYSQLServerInfoID), mysqlServerInfo.Identity(), "test Identity() failed")
}

func TestMYSQLServerInfo_IsDeleted(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMYSQLServerInfo()
	asst.False(mysqlServerInfo.IsDeleted(), "test IsDeleted() failed")
}

func TestMYSQLServerInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMYSQLServerInfo()
	asst.True(reflect.DeepEqual(mysqlServerInfo.CreateTime, mysqlServerInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestMYSQLServerInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMYSQLServerInfo()
	asst.True(reflect.DeepEqual(mysqlServerInfo.LastUpdateTime, mysqlServerInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMYSQLServerInfo_Get(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMYSQLServerInfo()

	clusterID, err := mysqlServerInfo.Get(clusterIDStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mysqlServerInfo.ClusterID, clusterID, "test Get() failed")

	hostIP, err := mysqlServerInfo.Get(hostIPStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mysqlServerInfo.HostIP, hostIP, "test Get() failed")

	portNum, err := mysqlServerInfo.Get(mSPortNumStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mysqlServerInfo.PortNum, portNum, "test Get() failed")

	deploymentType, err := mysqlServerInfo.Get(deploymentTypeStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mysqlServerInfo.DeploymentType, deploymentType, "test Get() failed")

	version, err := mysqlServerInfo.Get(versionStruct)
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(mysqlServerInfo.Version, version, "test Get() failed")
}

func TestMYSQLServerInfo_Set(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMYSQLServerInfo()

	newClusterID := defaultMYSQLServerInfoClusterID
	err := mysqlServerInfo.Set(map[string]interface{}{"ClusterID": newClusterID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newClusterID, mysqlServerInfo.ClusterID, "test Set() failed")

	newHostIP := defaultMYSQLServerInfoHostIP
	err = mysqlServerInfo.Set(map[string]interface{}{"HostIP": newHostIP})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newHostIP, mysqlServerInfo.HostIP, "test Set() failed")

	newPortNum := defaultMYSQLServerInfoPortNum
	err = mysqlServerInfo.Set(map[string]interface{}{"PortNum": newPortNum})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newPortNum, mysqlServerInfo.PortNum, "test Set() failed")

	newDeploymentType := defaultMYSQLServerInfoDeploymentType
	err = mysqlServerInfo.Set(map[string]interface{}{"DeploymentType": newDeploymentType})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newDeploymentType, mysqlServerInfo.DeploymentType, "test Set() failed")

	newVersion := defaultMYSQLServerInfoVersion
	err = mysqlServerInfo.Set(map[string]interface{}{"Version": newVersion})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newVersion, mysqlServerInfo.Version, "test Set() failed")
}

func TestMYSQLServerInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMYSQLServerInfo()
	mysqlServerInfo.Delete()
	asst.True(mysqlServerInfo.IsDeleted(), "test Delete() failed")
}

func TestMYSQLServerInfo_MarshalJSON(t *testing.T) {
	var mysqlServerInfoUnmarshal *MYSQLServerInfo

	asst := assert.New(t)

	mysqlServerInfo := initNewMYSQLServerInfo()
	data, err := mysqlServerInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &mysqlServerInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(equalMYSQLServerInfo(mysqlServerInfo, mysqlServerInfoUnmarshal), "test MarshalJSON() failed")
}

func TestMYSQLServerInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMYSQLServerInfo()
	data, err := mysqlServerInfo.MarshalJSONWithFields(hostIPStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{hostIPJSON: defaultMYSQLServerInfoHostIP})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
