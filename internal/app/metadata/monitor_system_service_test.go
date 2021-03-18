package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestMonitorSystemServiceAll(t *testing.T) {
	TestMonitorSystemService_GetEntities(t)
	TestMonitorSystemService_GetAll(t)
	TestMonitorSystemService_GetByID(t)
	TestMonitorSystemService_Create(t)
	TestMonitorSystemService_Update(t)
	TestMonitorSystemService_Delete(t)
	TestMonitorSystemService_Marshal(t)
	TestMonitorSystemService_MarshalWithFields(t)
}

func TestMonitorSystemService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	s := NewMonitorSystemService(monitorSystemRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMonitorSystemService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewMonitorSystemService(monitorSystemRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMonitorSystemService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewMonitorSystemService(monitorSystemRepo)
	err := s.GetByID("1")
	asst.Nil(err, "test GetByID() failed")
	id := s.Entities[constant.ZeroInt].Identity()
	asst.Equal("1", id, "test GetByID() failed")
}

func TestMonitorSystemService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMonitorSystemService(monitorSystemRepo)
	err := s.Create(map[string]interface{}{monitorSystemNameStruct: defaultMonitorSystemInfoSystemName, monitorSystemTypeStruct: defaultMonitorSystemInfoSystemType, monitorSystemHostIPStruct: defaultMonitorSystemInfoHostIP, monitorSystemPortNumStruct: defaultMonitorSystemInfoPortNum, portNumSlowStruct: defaultMonitorSystemInfoPortNumSlow, baseUrlStruct: defaultMonitorSystemInfoBaseUrl})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMonitorSystemByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMonitorSystemService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewMonitorSystemService(monitorSystemRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{monitorSystemNameStruct: newMonitorSystemName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	monitorSystemName, err := s.GetEntities()[constant.ZeroInt].Get(monitorSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newMonitorSystemName, monitorSystemName)
	// delete
	err = deleteMonitorSystemByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMonitorSystemService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewMonitorSystemService(monitorSystemRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMonitorSystemService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*MonitorSystemInfo

	asst := assert.New(t)

	s := NewMonitorSystemService(monitorSystemRepo)
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
		asst.True(monitorSystemEqual(entity.(*MonitorSystemInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestMonitorSystemService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewMonitorSystemService(monitorSystemRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(monitorSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(monitorSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
