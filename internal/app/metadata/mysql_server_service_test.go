package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestMySQLServerServiceAll(t *testing.T) {
	TestMySQLServerService_GetMySQLServers(t)
	TestMySQLServerService_GetAll(t)
	TestMySQLServerService_GetByClusterID(t)
	TestMySQLServerService_GetByID(t)
	TestMySQLServerService_GetByHostInfo(t)
	TestMySQLServerService_Create(t)
	TestMySQLServerService_Update(t)
	TestMySQLServerService_Delete(t)
	TestMySQLServerService_Marshal(t)
	TestMySQLServerService_MarshalWithFields(t)
}

func TestMySQLServerService_GetMySQLServers(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetMySQLServers()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestMySQLServerService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetMySQLServers()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestMySQLServerService_GetByClusterID(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
	err := s.GetByClusterID(testInitClusterID)
	asst.Nil(err, "test GetByClusterID() failed")
	clusterID := s.MySQLServers[constant.ZeroInt].GetClusterID()
	asst.Equal(testInitClusterID, clusterID, "test GetByClusterID() failed")
}

func TestMySQLServerService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
	err := s.GetByID(testInitServerID)
	asst.Nil(err, "test GetByID() failed")
	id := s.MySQLServers[constant.ZeroInt].Identity()
	asst.Equal(testInitServerID, id, "test GetByID() failed")
}

func TestMySQLServerService_GetByHostInfo(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
	err := s.GetByHostInfo(testInitHostIP, testInitPortNum)
	asst.Nil(err, "test GetByHostInfo() failed")
	hostIP := s.MySQLServers[constant.ZeroInt].GetHostIP()
	asst.Equal(testInitHostIP, hostIP, "test GetByHostInfo() failed")
	portNum := s.MySQLServers[constant.ZeroInt].GetPortNum()
	asst.Equal(testInitPortNum, portNum, "test GetByHostInfo() failed")
}

func TestMySQLServerService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
	err := s.Create(map[string]interface{}{
		clusterIDStruct:      defaultMySQLServerInfoClusterID,
		serverNameStruct:     defaultMySQLServerInfoServerName,
		serviceNameStruct:    defaultMySQLServerInfoServiceName,
		hostIPStruct:         testInsertHostIP,
		portNumStruct:        testInitPortNum,
		deploymentTypeStruct: defaultMySQLServerInfoDeploymentType,
		versionStruct:        defaultMySQLServerInfoVersion})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMySQLServerByID(s.MySQLServers[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMySQLServerService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewMySQLServerService(mysqlServerRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{
		hostIPStruct:  testUpdateHostIP,
		portNumStruct: testUpdatePortNum})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	hostIP := s.GetMySQLServers()[constant.ZeroInt].GetHostIP()
	asst.Equal(testUpdateHostIP, hostIP)
	portNum := s.GetMySQLServers()[constant.ZeroInt].GetPortNum()
	asst.Equal(testUpdatePortNum, portNum)
	// delete
	err = deleteMySQLServerByID(s.MySQLServers[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMySQLServerService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewMySQLServerService(mysqlServerRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMySQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMySQLServerService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*MySQLServerInfo

	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
	err := s.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	data, err := s.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	err = json.Unmarshal(data, &entitiesUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	entities := s.GetMySQLServers()
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		entityUnmarshal := entitiesUnmarshal[i]
		asst.True(equalMySQLServerInfo(entity.(*MySQLServerInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestMySQLServerService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewMySQLServerService(mysqlServerRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteMySQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
