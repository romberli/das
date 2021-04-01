package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestMiddlewareClusterServiceAll(t *testing.T) {
	TestMiddlewareClusterService_GetEntities(t)
	TestMiddlewareClusterService_GetAll(t)
	TestMiddlewareClusterService_GetByID(t)
	TestMiddlewareClusterService_Create(t)
	TestMiddlewareClusterService_Update(t)
	TestMiddlewareClusterService_Delete(t)
	TestMiddlewareClusterService_Marshal(t)
	TestMiddlewareClusterService_MarshalWithFields(t)
}

func TestMiddlewareClusterService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestMiddlewareClusterService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestMiddlewareClusterService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err := s.GetByID("3")
	asst.Nil(err, "test GetByID() failed")
	id := s.Entities[constant.ZeroInt].Identity()
	asst.Equal("3", id, "test GetByID() failed")
}

func TestMiddlewareClusterService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err := s.Create(map[string]interface{}{
		middlewareClusterNameStruct:  defaultMiddlewareClusterInfoClusterName,
		middlewareClusterEnvIDStruct: defaultMiddlewareClusterInfoEnvID,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMiddlewareClusterByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMiddlewareClusterService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{
		middlewareClusterNameStruct: newMiddlewareClusterName,
	})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	middlewareClusterName, err := s.GetEntities()[constant.ZeroInt].Get(middlewareClusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newMiddlewareClusterName, middlewareClusterName)
	// delete
	err = deleteMiddlewareClusterByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMiddlewareClusterService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMiddlewareClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMiddlewareClusterService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*MiddlewareClusterInfo

	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
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
		asst.True(middlewareClusterStructEqual(entity.(*MiddlewareClusterInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestMiddlewareClusterService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(middlewareClusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(middlewareClusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteMiddlewareClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
