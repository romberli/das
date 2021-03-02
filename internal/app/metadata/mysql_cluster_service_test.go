package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestMYSQLClusterServiceAll(t *testing.T) {
	TestMYSQLClusterService_GetEntities(t)
	TestMYSQLClusterService_GetAll(t)
	TestMYSQLClusterService_GetByID(t)
	TestMYSQLClusterService_Create(t)
	TestMYSQLClusterService_Update(t)
	TestMYSQLClusterService_Delete(t)
	TestMYSQLClusterService_Marshal(t)
	TestMYSQLClusterService_MarshalWithFields(t)
}

func TestMYSQLClusterService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	s := NewMYSQLClusterService(mysqlClusterRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMYSQLClusterService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewMYSQLClusterService(mysqlClusterRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMYSQLClusterService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewMYSQLClusterService(mysqlClusterRepo)
	err := s.GetByID("1")
	asst.Nil(err, "test GetByID() failed")
	id := s.Entities[constant.ZeroInt].Identity()
	asst.Equal("1", id, "test GetByID() failed")
}

func TestMYSQLClusterService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMYSQLClusterService(mysqlClusterRepo)
	err := s.Create(map[string]interface{}{clusterNameStruct: testInsertClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMYSQLClusterByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMYSQLClusterService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMYSQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewMYSQLClusterService(mysqlClusterRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{clusterNameStruct: testUpdateClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	mysqlClusterName, err := s.GetEntities()[constant.ZeroInt].Get(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testUpdateClusterName, mysqlClusterName)
	// delete
	err = deleteMYSQLClusterByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMYSQLClusterService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMYSQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewMYSQLClusterService(mysqlClusterRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMYSQLClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMYSQLClusterService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*MYSQLClusterInfo

	asst := assert.New(t)

	s := NewMYSQLClusterService(mysqlClusterRepo)
	err := s.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	data, err := s.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	err = json.Unmarshal(data, &entitiesUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	entities := s.GetEntities()
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		entityUnmarshal := entitiesUnmarshal[i]
		asst.True(equalMYSQLClusterInfo(entity.(*MYSQLClusterInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestMYSQLClusterService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMYSQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewMYSQLClusterService(mysqlClusterRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteMYSQLClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
