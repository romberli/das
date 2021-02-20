package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestDbServiceAll(t *testing.T) {
	TestDbService_GetEntities(t)
	TestDbService_GetAll(t)
	TestDbService_GetByID(t)
	TestDbService_Create(t)
	TestDbService_Update(t)
	TestDbService_Delete(t)
	TestDbService_Marshal(t)
	TestDbService_MarshalWithFields(t)
}

func TestDbService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	s := NewDbService(dbRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestDbService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewDbService(dbRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestDbService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewDbService(dbRepo)
	err := s.GetByID("1")
	asst.Nil(err, "test GetByID() failed")
	id := s.Entities[constant.ZeroInt].Identity()
	asst.Equal("1", id, "test GetByID() failed")
}

func TestDbService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewDbService(dbRepo)
	err := s.Create(map[string]interface{}{dbNameStruct: defaultDbInfoDbName, ownerIdStruct:defaultDbInfoOwnerId, envIdStruct:defaultDbInfoEnvId})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteDbByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestDbService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDb()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewDbService(dbRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{dbNameStruct: newDbName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	dbName, err := s.GetEntities()[constant.ZeroInt].Get(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newDbName, dbName)
	// delete
	err = deleteDbByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestDbService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDb()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewDbService(dbRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteDbByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestDbService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*DbInfo

	asst := assert.New(t)

	s := NewDbService(dbRepo)
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
		asst.True(dbEqual(entity.(*DbInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestDbService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDb()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewDbService(dbRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteDbByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
