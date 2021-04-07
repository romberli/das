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
	TestMonitorSystemService_GetMonitorSystems(t)
	TestMonitorSystemService_GetAll(t)
	TestMonitorSystemService_GetByID(t)
	TestMonitorSystemService_GetByEnv(t)
	TestMonitorSystemService_GetByHostInfo(t)
	TestMonitorSystemService_Create(t)
	TestMonitorSystemService_Update(t)
	TestMonitorSystemService_Delete(t)
	TestMonitorSystemService_Marshal(t)
	TestMonitorSystemService_MarshalWithFields(t)
}

func TestMonitorSystemService_GetMonitorSystems(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test GetMonitorSystems() failed", err))
	s := NewMonitorSystemService(monitorSystemRepo)
	err = s.GetAll()
	asst.Nil(err, "test GetMonitorSystems() failed")
	entities := s.GetMonitorSystems()
	asst.Greater(len(entities), constant.ZeroInt, "test GetMonitorSystems() failed")
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetMonitorSystems() failed", err))
}

func TestMonitorSystemService_GetAll(t *testing.T) {
	asst := assert.New(t)
	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	s := NewMonitorSystemService(monitorSystemRepo)
	err = s.GetAll()
	asst.Nil(err, "test GetAll() failed")
	entities := s.GetMonitorSystems()
	asst.Greater(len(entities), constant.ZeroInt, "test GetAll() failed")
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
}

func TestMonitorSystemService_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	s := NewMonitorSystemService(monitorSystemRepo)
	err = s.GetByID(entity.Identity())
	asst.Nil(err, "test GetByID() failed")
	id := s.MonitorSystems[constant.ZeroInt].Identity()
	asst.Equal(entity.Identity(), id, "test GetByID() failed")
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
}

func TestMonitorSystemService_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
	s := NewMonitorSystemService(monitorSystemRepo)
	err = s.GetByEnv(1)
	asst.Nil(err, "test GetByEnv() failed")
	envId := s.MonitorSystems[constant.ZeroInt].GetEnvID()
	asst.Equal(1, envId, "test GetByEnv() failed")
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
}

func TestMonitorSystemService_GetByHostInfo(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test GetByHostInfo() failed", err))
	s := NewMonitorSystemService(monitorSystemRepo)
	err = s.GetByHostInfo("0.0.0.0", 3306)
	asst.Nil(err, "test GetByHostInfo() failed")
	systemName := s.MonitorSystems[constant.ZeroInt].GetSystemName()
	asst.Equal(defaultMonitorSystemInfoSystemName, systemName, "test GetByHostInfo() failed")
	// delete
	err = deleteMonitorSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByHostInfo() failed", err))
}

func TestMonitorSystemService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMonitorSystemService(monitorSystemRepo)
	err := s.Create(map[string]interface{}{monitorSystemNameStruct: defaultMonitorSystemInfoSystemName,
		monitorSystemTypeStruct: defaultMonitorSystemInfoSystemType, monitorSystemHostIPStruct: defaultMonitorSystemInfoHostIP,
		monitorSystemPortNumStruct: defaultMonitorSystemInfoPortNum, monitorSystemPortNumSlowStruct: defaultMonitorSystemInfoPortNumSlow,
		monitorSystemBaseUrlStruct: defaultMonitorSystemInfoBaseUrl, monitorSystemEnvIDStruct: defaultMonitorSystemInfoEnvID})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMonitorSystemByID(s.MonitorSystems[0].Identity())
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
	monitorSystemName := s.GetMonitorSystems()[constant.ZeroInt].GetSystemName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newMonitorSystemName, monitorSystemName)
	// delete
	err = deleteMonitorSystemByID(s.MonitorSystems[0].Identity())
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
	entities := s.GetMonitorSystems()
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
