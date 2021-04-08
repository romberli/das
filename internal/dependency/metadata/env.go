package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type Env interface {
	// Identity returns the identity
	Identity() int
	// GetEnvName returns the env name
	GetEnvName() string
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// Set sets Env with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// MarshalJSON marshals Env to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified fields of Env to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type EnvRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns middleware.PoolConn, so it can run multiple statements in the same transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all environments from the middleware
	GetAll() ([]Env, error)
	// GetByID gets an environment by the identity from the middleware
	GetByID(id int) (Env, error)
	// GetID gets the identity with given environment name from the middleware
	GetID(envName string) (int, error)
	// GetEnvByName gets Env of given environment name
	GetEnvByName(envName string) (Env, error)
	// Create creates an environment in the middleware
	Create(env Env) (Env, error)
	// Update updates the environment in the middleware
	Update(env Env) error
	// Delete deletes the environment in the middleware
	Delete(id int) error
}

type EnvService interface {
	// GetEnvs returns environments of the service
	GetEnvs() []Env
	// GetAll gets all environments from the middleware
	GetAll() error
	// GetByID gets an environment of the given id from the middleware
	GetByID(id int) error
	// GetEnvByName returns Env of given env name
	GetEnvByName(envName string) error
	// Create creates an environment in the middleware
	Create(fields map[string]interface{}) error
	// Update gets the environment of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the environment of given id in the middleware
	Delete(id int) error
	// Marshal marshals EnvService.Envs to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the EnvService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
