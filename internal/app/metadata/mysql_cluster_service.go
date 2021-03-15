package metadata

import (
	"encoding/json"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/pkg/message"

	"github.com/romberli/das/internal/dependency"
)

const (
	clusterNameStruct         = "ClusterName"
	middlewareClusterIDStruct = "MiddlewareClusterID"
	monitorSystemIDStruct     = "MonitorSystemID"
	ownerIDStruct             = "OwnerID"
	ownerGroupStruct          = "OwnerGroup"
	envIDStruct               = "EnvID"
)

var _ dependency.Service = (*MySQLClusterService)(nil)

// MySQLClusterService implements Service interface
type MySQLClusterService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewMySQLClusterService returns a new *MySQLClusterService
func NewMySQLClusterService(repo dependency.Repository) *MySQLClusterService {
	return &MySQLClusterService{repo, []dependency.Entity{}}
}

// NewMySQLClusterServiceWithDefault returns a new *MySQLClusterService with default repository
func NewMySQLClusterServiceWithDefault() *MySQLClusterService {
	return NewMySQLClusterService(NewMySQLClusterRepoWithGlobal())
}

// GetEntities returns entities of the service
func (mcs *MySQLClusterService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(mcs.Entities))
	for i := range entityList {
		entityList[i] = mcs.Entities[i]
	}

	return entityList
}

// GetAll gets all environment entities from the middleware
func (mcs *MySQLClusterService) GetAll() error {
	var err error
	mcs.Entities, err = mcs.Repository.GetAll()

	return err
}

// GetByID gets an environment entity that contains the given id from the middleware
func (mcs *MySQLClusterService) GetByID(id string) error {
	entity, err := mcs.Repository.GetByID(id)
	if err != nil {
		return err
	}

	mcs.Entities = append(mcs.Entities, entity)

	return err
}

// Create creates a new environment entity and insert it into the middleware
func (mcs *MySQLClusterService) Create(fields map[string]interface{}) error {
	// generate new map
	clusterName, ok := fields[clusterNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, clusterNameStruct)
	}
	mysqlClusterInfo := NewMySQLClusterInfoWithDefault(clusterName.(string))
	// insert into middleware
	entity, err := mcs.Repository.Create(mysqlClusterInfo)
	if err != nil {
		return err
	}

	mcs.Entities = append(mcs.Entities, entity)
	return nil
}

// Update gets an environment entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mcs *MySQLClusterService) Update(id string, fields map[string]interface{}) error {
	err := mcs.GetByID(id)
	if err != nil {
		return err
	}
	err = mcs.Entities[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return mcs.Repository.Update(mcs.Entities[constant.ZeroInt])
}

// Delete deletes the environment entity that contains the given id in the middleware
func (mcs *MySQLClusterService) Delete(id string) error {
	return mcs.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (mcs *MySQLClusterService) Marshal() ([]byte, error) {
	return json.Marshal(mcs.Entities)
}

// MarshalWithFields marshals service.Entities with given fields
func (mcs *MySQLClusterService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(mcs.Entities))
	for i := range interfaceList {
		entity, err := common.CopyStructWithFields(mcs.Entities[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = entity
	}

	return json.Marshal(interfaceList)
}
