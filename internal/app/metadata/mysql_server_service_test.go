package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestMySQLServerServiceAll(t *testing.T) {
	TestMySQLServerService_GetEntities(t)
	TestMySQLServerService_GetAll(t)
	TestMySQLServerService_GetByID(t)
	TestMySQLServerService_Create(t)
	TestMySQLServerService_Update(t)
	TestMySQLServerService_Delete(t)
	TestMySQLServerService_Marshal(t)
	TestMySQLServerService_MarshalWithFields(t)
}

func TestMySQLServerService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMySQLServerService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMySQLServerService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
	err := s.GetByID("1")
	asst.Nil(err, "test GetByID() failed")
	id := s.Entities[constant.ZeroInt].Identity()
	asst.Equal("1", id, "test GetByID() failed")
}

func TestMySQLServerService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
	err := s.Create(map[string]interface{}{
		hostIPStruct:    testInsertHostIP,
		mSPortNumStruct: testInitPortNum})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMySQLServerByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMySQLServerService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewMySQLServerService(mysqlServerRepo)
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
	err = deleteMySQLServerByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMySQLServerService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewMySQLServerService(mysqlServerRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMySQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMySQLServerService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*MySQLServerInfo

	asst := assert.New(t)

	s := NewMySQLServerService(mysqlServerRepo)
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
		asst.True(equalMySQLServerInfo(entity.(*MySQLServerInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestMySQLServerService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewMySQLServerService(mysqlServerRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteMySQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
