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
	monitorSystemNameStruct    = "MonitorSystemName"
	monitorSystemTypeStruct    = "MonitorSystemType"
	monitorSystemHostIPStruct  = "MonitorSystemHostIP"
	monitorSystemPortNumStruct = "MonitorSystemPortNum"
	portNumSlowStruct          = "MonitorSystemPortNumSlow"
	baseUrlStruct              = "BaseUrl"
)

var _ dependency.Service = (*MonitorSystemService)(nil)

type MonitorSystemService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewMonitorSystemService returns a new *MonitorSystemService
func NewMonitorSystemService(repo dependency.Repository) *MonitorSystemService {
	return &MonitorSystemService{repo, []dependency.Entity{}}
}

// NewMonitorSystemServiceWithDefault returns a new *MonitorSystemService with default repository
func NewMonitorSystemServiceWithDefault() *MonitorSystemService {
	return NewMonitorSystemService(NewMonitorSystemRepoWithGlobal())
}

// GetEntities returns entities of the service
func (mss *MonitorSystemService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(mss.Entities))
	for i := range entityList {
		entityList[i] = mss.Entities[i]
	}

	return entityList
}

// GetAll gets all monitor system entities from the middleware
func (mss *MonitorSystemService) GetAll() error {
	var err error
	mss.Entities, err = mss.Repository.GetAll()

	return err
}

// GetByID gets an monitor system entity that contains the given id from the middleware
func (mss *MonitorSystemService) GetByID(id string) error {
	entity, err := mss.Repository.GetByID(id)
	if err != nil {
		return err
	}

	mss.Entities = append(mss.Entities, entity)

	return err
}

// Create creates a new monitor system entity and insert it into the middleware
func (mss *MonitorSystemService) Create(fields map[string]interface{}) error {
	// generate new map
	_, monitorSystemNameExists := fields[monitorSystemNameStruct]
	_, systemTypeExists := fields[monitorSystemTypeStruct]
	_, hostIPExists := fields[monitorSystemHostIPStruct]
	_, portNumExists := fields[monitorSystemPortNumStruct]
	_, portNumSlowExists := fields[portNumSlowStruct]
	_, baseUrlExists := fields[baseUrlStruct]
	if !monitorSystemNameExists && !systemTypeExists && !hostIPExists && !portNumExists && !portNumSlowExists && !baseUrlExists {
		return message.NewMessage(message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s and %s and %s", monitorSystemNameStruct, monitorSystemTypeStruct, monitorSystemHostIPStruct, monitorSystemPortNumStruct, portNumSlowStruct, baseUrlStruct))
	}
	// create a new entity
	monitorSystemInfo, err := NewMonitorSystemInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	entity, err := mss.Repository.Create(monitorSystemInfo)
	if err != nil {
		return err
	}

	mss.Entities = append(mss.Entities, entity)
	return nil
}

// Update gets an monitor system entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mss *MonitorSystemService) Update(id string, fields map[string]interface{}) error {
	err := mss.GetByID(id)
	if err != nil {
		return err
	}
	err = mss.Entities[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return mss.Repository.Update(mss.Entities[constant.ZeroInt])
}

// Delete deletes the monitor system entity that contains the given id in the middleware
func (mss *MonitorSystemService) Delete(id string) error {
	return mss.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (mss *MonitorSystemService) Marshal() ([]byte, error) {
	return json.Marshal(mss.Entities)
}

// Marshal marshals service.Entities with given fields
func (mss *MonitorSystemService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(mss.Entities))
	for i := range interfaceList {
		entity, err := common.CopyStructWithFields(mss.Entities[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = entity
	}

	return json.Marshal(interfaceList)
}
