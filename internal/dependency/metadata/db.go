package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type DB interface {
	// Identity returns the identity
	Identity() int
	// GetDBName returns the db name
	GetDBName() string
	// GetClusterID returns the cluster id
	GetClusterID() int
	// GetClusterType returns the cluster type
	GetClusterType() int
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
	// GetAppIDList gets app identity list that uses this db
	GetAppIDList() ([]int, error)
	// Set sets DB with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// AddDB adds a new map of the app and database in the middleware
	AddApp(appID int) error
	// DeleteApp deletes a new map of the app and database in the middleware
	DeleteApp(appID int) error
	// MarshalJSON marshals DB to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the DB to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type DBRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all databases from the middleware
	GetAll() ([]DB, error)
	// GetByEnv gets databases of given env id from the middleware
	GetByEnv(envID int) ([]DB, error)
	// GetByID gets a database by the identity from the middleware
	GetByID(id int) (DB, error)
	// GetByNameAndClusterInfo gets a database by the db name and cluster info from the middleware
	GetByNameAndClusterInfo(name string, clusterID, clusterType int) (DB, error)
	// GetID gets the identity with given database name, cluster id and cluster type from the middleware
	GetID(dbName string, clusterID int, clusterType int) (int, error)
	// GetAppIDList gets an app identity list that uses this db
	GetAppIDList(id int) ([]int, error)
	// Create creates a database in the middleware
	Create(db DB) (DB, error)
	// Update updates the database in the middleware
	Update(db DB) error
	// Delete deletes the database in the middleware
	Delete(id int) error
	// AddApp adds a new map of the app and database in the middleware
	AddApp(dbID, appID int) error
	// DeleteApp deletes a map of the app and database in the middleware
	DeleteApp(dbID, appID int) error
}

type DBService interface {
	// GetDBs returns databases of the service
	GetDBs() []DB
	// GetAll gets all databases from the middleware
	GetAll() error
	// GetByEnv gets databases of given env id
	GetByEnv(envID int) error
	// GetByID gets a database of the given id from the middleware
	GetByID(id int) error
	// GetByNameAndClusterInfo gets an database of the given db name and cluster info from the middleware
	GetByNameAndClusterInfo(name string, clusterID, clusterType int) error
	// GetAppIDList gets an app identity list that uses this db
	GetAppIDList(id int) error
	// Create creates a database in the middleware
	Create(fields map[string]interface{}) error
	// Update gets a database of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the database of given id in the middleware
	Delete(id int) error
	// AddApp adds a new map of app and database in the middleware
	AddApp(dbID, appID int) error
	// DeleteApp deletes the map of app and database in the middleware
	DeleteApp(dbID, appID int) error
	// Marshal marshals DBService.DBs to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the DBService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
