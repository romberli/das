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

var _ dependency.Service = (*MySQLServerService)(nil)

// MySQLServerService implements Service interface
type MySQLServerService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewMySQLServerService returns a new *MySQLServerService
func NewMySQLServerService(repo dependency.Repository) *MySQLServerService {
	return &MySQLServerService{repo, []dependency.Entity{}}
}

// NewMySQLServerServiceWithDefault returns a new *MySQLServerService with default repository
func NewMySQLServerServiceWithDefault() *MySQLServerService {
	return NewMySQLServerService(NewMySQLServerRepoWithGlobal())
}

// GetEntities returns entities of the service
func (mss *MySQLServerService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(mss.Entities))
	for i := range entityList {
		entityList[i] = mss.Entities[i]
	}

	return entityList
}

// GetAll gets all environment entities from the middleware
func (mss *MySQLServerService) GetAll() error {
	var err error
	mss.Entities, err = mss.Repository.GetAll()

	return err
}

// GetByID gets an environment entity that contains the given id from the middleware
func (mss *MySQLServerService) GetByID(id string) error {
	entity, err := mss.Repository.GetByID(id)
	if err != nil {
		return err
	}

	mss.Entities = append(mss.Entities, entity)

	return err
}

// Create creates a new environment entity and insert it into the middleware
func (mss *MySQLServerService) Create(fields map[string]interface{}) error {
	// generate new map
	hostIP, ok := fields[hostIPStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, hostIPStruct)
	}
	portNum, ok := fields[portNumStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, portNumStruct)
	}
	mysqlServerInfo := NewMySQLServerInfoWithDefault(hostIP.(string), portNum.(int))
	// insert into middleware
	entity, err := mss.Repository.Create(mysqlServerInfo)
	if err != nil {
		return err
	}

	mss.Entities = append(mss.Entities, entity)
	return nil
}

// Update gets an environment entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mss *MySQLServerService) Update(id string, fields map[string]interface{}) error {
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

// Delete deletes the environment entity that contains the given id in the middleware
func (mss *MySQLServerService) Delete(id string) error {
	return mss.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (mss *MySQLServerService) Marshal() ([]byte, error) {
	return json.Marshal(mss.Entities)
}

// MarshalWithFields marshals service.Entities with given fields
func (mss *MySQLServerService) MarshalWithFields(fields ...string) ([]byte, error) {
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
