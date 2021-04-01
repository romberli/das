package metadata

import (
	"encoding/json"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
	"github.com/romberli/das/pkg/message"
)

const envNameStruct = "EnvName"

var _ dependency.Service = (*EnvService)(nil)

// EnvService implements Service interface
type EnvService struct {
	dependency.Repository
	Envs []dependency.Entity
}

// NewEnvService returns a new *EnvService
func NewEnvService(repo dependency.Repository) *EnvService {
	return &EnvService{repo, []dependency.Entity{}}
}

// NewEnvServiceWithDefault returns a new *EnvService with default repository
func NewEnvServiceWithDefault() *EnvService {
	return NewEnvService(NewEnvRepoWithGlobal())
}

// GetEntities returns entities of the service
func (es *EnvService) GetEntities() []dependency.Entity {
	envList := make([]dependency.Entity, len(es.Envs))
	for i := range envList {
		envList[i] = es.Envs[i]
	}

	return envList
}

// GetAll gets all environment entities from the middleware
func (es *EnvService) GetAll() error {
	var err error
	es.Envs, err = es.Repository.GetAll()

	return err
}

// GetID gets identity of an entity with given fields
func (es *EnvService) GetID(fields map[string]interface{}) (string, error) {
	_, ok := fields[envNameStruct]
	if !ok {
		return constant.EmptyString, message.NewMessage(message.ErrFieldNotExists, envNameStruct)
	}
	// create a new entity
	envInfo, err := NewEnvInfoWithMapAndRandom(fields)
	if err != nil {
		return constant.EmptyString, err
	}
	// get identity from the middleware
	id, err := es.Repository.GetID(envInfo)
	if err != nil {
		return constant.EmptyString, err
	}

	return id, nil
}

// GetByID gets an environment entity that contains the given id from the middleware
func (es *EnvService) GetByID(id string) error {
	entity, err := es.Repository.GetByID(id)
	if err != nil {
		return err
	}

	es.Envs = append(es.Envs, entity)

	return err
}

// Create creates a new environment entity and insert it into the middleware
func (es *EnvService) Create(fields map[string]interface{}) error {
	// generate new map
	_, ok := fields[envNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, envNameStruct)
	}
	// create a new entity
	envInfo, err := NewEnvInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	env, err := es.Repository.Create(envInfo)
	if err != nil {
		return err
	}

	es.Envs = append(es.Envs, env)

	return nil
}

// Update gets an environment entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (es *EnvService) Update(id string, fields map[string]interface{}) error {
	err := es.GetByID(id)
	if err != nil {
		return err
	}
	err = es.Envs[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return es.Repository.Update(es.Envs[constant.ZeroInt])
}

// Delete deletes the environment entity that contains the given id in the middleware
func (es *EnvService) Delete(id string) error {
	return es.Repository.Delete(id)
}

// Marshal marshals service.Envs
func (es *EnvService) Marshal() ([]byte, error) {
	return json.Marshal(es.Envs)
}

// MarshalWithFields marshals service.Envs with given fields
func (es *EnvService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(es.Envs))
	for i := range interfaceList {
		env, err := common.CopyStructWithFields(es.Envs[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = env
	}

	return json.Marshal(interfaceList)
}
