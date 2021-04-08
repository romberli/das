package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type User interface {
	// Identity returns the identity
	Identity() int
	// GetUserName returns the user name
	GetUserName() string
	// GetDepartmentName returns the department name
	GetDepartmentName() string
	// GetEmployeeID returns the employee id
	GetEmployeeID() string
	// GetAccountName returns the account name
	GetAccountName() string
	// GetEmail returns the email
	GetEmail() string
	// GetEmail returns the telephone
	GetTelephone() string
	// GetMobile returns the mobile
	GetMobile() string
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// Set sets User with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// MarshalJSON marshals User to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the User to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type UserRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all databases from the middleware
	GetAll() ([]User, error)
	// GetByName gets users of given user name from the middleware
	GetByName(userName string) ([]User, error)
	// GetByID gets a user by the identity from the middleware
	GetByID(id int) (User, error)
	// GetByAccountName gets a user of given account name from the middleware
	GetByAccountName(accountName string) (User, error)
	// GetByEmail gets a user of given email from the middleware
	GetByEmail(email string) (User, error)
	// GetByTelephone gets a user of given telephone from the middleware
	GetByTelephone(telephone string) (User, error)
	// GetByTelephone gets a user of given mobile from the middleware
	GetByMobile(mobile string) (User, error)
	// GetID gets the identity with given accountName from the middleware
	GetID(accountName string) (int, error)
	// Create creates a user in the middleware
	Create(db User) (User, error)
	// Update updates a user in the middleware
	Update(db User) error
	// Delete deletes a user in the middleware
	Delete(id int) error
	// GetByEmployeeID gets a user of given employee id from the middleware
	GetByEmployeeID(employeeID string) (User, error)
}

type UserService interface {
	// GetUsers returns users of the service
	GetUsers() []User
	// GetAll gets all users
	GetAll() error
	// GetByName gets users of given user name
	GetByName(userName string) error
	// GetByID gets a user by the identity
	GetByID(id int) error
	// GetByAccountName gets a user of given account name
	GetByAccountName(accountName string) error
	// GetByEmail gets a user of given email
	GetByEmail(email string) error
	// GetByTelephone gets a user of given telephone
	GetByTelephone(telephone string) error
	// GetByTelephone gets a user of given mobile
	GetByMobile(mobile string) error
	// Create creates a user in the middleware
	Create(fields map[string]interface{}) error
	// Update gets a user of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the user of given id in the middleware
	Delete(id int) error
	// Marshal marshals UserService.Users to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the UserService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
	// GetByEmployeeID gets a user of given employee id
	GetByEmployeeID(employeeID string) error
}
