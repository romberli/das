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
	dbNameStruct        = "DBName"
	dbClusterIDStruct   = "ClusterID"
	dbClusterTypeStruct = "ClusterType"
	dbOwnerIDStruct     = "OwnerID"
	dbOwnerGroupStruct  = "OwnerGroup"
	dbEnvIDStruct       = "EnvID"
)

var _ dependency.Service = (*DBService)(nil)

type DBService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewDBService returns a new *DBService
func NewDBService(repo dependency.Repository) *DBService {
	return &DBService{repo, []dependency.Entity{}}
}

// NewDBServiceWithDefault returns a new *DBService with default repository
func NewDBServiceWithDefault() *DBService {
	return NewDBService(NewDBRepoWithGlobal())
}

// GetEntities returns entities of the service
func (dbs *DBService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(dbs.Entities))
	for i := range entityList {
		entityList[i] = dbs.Entities[i]
	}

	return entityList
}

// GetAll gets all database entities from the middleware
func (dbs *DBService) GetAll() error {
	var err error
	dbs.Entities, err = dbs.Repository.GetAll()

	return err
}

// GetByID gets an database entity that contains the given id from the middleware
func (dbs *DBService) GetByID(id string) error {
	entity, err := dbs.Repository.GetByID(id)
	if err != nil {
		return err
	}

	dbs.Entities = append(dbs.Entities, entity)

	return err
}

// Create creates a new database entity and insert it into the middleware
func (dbs *DBService) Create(fields map[string]interface{}) error {
	// generate new map
	_, dbNameExists := fields[dbNameStruct]
	_, clusterIDExists := fields[dbClusterIDStruct]
	_, clusterTypeExists := fields[dbClusterTypeStruct]
	_, envIDExists := fields[dbEnvIDStruct]
	if !dbNameExists && !clusterIDExists && !clusterTypeExists && !envIDExists {
		return message.NewMessage(message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s", dbNameStruct, dbClusterIDStruct, dbClusterTypeStruct, dbEnvIDStruct))
	}
	// create a new entity
	dbInfo, err := NewDBInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	entity, err := dbs.Repository.Create(dbInfo)
	if err != nil {
		return err
	}

	dbs.Entities = append(dbs.Entities, entity)
	return nil
}

// Update gets an database entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (dbs *DBService) Update(id string, fields map[string]interface{}) error {
	err := dbs.GetByID(id)
	if err != nil {
		return err
	}
	err = dbs.Entities[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return dbs.Repository.Update(dbs.Entities[constant.ZeroInt])
}

// Delete deletes the database entity that contains the given id in the middleware
func (dbs *DBService) Delete(id string) error {
	return dbs.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (dbs *DBService) Marshal() ([]byte, error) {
	return json.Marshal(dbs.Entities)
}

// Marshal marshals service.Entities with given fields
func (dbs *DBService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(dbs.Entities))
	for i := range interfaceList {
		entity, err := common.CopyStructWithFields(dbs.Entities[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = entity
	}

	return json.Marshal(interfaceList)
}
