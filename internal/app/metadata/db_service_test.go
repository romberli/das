package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestDBServiceAll(t *testing.T) {
	TestDBService_GetEntities(t)
	TestDBService_GetAll(t)
	TestDBService_GetByID(t)
	TestDBService_Create(t)
	TestDBService_Update(t)
	TestDBService_Delete(t)
	TestDBService_Marshal(t)
	TestDBService_MarshalWithFields(t)
}

func TestDBService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestDBService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestDBService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.GetByID("1")
	asst.Nil(err, "test GetByID() failed")
	id := s.Entities[constant.ZeroInt].Identity()
	asst.Equal("1", id, "test GetByID() failed")
}

func TestDBService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.Create(map[string]interface{}{dbNameStruct: defaultDBInfoDBName, dbClusterIDStruct: defaultDBInfoClusterID, dbClusterTypeStruct: defaultDBInfoClusterType, dbEnvIDStruct: defaultDBInfoEnvID})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteDBByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestDBService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewDBService(dbRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{dbNameStruct: newDBName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	dbName, err := s.GetEntities()[constant.ZeroInt].Get(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newDBName, dbName)
	// delete
	err = deleteDBByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestDBService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewDBService(dbRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestDBService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*DBInfo

	asst := assert.New(t)

	s := NewDBService(dbRepo)
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
		asst.True(dbEqual(entity.(*DBInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestDBService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewDBService(dbRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}