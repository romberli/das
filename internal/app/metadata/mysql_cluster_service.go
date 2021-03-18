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

// GetAll gets all mysql cluster entities from the middleware
func (mcs *MySQLClusterService) GetAll() error {
	var err error
	mcs.Entities, err = mcs.Repository.GetAll()

	return err
}

// GetByID gets an mysql cluster entity that contains the given id from the middleware
func (mcs *MySQLClusterService) GetByID(id string) error {
	entity, err := mcs.Repository.GetByID(id)
	if err != nil {
		return err
	}

	mcs.Entities = append(mcs.Entities, entity)

	return err
}

// Create creates a new mysql cluster entity and insert it into the middleware
func (mcs *MySQLClusterService) Create(fields map[string]interface{}) error {
	// generate new map
	if _, ok := fields[clusterNameStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, clusterNameStruct)
	}
	if _, ok := fields[middlewareClusterIDStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, middlewareClusterIDStruct)
	}
	if _, ok := fields[monitorSystemIDStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, monitorSystemIDStruct)
	}
	if _, ok := fields[ownerIDStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, ownerIDStruct)
	}
	if _, ok := fields[ownerGroupStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, ownerGroupStruct)
	}
	if _, ok := fields[envIDStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, envIDStruct)
	}
	// create a new entity
	mysqlClusterInfo, err := NewMySQLClusterInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	entity, err := mcs.Repository.Create(mysqlClusterInfo)
	if err != nil {
		return err
	}

	mcs.Entities = append(mcs.Entities, entity)
	return nil
}

// Update gets an mysql cluster entity that contains the given id from the middleware,
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

// Delete deletes the mysql cluster entity that contains the given id in the middleware
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
