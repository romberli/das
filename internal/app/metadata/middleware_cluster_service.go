package metadata

import (
	"encoding/json"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
	"github.com/romberli/das/pkg/message"
)

const (
	middlewareClusterNameStruct  = "ClusterName"
	middlewareClusterEnvIDStruct = "EnvID"
)

var _ dependency.Service = (*MiddlewareClusterService)(nil)

type MiddlewareClusterService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewMiddlewareClusterService returns a new *MiddlewareClusterService
func NewMiddlewareClusterService(repo dependency.Repository) *MiddlewareClusterService {
	return &MiddlewareClusterService{repo, []dependency.Entity{}}
}

// NewMiddlewareClusterServiceWithDefault returns a new *MiddlewareClusterService with default repository
func NewMiddlewareClusterServiceWithDefault() *MiddlewareClusterService {
	return NewMiddlewareClusterService(NewMiddlewareClusterRepoWithGlobal())
}

// GetEntities returns entities of the service
func (mcs *MiddlewareClusterService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(mcs.Entities))
	for i := range entityList {
		entityList[i] = mcs.Entities[i]
	}

	return entityList
}

// GetAll gets all middleware cluster entities from the middleware
func (mcs *MiddlewareClusterService) GetAll() error {
	var err error
	mcs.Entities, err = mcs.Repository.GetAll()

	return err
}

// GetByID gets an middleware cluster entity that contains the given id from the middleware
func (mcs *MiddlewareClusterService) GetByID(id string) error {
	entity, err := mcs.Repository.GetByID(id)
	if err != nil {
		return err
	}

	mcs.Entities = append(mcs.Entities, entity)

	return err
}

// Create creates a new middleware cluster entity and insert it into the middleware
func (mcs *MiddlewareClusterService) Create(fields map[string]interface{}) error {
	// generate new map
	_, ok := fields[middlewareClusterNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, middlewareClusterNameStruct)
	}
	_, ok = fields[middlewareClusterEnvIDStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, middlewareClusterEnvIDStruct)
	}
	// create a new entity
	middlewareClusterInfo, err := NewMiddlewareClusterInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	entity, err := mcs.Repository.Create(middlewareClusterInfo)
	if err != nil {
		return err
	}

	mcs.Entities = append(mcs.Entities, entity)
	return nil
}

// Update gets an middleware cluster entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mcs *MiddlewareClusterService) Update(id string, fields map[string]interface{}) error {
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

// Delete deletes the middleware cluster entity that contains the given id in the middleware
func (mcs *MiddlewareClusterService) Delete(id string) error {
	return mcs.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (mcs *MiddlewareClusterService) Marshal() ([]byte, error) {
	return json.Marshal(mcs.Entities)
}

// Marshal marshals service.Entities with given fields
func (mcs *MiddlewareClusterService) MarshalWithFields(fields ...string) ([]byte, error) {
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
