package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestEnvServiceAll(t *testing.T) {
	TestEnvService_GetEntities(t)
	TestEnvService_GetAll(t)
	TestEnvService_GetByID(t)
	TestEnvService_Create(t)
	TestEnvService_Update(t)
	TestEnvService_Delete(t)
	TestEnvService_Marshal(t)
	TestEnvService_MarshalWithFields(t)
}

func TestEnvService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	s := NewEnvService(envRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetEnvs()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestEnvService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewEnvService(envRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetEnvs()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestEnvService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewEnvService(envRepo)
	err := s.GetByID(1)
	asst.Nil(err, "test GetByID() failed")
	id := s.Envs[constant.ZeroInt].Identity()
	asst.Equal(1, id, "test GetByID() failed")
}

func TestEnvService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewEnvService(envRepo)
	err := s.Create(map[string]interface{}{envNameStruct: defaultEnvInfoEnvName})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteEnvByID(s.Envs[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestEnvService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createEnv()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewEnvService(envRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{envNameStruct: newEnvName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	envName := s.GetEnvs()[constant.ZeroInt].GetEnvName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newEnvName, envName)
	// delete
	err = deleteEnvByID(s.Envs[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestEnvService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createEnv()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewEnvService(envRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteEnvByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestEnvService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*EnvInfo

	asst := assert.New(t)

	s := NewEnvService(envRepo)
	err := s.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	data, err := s.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	err = json.Unmarshal(data, &entitiesUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	entities := s.GetEnvs()
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		entityUnmarshal := entitiesUnmarshal[i]
		asst.True(equal(entity.(*EnvInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestEnvService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createEnv()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewEnvService(envRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(envNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(envNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteEnvByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
