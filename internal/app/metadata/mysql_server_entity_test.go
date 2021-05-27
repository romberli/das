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
	defaultMySQLServerInfoID                   = 1
	defaultMySQLServerInfoClusterID            = 1
	defaultMySQLServerInfoServerName           = "server1"
	defaultMySQLServerInfoServiceName          = "service1"
	defaultMySQLServerInfoHostIP               = "127.0.0.1"
	defaultMySQLServerInfoPortNum              = 3306
	defaultMySQLServerInfoDeploymentType       = 1
	defaultMySQLServerInfoVersion              = "1.1.1"
	defaultMySQLServerInfoDelFlag              = 0
	defaultMySQLServerInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultMySQLServerInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	hostIPJSON                                 = "host_ip"
)

func initNewMySQLServerInfo() *MySQLServerInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultMySQLServerInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultMySQLServerInfoLastUpdateTimeString)
	return NewMySQLServerInfoWithGlobal(
		defaultMySQLServerInfoID,
		defaultMySQLServerInfoClusterID,
		defaultMySQLServerInfoServerName,
		defaultMySQLServerInfoServiceName,
		defaultMySQLServerInfoHostIP,
		defaultMySQLServerInfoPortNum,
		defaultMySQLServerInfoDeploymentType,
		defaultMySQLServerInfoVersion,
		defaultMySQLServerInfoDelFlag,
		createTime,
		lastUpdateTime)
}

func equalMySQLServerInfo(a, b *MySQLServerInfo) bool {
	return a.ID == b.ID &&
		a.ClusterID == b.ClusterID &&
		a.ServerName == b.ServerName &&
		a.ServiceName == b.ServiceName &&
		a.HostIP == b.HostIP &&
		a.PortNum == b.PortNum &&
		a.DeploymentType == b.DeploymentType &&
		a.Version == b.Version &&
		a.DelFlag == b.DelFlag &&
		a.CreateTime == b.CreateTime &&
		a.LastUpdateTime == b.LastUpdateTime
}

func TestMySQLServerEntityAll(t *testing.T) {
	TestMySQLServerInfo_Identity(t)
	TestMySQLServerInfo_Get(t)
	TestMySQLServerInfo_Set(t)
	TestMySQLServerInfo_Delete(t)
	TestMySQLServerInfo_MarshalJSON(t)
	TestMySQLServerInfo_MarshalJSONWithFields(t)
}

func TestMySQLServerInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMySQLServerInfo()
	asst.Equal(defaultMySQLServerInfoID, mysqlServerInfo.Identity(), "test Identity() failed")
}

func TestMySQLServerInfo_Get(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMySQLServerInfo()

	clusterID := mysqlServerInfo.GetClusterID()
	asst.Equal(mysqlServerInfo.ClusterID, clusterID, "test GetClusterID() failed")

	hostIP := mysqlServerInfo.GetHostIP()
	asst.Equal(mysqlServerInfo.HostIP, hostIP, "test GetHostIP() failed")

	portNum := mysqlServerInfo.GetPortNum()
	asst.Equal(mysqlServerInfo.PortNum, portNum, "test GetPortNum() failed")

	deploymentType := mysqlServerInfo.GetDeploymentType()
	asst.Equal(mysqlServerInfo.DeploymentType, deploymentType, "test GetDeploymentType() failed")

	version := mysqlServerInfo.GetVersion()
	asst.Equal(mysqlServerInfo.Version, version, "test GetVersion() failed")

	delFlag := mysqlServerInfo.GetDelFlag()
	asst.Equal(mysqlServerInfo.DelFlag, delFlag, "test GetDelFlag() failed")

	createTime := mysqlServerInfo.GetCreateTime()
	asst.True(reflect.DeepEqual(mysqlServerInfo.CreateTime, createTime), "test GetCreateTime() failed")

	lastUpdateTime := mysqlServerInfo.GetLastUpdateTime()
	asst.True(reflect.DeepEqual(mysqlServerInfo.LastUpdateTime, lastUpdateTime), "test GetLastUpdateTime() failed")
}

func TestMySQLServerInfo_Set(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMySQLServerInfo()

	newClusterID := defaultMySQLServerInfoClusterID
	err := mysqlServerInfo.Set(map[string]interface{}{"ClusterID": newClusterID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newClusterID, mysqlServerInfo.ClusterID, "test Set() failed")

	newHostIP := defaultMySQLServerInfoHostIP
	err = mysqlServerInfo.Set(map[string]interface{}{"HostIP": newHostIP})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newHostIP, mysqlServerInfo.HostIP, "test Set() failed")

	newPortNum := defaultMySQLServerInfoPortNum
	err = mysqlServerInfo.Set(map[string]interface{}{"PortNum": newPortNum})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newPortNum, mysqlServerInfo.PortNum, "test Set() failed")

	newDeploymentType := defaultMySQLServerInfoDeploymentType
	err = mysqlServerInfo.Set(map[string]interface{}{"DeploymentType": newDeploymentType})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newDeploymentType, mysqlServerInfo.DeploymentType, "test Set() failed")

	newVersion := defaultMySQLServerInfoVersion
	err = mysqlServerInfo.Set(map[string]interface{}{"Version": newVersion})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newVersion, mysqlServerInfo.Version, "test Set() failed")
}

func TestMySQLServerInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMySQLServerInfo()
	mysqlServerInfo.Delete()
	asst.Equal(1, mysqlServerInfo.GetDelFlag(), "test Delete() failed")
}

func TestMySQLServerInfo_MarshalJSON(t *testing.T) {
	var mysqlServerInfoUnmarshal *MySQLServerInfo

	asst := assert.New(t)

	mysqlServerInfo := initNewMySQLServerInfo()
	data, err := mysqlServerInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &mysqlServerInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(equalMySQLServerInfo(mysqlServerInfo, mysqlServerInfoUnmarshal), "test MarshalJSON() failed")
}

func TestMySQLServerInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	mysqlServerInfo := initNewMySQLServerInfo()
	data, err := mysqlServerInfo.MarshalJSONWithFields(hostIPStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{hostIPJSON: defaultMySQLServerInfoHostIP})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
