package metadata

import (
	"encoding/json"
	"github.com/romberli/das/internal/dependency"
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

const (
	middlewareServerClusterIDStruct      = "ClusterIDMiddleware"
	middlewareServerNameStruct           = "ServerName"
	middlewareServerMiddlewareRoleStruct = "MiddlewareRole"
	middlewareServerHostIPStruct         = "HostIP"
	middlewareServerPortNumStruct        = "PortNum"
)

var _ dependency.Service = (*MiddlewareServerService)(nil)

type MiddlewareServerService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewMiddlewareServerService returns a new *MiddlewareServerService
func NewMiddlewareServerService(repo dependency.Repository) *MiddlewareServerService {
	return &MiddlewareServerService{repo, []dependency.Entity{}}
}

// NewMiddlewareServerServiceWithDefault returns a new *MiddlewareServerService with default repository
func NewMiddlewareServerServiceWithDefault() *MiddlewareServerService {
	return NewMiddlewareServerService(NewMiddlewareServerRepoWithGlobal())
}

// GetEntities returns entities of the service
func (mss *MiddlewareServerService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(mss.Entities))
	for i := range entityList {
		entityList[i] = mss.Entities[i]
	}

	return entityList
}

// GetAll gets all middlewareServerironment entities from the middleware
func (mss *MiddlewareServerService) GetAll() error {
	var err error
	mss.Entities, err = mss.Repository.GetAll()

	return err
}

// GetByID gets an middlewareServerironment entity that contains the given id from the middleware
func (mss *MiddlewareServerService) GetByID(id string) error {
	entity, err := mss.Repository.GetByID(id)
	if err != nil {
		return err
	}

	mss.Entities = append(mss.Entities, entity)

	return err
}

// Create creates a new middlewareServerironment entity and insert it into the middleware
func (mss *MiddlewareServerService) Create(fields map[string]interface{}) error {
	// generate new map
	middlewareServerClusterID, ok := fields[middlewareServerClusterIDStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, middlewareServerClusterIDStruct)
	}
	middlewareServerName, ok := fields[middlewareServerNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, middlewareServerNameStruct)
	}
	middlewareServerMiddlewareRole, ok := fields[middlewareServerMiddlewareRoleStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, middlewareServerMiddlewareRoleStruct)
	}
	middlewareServerHostIP, ok := fields[middlewareServerHostIPStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, middlewareServerHostIPStruct)
	}
	middlewareServerPortNum, ok := fields[middlewareServerPortNumStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, middlewareServerPortNumStruct)
	}
	middlewareServerInfo := NewMiddlewareServerInfoWithDefault(middlewareServerClusterID.(int), middlewareServerName.(string), middlewareServerMiddlewareRole.(int), middlewareServerHostIP.(string), middlewareServerPortNum.(int))
	// insert into middleware
	entity, err := mss.Repository.Create(middlewareServerInfo)
	if err != nil {
		return err
	}

	mss.Entities = append(mss.Entities, entity)
	return nil
}

// Update gets an middlewareServerironment entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mss *MiddlewareServerService) Update(id string, fields map[string]interface{}) error {
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

// Delete deletes the middlewareServerironment entity that contains the given id in the middleware
func (mss *MiddlewareServerService) Delete(id string) error {
	return mss.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (mss *MiddlewareServerService) Marshal() ([]byte, error) {
	return json.Marshal(mss.Entities)
}

// Marshal marshals service.Entities with given fields
func (mss *MiddlewareServerService) MarshalWithFields(fields ...string) ([]byte, error) {
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
