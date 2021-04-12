package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type MiddlewareServer interface {
	// Identity returns the identity
	Identity() int
	// GetClusterID returns the middleware cluster id
	GetClusterID() int
	// GetServerName returns the server name
	GetServerName() string
	// GetMiddlewareRole returns the middleware role
	GetMiddlewareRole() int
	// GetHostIP returns the host ip
	GetHostIP() string
	// GetPortNum returns the port number
	GetPortNum() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// Set sets MiddlewareServer with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// MarshalJSON marshals MiddlewareServer to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the MiddlewareServer to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type MiddlewareServerRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all middleware servers from the middleware
	GetAll() ([]MiddlewareServer, error)
	// GetByClusterID gets middleware servers with given cluster id
	GetByClusterID(clusterID int) ([]MiddlewareServer, error)
	// GetByID gets a middleware server by the identity from the middleware
	GetByID(id int) (MiddlewareServer, error)
	// GetByHostInfo gets a middleware server with given host ip and port number
	GetByHostInfo(hostIP string, portNum int) (MiddlewareServer, error)
	// GetID gets the identity with given host ip and port number from the middleware
	GetID(hostIP string, portNum int) (int, error)
	// Create creates a middleware server in the middleware
	Create(ms MiddlewareServer) (MiddlewareServer, error)
	// Update updates the middleware server in the middleware
	Update(ms MiddlewareServer) error
	// Delete deletes the middleware server in the middleware
	Delete(id int) error
}

type MiddlewareServerService interface {
	// GetMiddlewareServers returns middleware servers of the service
	GetMiddlewareServers() []MiddlewareServer
	// GetAll gets all middleware servers from the middleware
	GetAll() error
	// GetByClusterID gets middleware servers with given cluster id
	GetByClusterID(clusterID int) error
	// GetByID gets a middleware server of the given id from the middleware
	GetByID(id int) error
	// GetByHostInfo gets a middleware server with given host ip and port number
	GetByHostInfo(hostIP string, portNum int) error
	// Create creates a middleware server in the middleware
	Create(fields map[string]interface{}) error
	// Update gets a middleware server of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the middleware server of given id in the middleware
	Delete(id int) error
	// Marshal marshals MiddlewareServerService.MiddlewareServers to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the MiddlewareServerService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
