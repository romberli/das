package metadata

import (
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
)

const envEnvsStruct = "Envs"

var _ metadata.EnvService = (*EnvService)(nil)

type EnvService struct {
	metadata.EnvRepo
	Envs []metadata.Env `json:"Envs"`
}

// NewEnvService returns a new *EnvService
func NewEnvService(repo metadata.EnvRepo) *EnvService {
	return &EnvService{repo, []metadata.Env{}}
}

// NewEnvServiceWithDefault returns a new *EnvService with default EnvRepo
func NewEnvServiceWithDefault() *EnvService {
	return NewEnvService(NewEnvRepoWithGlobal())
}

// GetEnvs returns environments of the service
func (es *EnvService) GetEnvs() []metadata.Env {
	return es.Envs
}

// GetAll gets all environments from the middleware
func (es *EnvService) GetAll() error {
	var err error
	es.Envs, err = es.EnvRepo.GetAll()

	return err
}

// GetID gets identity of an entity with given fields
func (es *EnvService) GetID(fields map[string]interface{}) (int, error) {
	_, ok := fields[envNameStruct]
	if !ok {
		return constant.ZeroInt, message.NewMessage(message.ErrFieldNotExists, envNameStruct)
	}
	// create a new entity
	envInfo, err := NewEnvInfoWithMapAndRandom(fields)
	if err != nil {
		return constant.ZeroInt, err
	}
	// get identity from the middleware
	id, err := es.EnvRepo.GetID(envInfo.EnvName)
	if err != nil {
		return constant.ZeroInt, err
	}

	return id, nil
}

// GetByID gets an environment of the given id from the middleware
func (es *EnvService) GetByID(id int) error {
	entity, err := es.EnvRepo.GetByID(id)
	if err != nil {
		return err
	}

	es.Envs = append(es.Envs, entity)

	return err
}

// GetEnvByName returns Env of given env name
func (es *EnvService) GetEnvByName(envName string) error {
	env, err := es.EnvRepo.GetEnvByName(envName)
	if err != nil {
		return err
	}

	es.Envs = append(es.Envs, env)

	return nil
}

// Create creates an environment in the middleware
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
	env, err := es.EnvRepo.Create(envInfo)
	if err != nil {
		return err
	}

	es.Envs = append(es.Envs, env)

	return nil
}

// Update gets the environment of the given id from the middleware,
// and then updates its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (es *EnvService) Update(id int, fields map[string]interface{}) error {
	err := es.GetByID(id)
	if err != nil {
		return err
	}
	err = es.Envs[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return es.EnvRepo.Update(es.Envs[constant.ZeroInt])
}

// Delete deletes the environment of given id in the middleware
func (es *EnvService) Delete(id int) error {
	return es.EnvRepo.Delete(id)
}

// Marshal marshals EnvService.Envs to json bytes
func (es *EnvService) Marshal() ([]byte, error) {
	return es.MarshalWithFields(envEnvsStruct)
}

// MarshalWithFields marshals only specified fields of the EnvService to json bytes
func (es *EnvService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(es, fields...)
}
