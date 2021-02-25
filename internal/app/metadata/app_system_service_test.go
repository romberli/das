package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestAppSystemServiceAll(t *testing.T) {
	TestAppSystemService_GetEntities(t)
	TestAppSystemService_GetAll(t)
	TestAppSystemService_GetByID(t)
	TestAppSystemService_Create(t)
	TestAppSystemService_Update(t)
	TestAppSystemService_Delete(t)
	TestAppSystemService_Marshal(t)
	TestAppSystemService_MarshalWithFields(t)
}

func TestAppSystemService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	s := NewAppSystemService(appSystemRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestAppSystemService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewAppSystemService(appSystemRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestAppSystemService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewAppSystemService(appSystemRepo)
	err := s.GetByID("66")
	asst.Nil(err, "test GetByID() failed")
	id := s.Entities[constant.ZeroInt].Identity()
	asst.Equal("66", id, "test GetByID() failed")
}

func TestAppSystemService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewAppSystemService(appSystemRepo)
	err := s.Create(map[string]interface{}{appSystemNameStruct: defaultAppSystemInfoAppSystemName, appSystemLevelStruct: defaultAppSystemInfoLevel, appSystemOwnerIDStruct: defaultAppSystemInfoOwnerID, appSystemOwnerGroupStruct: defaultAppSystemInfoOwnerGroup})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteAppSystemByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestAppSystemService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createAppSystem()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewAppSystemService(appSystemRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{appSystemNameStruct: newAppSystemName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	appSystemName, err := s.GetEntities()[constant.ZeroInt].Get(appSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newAppSystemName, appSystemName)
	// delete
	err = deleteAppSystemByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestAppSystemService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createAppSystem()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewAppSystemService(appSystemRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteAppSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestAppSystemService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*AppSystemInfo

	asst := assert.New(t)

	s := NewAppSystemService(appSystemRepo)
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
		asst.True(appSystemEqual(entity.(*AppSystemInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestAppSystemService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createAppSystem()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewAppSystemService(appSystemRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(appSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(appSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteAppSystemByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
