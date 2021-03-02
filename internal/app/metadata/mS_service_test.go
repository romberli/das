package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func TestMSServiceAll(t *testing.T) {
	TestMSService_GetEntities(t)
	TestMSService_GetAll(t)
	TestMSService_GetByID(t)
	TestMSService_Create(t)
	TestMSService_Update(t)
	TestMSService_Delete(t)
	TestMSService_Marshal(t)
	TestMSService_MarshalWithFields(t)
}

func TestMSService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	s := NewMSService(mSRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMSService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewMSService(mSRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMSService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewMSService(mSRepo)
	err := s.GetByID("1")
	asst.Nil(err, "test GetByID() failed")
	id := s.Entities[constant.ZeroInt].Identity()
	asst.Equal("1", id, "test GetByID() failed")
}

func TestMSService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMSService(mSRepo)
	err := s.Create(map[string]interface{}{mSNameStruct: defaultMSInfoMSName, systemTypeStruct: defaultMSInfoSystemType, mSHostIpStruct: defaultMSInfoHostIp, mSPortNumStruct: defaultMSInfoPortNum, portNumSlowStruct: defaultMSInfoPortNumSlow, baseUrlStruct: defaultMSInfoBaseUrl})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMSByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMSService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMS()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewMSService(mSRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{mSNameStruct: newMSName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	mSName, err := s.GetEntities()[constant.ZeroInt].Get(mSNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newMSName, mSName)
	// delete
	err = deleteMSByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMSService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMS()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewMSService(mSRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMSByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMSService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*MSInfo

	asst := assert.New(t)

	s := NewMSService(mSRepo)
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
		asst.True(mSEqual(entity.(*MSInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestMSService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMS()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewMSService(mSRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(mSNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(mSNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteMSByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
