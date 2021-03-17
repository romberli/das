package metadata

import (
	"encoding/json"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/pkg/message"

	"github.com/romberli/das/internal/dependency"
)

const (
	userNameStruct       = "UserName"
	departmentNameStruct = "DepartmentName"
	employeeIDStruct     = "EmployeeID"
	accountNameStruct    = "accountName"
	emailStruct          = "Email"
	telephoneStruct      = "Telephone"
	mobileStruct         = "Mobile"
	roleStruct           = "Role"
)

var _ dependency.Service = (*UserService)(nil)

type UserService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewUserService returns a new *UserService
func NewUserService(repo dependency.Repository) *UserService {
	return &UserService{repo, []dependency.Entity{}}
}

// NewUserServiceWithDefault returns a new *UserService with default repository
func NewUserServiceWithDefault() *UserService {
	return NewUserService(NewUserRepoWithGlobal())
}

// GetEntities returns entities of the service
func (us *UserService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(us.Entities))
	for i := range entityList {
		entityList[i] = us.Entities[i]
	}

	return entityList
}

// GetAll gets all Userironment entities from the middleware
func (us *UserService) GetAll() error {
	var err error
	us.Entities, err = us.Repository.GetAll()

	return err
}

// GetByID gets an Userironment entity that contains the given id from the middleware
func (us *UserService) GetByID(id string) error {
	entity, err := us.Repository.GetByID(id)
	if err != nil {
		return err
	}

	us.Entities = append(us.Entities, entity)

	return err
}

// Create creates a new Userironment entity and insert it into the middleware
func (us *UserService) Create(fields map[string]interface{}) error {
	// generate new map
	userName, ok := fields[userNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, userNameStruct)
	}
	departmentName, ok := fields[departmentNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, departmentName)
	}
	employeeID, ok := fields[employeeIDStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, employeeIDStruct)
	}
	accountName, ok := fields[accountNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, accountNameStruct)
	}
	email, ok := fields[emailStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, emailStruct)
	}
	telephone, ok := fields[telephoneStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, telephoneStruct)
	}
	mobile, ok := fields[mobileStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, mobileStruct)
	}
	role, ok := fields[roleStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, roleStruct)
	}

	userInfo := NewUserInfoWithDefault(userName.(string), departmentName.(string), employeeID.(int), accountName.(string), email.(string), telephone.(string), mobile.(string), role.(int))
	// insert into middleware
	entity, err := us.Repository.Create(userInfo)
	if err != nil {
		return err
	}

	us.Entities = append(us.Entities, entity)
	return nil
}

// Update gets an Userironment entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (us *UserService) Update(id string, fields map[string]interface{}) error {
	err := us.GetByID(id)
	if err != nil {
		return err
	}
	err = us.Entities[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return us.Repository.Update(us.Entities[constant.ZeroInt])
}

// Delete deletes the Userironment entity that contains the given id in the middleware
func (us *UserService) Delete(id string) error {
	return us.Repository.Delete(id)
}

// Marshal marshals service.Entities
func (us *UserService) Marshal() ([]byte, error) {
	return json.Marshal(us.Entities)
}

// Marshal marshals service.Entities with given fields
func (us *UserService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(us.Entities))
	for i := range interfaceList {
		entity, err := common.CopyStructWithFields(us.Entities[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = entity
	}

	return json.Marshal(interfaceList)
}
