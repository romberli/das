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
	dbNameStruct = "DbName"
	ownerIdStruct = "OwnerId"
	envIdStruct = "EnvId"
)

var _ dependency.Service = (*DbService)(nil)

type DbService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewDbService returns a new *DbService
func NewDbService(repo dependency.Repository) *DbService {
	return &DbService{repo, []dependency.Entity{}}
}

// NewDbServiceWithDefault returns a new *DbService with default repository
func NewDbServiceWithDefault() *DbService {
	return NewDbService(NewDbRepoWithGlobal())
}

// GetEntities returns entities of the service
func (es *DbService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(es.Entities))
	for i := range entityList {
		entityList[i] = es.Entities[i]
	}

	return entityList
}

// GetAll gets all database entities from the middleware
func (es *DbService) GetAll() error {
	var err error
	es.Entities, err = es.Repository.GetAll()

	return err
}

// GetByID gets an database entity that contains the given id from the middleware
func (es *DbService) GetByID(id string) error {
	entity, err := es.Repository.GetByID(id)
	if err != nil {
		return err
	}

	es.Entities = append(es.Entities, entity)

	return err
}

// Create creates a new database entity and insert it into the middleware
func (es *DbService) Create(fields map[string]interface{}) error {
	// generate new map
	dbName, dbNameExists := fields[dbNameStruct]
	ownerId, ownerIdExists := fields[ownerIdStruct]
	envId, envIdExists := fields[envIdStruct]
	if !dbNameExists && !ownerIdExists && !envIdExists {
		return message.NewMessage(message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s", dbNameStruct, ownerIdStruct, envIdStruct))
	}
	dbInfo := NewDbInfoWithDefault(dbName.(string), ownerId.(string), envId.(string))
	// insert into middleware
	entity, err := es.Repository.Create(dbInfo)
	if err != nil {
		return err
	}

	es.Entities = append(es.Entities, entity)
	return nil
}

// Update gets an database entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (es *DbService) Update(id string, fields map[string]interface{}) error {
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

// Delete deletes the database entity that contains the given id in the middleware
func (es *DbService) Delete(id string) error {
	return es.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (es *DbService) Marshal() ([]byte, error) {
	return json.Marshal(es.Entities)
}

// Marshal marshals service.Entities with given fields
func (es *DbService) MarshalWithFields(fields ...string) ([]byte, error) {
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
