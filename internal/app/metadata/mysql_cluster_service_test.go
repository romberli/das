package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestMySQLClusterServiceAll(t *testing.T) {
	TestMySQLClusterService_GetMySQLServers(t)
	TestMySQLClusterService_GetAll(t)
	TestMySQLClusterService_GetByID(t)
	TestMySQLClusterService_Create(t)
	TestMySQLClusterService_Update(t)
	TestMySQLClusterService_Delete(t)
	TestMySQLClusterService_Marshal(t)
	TestMySQLClusterService_MarshalWithFields(t)
}

func TestMySQLClusterService_GetMySQLServers(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLClusterService(mysqlClusterRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetMySQLClusters()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestMySQLClusterService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLClusterService(mysqlClusterRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetMySQLClusters()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestMySQLClusterService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLClusterService(mysqlClusterRepo)
	err := s.GetByID(1)
	asst.Nil(err, "test GetByID() failed")
	id := s.MySQLClusters[constant.ZeroInt].Identity()
	asst.Equal("1", id, "test GetByID() failed")
}

func TestMySQLClusterService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLClusterService(mysqlClusterRepo)

	err := s.Create(map[string]interface{}{
		clusterNameStruct:         testInsertClusterName,
		middlewareClusterIDStruct: defaultMySQLClusterInfoMiddlewareClusterID,
		monitorSystemIDStruct:     defaultMySQLClusterInfoMonitorSystemID,
		ownerIDStruct:             defaultMySQLClusterInfoOwnerID,
		envIDStruct:               defaultMySQLClusterInfoEnvID,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMySQLClusterByID(s.MySQLClusters[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMySQLClusterService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewMySQLClusterService(mysqlClusterRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{clusterNameStruct: testUpdateClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	mysqlClusterName := s.GetMySQLClusters()[constant.ZeroInt].GetClusterName()
	asst.Equal(testUpdateClusterName, mysqlClusterName)
	// delete
	err = deleteMySQLClusterByID(s.MySQLClusters[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMySQLClusterService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewMySQLClusterService(mysqlClusterRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMySQLClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMySQLClusterService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*MySQLClusterInfo

	asst := assert.New(t)

	s := NewMySQLClusterService(mysqlClusterRepo)
	err := s.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	data, err := s.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	err = json.Unmarshal(data, &entitiesUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	entities := s.GetMySQLClusters()
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		entityUnmarshal := entitiesUnmarshal[i]
		asst.True(equalMySQLClusterInfo(entity.(*MySQLClusterInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestMySQLClusterService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewMySQLClusterService(mysqlClusterRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteMySQLClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
