package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

// MySQLServer is the entity interface
type MySQLServer interface {
	// Identity returns the identity
	Identity() int
	// GetClusterID returns the mysql cluster id
	GetClusterID() int
	// GetServerName returns the server name
	GetServerName() string
	// GetServiceName returns the service name
	GetServiceName() string
	// GetHostIP returns the host ip
	GetHostIP() string
	// GetPortNum returns the port number
	GetPortNum() int
	// GetDeploymentType returns the deployment type
	GetDeploymentType() int
	// GetVersion returns the version
	GetVersion() string
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// GetMonitorSystem gets monitor system from the mysql
	GetMonitorSystem() (MonitorSystem, error)
	// Set sets MySQLServer with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// MarshalJSON marshals MySQLServer to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the MySQLServer to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

// MySQLServerRepo is the repository interface
type MySQLServerRepo interface {
	// Execute executes given command and placeholders on the mysql
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a mysql.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all mysql servers from the mysql
	GetAll() ([]MySQLServer, error)
	// GetByClusterID gets mysql servers with given cluster id
	GetByClusterID(clusterID int) ([]MySQLServer, error)
	// GetByID gets a mysql server by the identity from the mysql
	GetByID(id int) (MySQLServer, error)
	// GetByHostInfo gets a mysql server with given host ip and port number
	GetByHostInfo(hostIP string, portNum int) (MySQLServer, error)
	// GetID gets the identity with given host ip and port number from the mysql
	GetID(hostIP string, portNum int) (int, error)
	// GetMonitorSystem gets monitor system with given mysql server id from the mysql
	GetMonitorSystem(id int) (MonitorSystem, error)
	// Create creates a mysql server in the mysql
	Create(ms MySQLServer) (MySQLServer, error)
	// Update updates the mysql server in the mysql
	Update(ms MySQLServer) error
	// Delete deletes the mysql server in the mysql
	Delete(id int) error
}

// MySQLServerService is the service interface
type MySQLServerService interface {
	// GetMySQLServers returns mysql servers of the service
	GetMySQLServers() []MySQLServer
	// GetAll gets all mysql servers from the mysql
	GetAll() error
	// GetByClusterID gets mysql servers with given cluster id
	GetByClusterID(clusterID int) error
	// GetByID gets a mysql server of the given id from the mysql
	GetByID(id int) error
	// GetByHostInfo gets a mysql server with given host ip and port number
	GetByHostInfo(hostIP string, portNum int) error
	// Create creates a mysql server in the mysql
	Create(fields map[string]interface{}) error
	// Update gets a mysql server of the given id from the mysql,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the mysql
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the mysql server of given id in the mysql
	Delete(id int) error
	// Marshal marshals MySQLServerService.MySQLServers to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the MySQLServerService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
