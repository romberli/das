package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestMiddlewareServerServiceAll(t *testing.T) {
	TestMiddlewareServerService_GetMiddlewareServers(t)
	TestMiddlewareServerService_GetAll(t)
	TestMiddlewareServerService_GetByClusterID(t)
	TestMiddlewareServerService_GetByID(t)
	TestMiddlewareServerService_GetByHostInfo(t)
	TestMiddlewareServerService_Create(t)
	TestMiddlewareServerService_Update(t)
	TestMiddlewareServerService_Delete(t)
	TestMiddlewareServerService_Marshal(t)
	TestMiddlewareServerService_MarshalWithFields(t)
}

func TestMiddlewareServerService_GetMiddlewareServers(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetMiddlewareServers() failed")
	entities := s.GetMiddlewareServers()
	asst.Greater(len(entities), constant.ZeroInt, "test GetMiddlewareServers() failed")
}

func TestMiddlewareServerService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetAll() failed")
	entities := s.GetMiddlewareServers()
	asst.Greater(len(entities), constant.ZeroInt, "test GetAll() failed")
}

func TestMiddlewareServerService_GetByClusterID(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.GetByClusterID(13)
	asst.Nil(err, "test GetByClusterID() failed")
	clusterID := s.MiddlewareServers[constant.ZeroInt].GetClusterID()
	asst.Equal(13, clusterID, "test GetByClusterID() failed")
}

func TestMiddlewareServerService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.GetByID(1)
	asst.Nil(err, "test GetByID() failed")
	id := s.MiddlewareServers[constant.ZeroInt].Identity()
	asst.Equal(1, id, "test GetByID() failed")
}

func TestMiddlewareServerService_GetByHostInfo(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.GetByHostInfo("1", 1)
	asst.Nil(err, "test GetByHostInfo() failed")
	id := s.MiddlewareServers[constant.ZeroInt].Identity()
	asst.Equal(1, id, "test GetByHostInfo() failed")
}

func TestMiddlewareServerService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.Create(map[string]interface{}{
		middlewareServerClusterIDStruct:      defaultMiddlewareServerInfoClusterID,
		middlewareServerNameStruct:           defaultMiddlewareServerInfoServerName,
		middlewareServerMiddlewareRoleStruct: defaultMiddlewareServerInfoMiddlewareRole,
		middlewareServerHostIPStruct:         defaultMiddlewareServerInfoHostIP,
		middlewareServerPortNumStruct:        defaultMiddlewareServerInfoPortNum,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMiddlewareServerByID(s.MiddlewareServers[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMiddlewareServerService_Update(t *testing.T) {
	asst := assert.New(t)
	entity, err := createMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewMiddlewareServerService(middlewareServerRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{
		middlewareServerNameStruct: newMiddlewareServerName,
	})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	middlewareServerName := s.GetMiddlewareServers()[constant.ZeroInt].GetServerName()
	asst.Equal(newMiddlewareServerName, middlewareServerName)
	// delete
	err = deleteMiddlewareServerByID(s.MiddlewareServers[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMiddlewareServerService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewMiddlewareServerService(middlewareServerRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMiddlewareServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMiddlewareServerService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*MiddlewareServerInfo

	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	data, err := s.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	err = json.Unmarshal(data, &entitiesUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	entities := s.GetMiddlewareServers()
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		entityUnmarshal := entitiesUnmarshal[i]
		asst.True(middlewareServerStructEqual(entity.(*MiddlewareServerInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestMiddlewareServerService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewMiddlewareServerService(middlewareServerRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(middlewareServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(middlewareServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteMiddlewareServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
