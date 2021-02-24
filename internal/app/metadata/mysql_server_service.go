package metadata

import (
	"encoding/json"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/pkg/message"

	"github.com/romberli/das/internal/dependency"
)

const (
	clusterIDStruct      = "ClusterID"
	hostIPStruct         = "HostIP"
	portNumStruct        = "PortNum"
	deploymentTypeStruct = "DeploymentType"
	versionStruct        = "Version"
)

var _ dependency.Service = (*MYSQLServerService)(nil)

// MYSQLServerService implements Service interface
type MYSQLServerService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewMYSQLServerService returns a new *MYSQLServerService
func NewMYSQLServerService(repo dependency.Repository) *MYSQLServerService {
	return &MYSQLServerService{repo, []dependency.Entity{}}
}

// NewMYSQLServerServiceWithDefault returns a new *MYSQLServerService with default repository
func NewMYSQLServerServiceWithDefault() *MYSQLServerService {
	return NewMYSQLServerService(NewMYSQLServerRepoWithGlobal())
}

// GetEntities returns entities of the service
func (es *MYSQLServerService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(es.Entities))
	for i := range entityList {
		entityList[i] = es.Entities[i]
	}

	return entityList
}

// GetAll gets all environment entities from the middleware
func (es *MYSQLServerService) GetAll() error {
	var err error
	es.Entities, err = es.Repository.GetAll()

	return err
}

// GetByID gets an environment entity that contains the given id from the middleware
func (es *MYSQLServerService) GetByID(id string) error {
	entity, err := es.Repository.GetByID(id)
	if err != nil {
		return err
	}

	es.Entities = append(es.Entities, entity)

	return err
}

// Create creates a new environment entity and insert it into the middleware
func (es *MYSQLServerService) Create(fields map[string]interface{}) error {
	// generate new map
	hostIP, ok := fields[hostIPStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, hostIPStruct)
	}
	portNum, ok := fields[portNumStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, portNumStruct)
	}
	mysqlServerInfo := NewMYSQLServerInfoWithDefault(hostIP.(string), portNum.(int))
	// insert into middleware
	entity, err := es.Repository.Create(mysqlServerInfo)
	if err != nil {
		return err
	}

	es.Entities = append(es.Entities, entity)
	return nil
}

// Update gets an environment entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (es *MYSQLServerService) Update(id string, fields map[string]interface{}) error {
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

// Delete deletes the environment entity that contains the given id in the middleware
func (es *MYSQLServerService) Delete(id string) error {
	return es.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (es *MYSQLServerService) Marshal() ([]byte, error) {
	return json.Marshal(es.Entities)
}

// MarshalWithFields marshals service.Entities with given fields
func (es *MYSQLServerService) MarshalWithFields(fields ...string) ([]byte, error) {
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
