package metadata

import (
	"encoding/json"
	"fmt"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/pkg/message"

	"github.com/romberli/das/internal/dependency"
)

const (
	mSNameStruct      = "MSName"
	systemTypeStruct  = "SystemType"
	mSHostIpStruct    = "HostIp"
	mSPortNumStruct   = "PortNum"
	portNumSlowStruct = "PortNumSlow"
	baseUrlStruct     = "BaseUrl"
)

var _ dependency.Service = (*MSService)(nil)

type MSService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewMSService returns a new *MSService
func NewMSService(repo dependency.Repository) *MSService {
	return &MSService{repo, []dependency.Entity{}}
}

// NewMonitorSystemServiceWithDefault returns a new *MSService with default repository
func NewMonitorSystemServiceWithDefault() *MSService {
	return NewMSService(NewMSRepoWithGlobal())
}

// GetEntities returns entities of the service
func (es *MSService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(es.Entities))
	for i := range entityList {
		entityList[i] = es.Entities[i]
	}

	return entityList
}

// GetAll gets all monitor system entities from the middleware
func (es *MSService) GetAll() error {
	var err error
	es.Entities, err = es.Repository.GetAll()

	return err
}

// GetByID gets an monitor system entity that contains the given id from the middleware
func (es *MSService) GetByID(id string) error {
	entity, err := es.Repository.GetByID(id)
	if err != nil {
		return err
	}

	es.Entities = append(es.Entities, entity)

	return err
}

// Create creates a new monitor system entity and insert it into the middleware
func (es *MSService) Create(fields map[string]interface{}) error {
	// generate new map
	mSName, mSNameExists := fields[mSNameStruct]
	systemType, systemTypeExists := fields[systemTypeStruct]
	hostIp, hostIpExists := fields[mSHostIpStruct]
	portNum, portNumExists := fields[mSPortNumStruct]
	portNumSlow, portNumSlowExists := fields[portNumSlowStruct]
	baseUrl, baseUrlExists := fields[baseUrlStruct]

	if !mSNameExists && !systemTypeExists && !hostIpExists && !portNumExists && !portNumSlowExists && !baseUrlExists {
		return message.NewMessage(message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s and %s and %s", mSNameStruct, systemTypeStruct, mSHostIpStruct, mSPortNumStruct, portNumSlowStruct, baseUrlStruct))
	}
	mSInfo := NewMSInfoWithDefault(mSName.(string), systemType.(string), hostIp.(string), portNum.(string), portNumSlow.(string), baseUrl.(string))
	// insert into middleware
	entity, err := es.Repository.Create(mSInfo)
	if err != nil {
		return err
	}

	es.Entities = append(es.Entities, entity)
	return nil
}

// Update gets an monitor system entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (es *MSService) Update(id string, fields map[string]interface{}) error {
	err := es.GetByID(id)
	if err != nil {
		return err
	}
	err = es.Entities[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return es.Repository.Update(es.Entities[constant.ZeroInt])
}

// Delete deletes the monitor system entity that contains the given id in the middleware
func (es *MSService) Delete(id string) error {
	return es.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (es *MSService) Marshal() ([]byte, error) {
	return json.Marshal(es.Entities)
}

// Marshal marshals service.Entities with given fields
func (es *MSService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(es.Entities))
	for i := range interfaceList {
		entity, err := common.CopyStructWithFields(es.Entities[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = entity
	}

	return json.Marshal(interfaceList)
}
