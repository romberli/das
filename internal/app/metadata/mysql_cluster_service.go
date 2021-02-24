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

var _ dependency.Service = (*MYSQLClusterService)(nil)

// MYSQLClusterService implements Service interface
type MYSQLClusterService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewMYSQLClusterService returns a new *MYSQLClusterService
func NewMYSQLClusterService(repo dependency.Repository) *MYSQLClusterService {
	return &MYSQLClusterService{repo, []dependency.Entity{}}
}

// NewMYSQLClusterServiceWithDefault returns a new *MYSQLClusterService with default repository
func NewMYSQLClusterServiceWithDefault() *MYSQLClusterService {
	return NewMYSQLClusterService(NewMYSQLClusterRepoWithGlobal())
}

// GetEntities returns entities of the service
func (es *MYSQLClusterService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(es.Entities))
	for i := range entityList {
		entityList[i] = es.Entities[i]
	}

	return entityList
}

// GetAll gets all environment entities from the middleware
func (es *MYSQLClusterService) GetAll() error {
	var err error
	es.Entities, err = es.Repository.GetAll()

	return err
}

// GetByID gets an environment entity that contains the given id from the middleware
func (es *MYSQLClusterService) GetByID(id string) error {
	entity, err := es.Repository.GetByID(id)
	if err != nil {
		return err
	}

	es.Entities = append(es.Entities, entity)

	return err
}

// Create creates a new environment entity and insert it into the middleware
func (es *MYSQLClusterService) Create(fields map[string]interface{}) error {
	// generate new map
	clusterName, ok := fields[clusterNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, clusterNameStruct)
	}
	mysqlClusterInfo := NewMYSQLClusterInfoWithDefault(clusterName.(string))
	// insert into middleware
	entity, err := es.Repository.Create(mysqlClusterInfo)
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
func (es *MYSQLClusterService) Update(id string, fields map[string]interface{}) error {
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
func (es *MYSQLClusterService) Delete(id string) error {
	return es.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (es *MYSQLClusterService) Marshal() ([]byte, error) {
	return json.Marshal(es.Entities)
}

// MarshalWithFields marshals service.Entities with given fields
func (es *MYSQLClusterService) MarshalWithFields(fields ...string) ([]byte, error) {
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
