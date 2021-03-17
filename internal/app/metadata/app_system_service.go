package metadata

import (
	"encoding/json"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/pkg/message"

	"github.com/romberli/das/internal/dependency"
)

const (
	appSystemNameStruct       = "AppSystemName"
	appSystemLevelStruct      = "Level"
	appSystemOwnerIDStruct    = "OwnerID"
	appSystemOwnerGroupStruct = "OwnerGroup"
)

var _ dependency.Service = (*AppSystemService)(nil)

// AppSystemService service struct
type AppSystemService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewAppSystemService returns a new *AppSystemService
func NewAppSystemService(repo dependency.Repository) *AppSystemService {
	return &AppSystemService{repo, []dependency.Entity{}}
}

// NewAppSystemServiceWithDefault returns a new *AppSystemService with default repository
func NewAppSystemServiceWithDefault() *AppSystemService {
	return NewAppSystemService(NewAppSystemRepoWithGlobal())
}

// GetEntities returns entities of the service
func (ass *AppSystemService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(ass.Entities))
	for i := range entityList {
		entityList[i] = ass.Entities[i]
	}

	return entityList
}

// GetAll ge entities from the middleware
func (ass *AppSystemService) GetAll() error {
	var err error
	ass.Entities, err = ass.Repository.GetAll()

	return err
}

// GetByID g entity that contains the given id from the middleware
func (ass *AppSystemService) GetByID(id string) error {
	entity, err := ass.Repository.GetByID(id)
	if err != nil {
		return err
	}

	ass.Entities = append(ass.Entities, entity)

	return err
}

// Create creates entity and insert it into the middleware
func (ass *AppSystemService) Create(fields map[string]interface{}) error {
	// generate new map
	appSystemName, ok := fields[appSystemNameStruct]

	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, appSystemNameStruct)
	}
	level, ok := fields[appSystemLevelStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, appSystemLevelStruct)
	}
	ownerID, ok := fields[appSystemOwnerIDStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, appSystemOwnerIDStruct)
	}
	ownerGroup, ok := fields[appSystemOwnerGroupStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, appSystemOwnerGroupStruct)
	}
	appSystemInfo := NewAppSystemInfoWithDefault(appSystemName.(string), level.(int), ownerID.(int), ownerGroup.(string))
	// insert into middleware
	entity, err := ass.Repository.Create(appSystemInfo)
	if err != nil {
		return err
	}

	ass.Entities = append(ass.Entities, entity)
	return nil
}

// Update gets an AppSystem entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (ass *AppSystemService) Update(id string, fields map[string]interface{}) error {
	err := ass.GetByID(id)
	if err != nil {
		return err
	}
	err = ass.Entities[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return ass.Repository.Update(ass.Entities[constant.ZeroInt])
}

// Delete delete entity that contains the given id in the middleware
func (ass *AppSystemService) Delete(id string) error {
	return ass.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (ass *AppSystemService) Marshal() ([]byte, error) {
	return json.Marshal(ass.Entities)
}

// MarshalWithFields marshals service.Entities with given fields
func (ass *AppSystemService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(ass.Entities))
	for i := range interfaceList {
		entity, err := common.CopyStructWithFields(ass.Entities[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = entity
	}

	return json.Marshal(interfaceList)
}
