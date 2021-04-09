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
	TestMiddlewareClusterService_GetMiddlewareClusters(t)
	TestMiddlewareClusterService_GetAll(t)
	TestMiddlewareClusterService_GetByEnv(t)
	TestMiddlewareClusterService_GetByID(t)
	TestMiddlewareClusterService_GetByName(t)
	TestMiddlewareClusterService_GetMiddlewareServerIDList(t)
	TestMiddlewareClusterService_Create(t)
	TestMiddlewareClusterService_Update(t)
	TestMiddlewareClusterService_Delete(t)
	TestMiddlewareClusterService_Marshal(t)
	TestMiddlewareClusterService_MarshalWithFields(t)
}

func TestMiddlewareClusterService_GetMiddlewareClusters(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetMiddlewareClusters() failed")
	entities := s.GetMiddlewareClusters()
	asst.Greater(len(entities), constant.ZeroInt, "test GetMiddlewareClusters() failed")
}

func TestMiddlewareClusterService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetAll() failed")
	entities := s.GetMiddlewareClusters()
	asst.Greater(len(entities), constant.ZeroInt, "test GetAll() failed")
}

func TestMiddlewareClusterService_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err := s.GetByEnv(1)
	asst.Nil(err, "test GetByEnv() failed")
	envID := s.MiddlewareClusters[constant.ZeroInt].GetEnvID()
	asst.Equal(1, envID, "test GetByEnvID() failed")
}

func TestMiddlewareClusterService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err := s.GetByID(8)
	asst.Nil(err, "test GetByID() failed")
	id := s.MiddlewareClusters[constant.ZeroInt].Identity()
	asst.Equal(8, id, "test GetByID() failed")
}

func TestMiddlewareClusterService_GetByName(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err := s.GetByName("test")
	asst.Nil(err, "test GetByName() failed")
	clusterName := s.MiddlewareClusters[constant.ZeroInt].GetClusterName()
	asst.Equal("test", clusterName, "test GetByName() failed")
}

func TestMiddlewareClusterService_GetMiddlewareServerIDList(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
	_, err := s.GetMiddlewareServerIDList(13)
	asst.Nil(err, "test GetMiddlewareServerIDList failed")
	middlewareServerList := s.MiddlewareServerList
	asst.Equal(2, len(middlewareServerList), "test GetMiddlewareServerIDList failed")
}

func TestMiddlewareClusterService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareClusterService(middlewareClusterRepo)
	err := s.Create(map[string]interface{}{
		middlewareClusterNameStruct:    defaultMiddlewareClusterInfoClusterName,
		middlewareClusterOwnerIDStruct: defaultMiddlewareClusterInfoOwnerID,
		middlewareClusterEnvIDStruct:   defaultMiddlewareClusterInfoEnvID,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMiddlewareClusterByID(s.MiddlewareClusters[0].Identity())
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
	middlewareClusterName := s.GetMiddlewareClusters()[constant.ZeroInt].GetClusterName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newMiddlewareClusterName, middlewareClusterName)
	// delete
	err = deleteMiddlewareClusterByID(s.MiddlewareClusters[0].Identity())
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
	entities := s.GetMiddlewareClusters()
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
