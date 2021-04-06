package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

// MySQLCluster is the entity interface
type MySQLCluster interface {
	// Identity returns the identity
	Identity() int
	// GetClusterName returns the env name
	GetClusterName() string
	// GetMiddlewareClusterID returns the middleware cluster id
	GetMiddlewareClusterID() int
	// GetMonitorSystemID returns the monitor system id
	GetMonitorSystemID() int
	// GetOwnerID returns the owner id
	GetOwnerID() int
	// GetEnvID returns the env id
	GetEnvID() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// GetMySQLServerIDList gets the mysql server id list of this cluster
	GetMySQLServerIDList() ([]int, error)
	// Set sets MySQLCluster with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// MarshalJSON marshals MySQLCluster to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the MySQLCluster to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

// MySQLClusterRepo is the repository interface
type MySQLClusterRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all mysql clusters from the middleware
	GetAll() ([]MySQLCluster, error)
	// GetByEnv gets mysql clusters of given env id from the middleware
	GetByEnv(envID int) ([]MySQLCluster, error)
	// GetByID gets a mysql cluster by the identity from the middleware
	GetByID(id int) (MySQLCluster, error)
	// GetByName gets a mysql cluster of given cluster name from the middle ware
	GetByName(clusterName string) (MySQLCluster, error)
	// GetID gets the identity with given cluster name from the middleware
	GetID(clusterName string) (int, error)
	// GetMySQLServerIDList gets the mysql server id list of given cluster id
	GetMySQLServerIDList(clusterID int) ([]int, error)
	// Create creates a mysql cluster in the middleware
	Create(mc MySQLCluster) (MySQLCluster, error)
	// Update updates the mysql cluster in the middleware
	Update(mc MySQLCluster) error
	// Delete deletes the mysql cluster in the middleware
	Delete(id int) error
}

// MySQLClusterService is the service interface
type MySQLClusterService interface {
	// GetMySQLClusters returns mysql clusters of the service
	GetMySQLClusters() []MySQLCluster
	// GetAll gets all mysql clusters from the middleware
	GetAll() error
	// GetByEnv gets mysql clusters of given env id
	GetByEnv(envID int) error
	// GetByID gets a mysql cluster of the given id from the middleware
	GetByID(id int) error
	// GetByName gets a mysql cluster of given cluster name
	GetByName(clusterName string) error
	// GetMySQLServerIDList gets the mysql server id list of given cluster id
	GetMySQLServerIDList(clusterID int) error
	// Create creates a mysql cluster in the middleware
	Create(fields map[string]interface{}) error
	// Update gets a mysql cluster of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the mysql cluster of given id in the middleware
	Delete(id int) error
	// Marshal marshals MySQLClusterService.MySQLClusters to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the MySQLClusterService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
