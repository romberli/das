package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type MonitorSystem interface {
	// Identity returns the identity
	Identity() int
	// GetSystemName returns the system name
	GetSystemName() string
	// GetSystemType returns the system type
	GetSystemType() int
	// GetHostIP returns the host ip
	GetHostIP() string
	// GetPortNum returns the port number
	GetPortNum() int
	// GetPortNumSlow returns the slow log port number
	GetPortNumSlow() int
	// GetBaseURL returns the base url
	GetBaseURL() string
	// GetEnvID returns env id
	GetEnvID() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// Set sets DB with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// MarshalJSON marshals DB to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the DB to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type MonitorSystemRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all monitor systems from the middleware
	GetAll() ([]MonitorSystem, error)
	// GetByEnv gets monitor systems of given env id from the middleware
	GetByEnv(envID int) ([]MonitorSystem, error)
	// GetByID gets a monitor system by the identity from the middleware
	GetByID(id int) (MonitorSystem, error)
	// GetByHostInfo gets a monitor system with given host ip and port number
	GetByHostInfo(hostIP string, portNum int) (MonitorSystem, error)
	// GetID gets the identity with given host ip and port number from the middleware
	GetID(hostIP string, portNum int) (int, error)
	// Create creates a monitor system in the middleware
	Create(ms MonitorSystem) (MonitorSystem, error)
	// Update updates the monitor system in the middleware
	Update(ms MonitorSystem) error
	// Delete deletes the monitor system in the middleware
	Delete(id int) error
}

type MonitorSystemService interface {
	// GetDBs returns monitor systems of the service
	GetMonitorSystems() []MonitorSystem
	// GetAll gets all monitor systems from the middleware
	GetAll() error
	// GetByEnv gets monitor systems of given env id
	GetByEnv(envID int) error
	// GetByID gets a monitor system of the given id from the middleware
	GetByID(id int) error
	// GetByHostInfo gets a monitor system with given host ip and port number
	GetByHostInfo(hostIP string, portNum int) error
	// Create creates a monitor system in the middleware
	Create(fields map[string]interface{}) error
	// Update gets a monitor system of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the monitor system of given id in the middleware
	Delete(id int) error
	// Marshal marshals MonitorSystemService.MonitorSystems to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the MonitorSystemService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
