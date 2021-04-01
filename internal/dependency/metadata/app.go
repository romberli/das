package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type App interface {
	// Identity returns the identity
	Identity() int
	// GetSystemName returns the app name
	GetAppName() string
	// GetLevel returns the level
	GetLevel() int
	// GetOwnerID returns the owner id
	GetOwnerID() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// GetDBIDList gets database identity list that the app uses
	GetDBIDList() ([]int, error)
	// Set sets App with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// AddDB adds a new map of the app and database in the middleware
	AddDB(dbID int) error
	// DeleteDB deletes the map of the app and database in the middleware
	DeleteDB(dbID int) error
	// MarshalJSON marshals App to json bytes
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified fields of the App to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type AppRepo interface {
	// Execute executes command with arguments on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all apps from the middleware
	GetAll() ([]App, error)
	// GetByID gets an app by the identity from the middleware
	GetByID(id int) (App, error)
	// GetID gets the identity with given app name from the middleware
	GetID(appName string) (int, error)
	// GetAppSystemByName gets the app by name from the middleware
	GetAppByName(appName string) (App, error)
	// GetDBIDList gets a database identity list that app uses
	GetDBIDList(id int) ([]int, error)
	// Create creates an app in the middleware
	Create(appSystem App) (App, error)
	// Update updates the app in the middleware
	Update(appSystem App) error
	// Delete deletes the app in the middleware
	Delete(id int) error
	// AddDB adds a new map of app and database in the middleware
	AddDB(appID, dbID int) error
	// DeleteDB delete the map of app and database in the middleware
	DeleteDB(appID, dbID int) error
}

type AppService interface {
	// GetApps returns apps of the service
	GetApps() []App
	// GetAll gets all apps from the middleware
	GetAll() error
	// GetByID gets an app of the given id from the middleware
	GetByID(id int) error
	// GetAppByName gets App from the middleware by name
	GetAppByName(appName string) error
	// GetDBIDList gets a database identity list that the app uses
	GetDBIDList(id int) error
	// Create creates an app in the middleware
	Create(fields map[string]interface{}) error
	// Update gets the app of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the app of given id in the middleware
	Delete(id int) error
	// AddDB adds a new map of app and database in the middleware
	AddDB(appID, dbID int) error
	// DeleteDB deletes the map of app and database in the middleware
	DeleteDB(appID, dbID int) error
	// Marshal marshals AppService.Apps to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the AppService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
