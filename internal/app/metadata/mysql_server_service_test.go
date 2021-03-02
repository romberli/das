package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestMYSQLServerServiceAll(t *testing.T) {
	TestMYSQLServerService_GetEntities(t)
	TestMYSQLServerService_GetAll(t)
	TestMYSQLServerService_GetByID(t)
	TestMYSQLServerService_Create(t)
	TestMYSQLServerService_Update(t)
	TestMYSQLServerService_Delete(t)
	TestMYSQLServerService_Marshal(t)
	TestMYSQLServerService_MarshalWithFields(t)
}

func TestMYSQLServerService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	s := NewMYSQLServerService(mysqlServerRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMYSQLServerService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewMYSQLServerService(mysqlServerRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMYSQLServerService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewMYSQLServerService(mysqlServerRepo)
	err := s.GetByID("1")
	asst.Nil(err, "test GetByID() failed")
	id := s.Entities[constant.ZeroInt].Identity()
	asst.Equal("1", id, "test GetByID() failed")
}

func TestMYSQLServerService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMYSQLServerService(mysqlServerRepo)
	err := s.Create(map[string]interface{}{
		hostIPStruct:    testInsertHostIP,
		mSPortNumStruct: testInitPortNum})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMYSQLServerByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMYSQLServerService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMYSQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewMYSQLServerService(mysqlServerRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{
		hostIPStruct:    testUpdateHostIP,
		mSPortNumStruct: testUpdatePortNum})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	hostIP, err := s.GetEntities()[constant.ZeroInt].Get(hostIPStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testUpdateHostIP, hostIP)
	portNum, err := s.GetEntities()[constant.ZeroInt].Get(mSPortNumStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testUpdatePortNum, portNum)
	// delete
	err = deleteMYSQLServerByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMYSQLServerService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMYSQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewMYSQLServerService(mysqlServerRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMYSQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMYSQLServerService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*MYSQLServerInfo

	asst := assert.New(t)

	s := NewMYSQLServerService(mysqlServerRepo)
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
		asst.True(equalMYSQLServerInfo(entity.(*MYSQLServerInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestMYSQLServerService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMYSQLServer()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewMYSQLServerService(mysqlServerRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteMYSQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
