package metadata

import (
	"encoding/json"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/pkg/message"

	"github.com/romberli/das/internal/dependency/metadata"
)

var _ metadata.UserService = (*UserService)(nil)

const (
	userNameStruct       = "UserName"
	departmentNameStruct = "DepartmentName"
	employeeIDStruct     = "EmployeeID"
	accountNameStruct    = "AccountName"
	emailStruct          = "Email"
	telephoneStruct      = "Telephone"
	mobileStruct         = "Mobile"
	roleStruct           = "Role"
)

// UserService struct
type UserService struct {
	metadata.UserRepo
	Users []metadata.User
}

// NewUserService returns a new *UserService
func NewUserService(repo metadata.UserRepo) *UserService {
	return &UserService{repo, []metadata.User{}}
}

// NewUserServiceWithDefault returns a new *UserService with default repository
func NewUserServiceWithDefault() *UserService {
	return NewUserService(NewUserRepoWithGlobal())
}

// GetAll gets all user entities from the middleware
func (us *UserService) GetAll() error {
	var err error
	us.Users, err = us.UserRepo.GetAll()

	return err
}

// GetByID gets an user user that contains the given id from the middleware
func (us *UserService) GetByID(id int) error {
	user, err := us.UserRepo.GetByID(id)
	if err != nil {
		return err
	}

	us.Users = append(us.Users, user)

	return err
}

// Create creates a new user user and insert it into the middleware
func (us *UserService) Create(fields map[string]interface{}) error {
	// generate new map
	_, ok := fields[userNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, userNameStruct)
	}
	_, ok = fields[departmentNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, departmentNameStruct)
	}
	_, ok = fields[employeeIDStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, employeeIDStruct)
	}
	_, ok = fields[accountNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, accountNameStruct)
	}
	_, ok = fields[emailStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, emailStruct)
	}

	_, ok = fields[roleStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, roleStruct)
	}

	// create a new user
	userInfo, err := NewUserInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}

	// insert into middleware
	user, err := us.UserRepo.Create(userInfo)
	if err != nil {
		return err
	}

	us.Users = append(us.Users, user)
	return nil
}

// Update gets an user user that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (us *UserService) Update(id int, fields map[string]interface{}) error {
	err := us.GetByID(id)
	if err != nil {
		return err
	}
	err = us.Users[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return us.UserRepo.Update(us.Users[constant.ZeroInt])
}

// Delete deletes the user user that contains the given id in the middleware
func (us *UserService) Delete(id int) error {
	return us.UserRepo.Delete(id)
}

// Marshal marshals service.Users
func (us *UserService) Marshal() ([]byte, error) {
	return json.Marshal(us.Users)
}

// MarshalWithFields marshals service.Users with given fields
func (us *UserService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(us.Users))
	for i := range interfaceList {
		user, err := common.CopyStructWithFields(us.Users[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = user
	}

	return json.Marshal(interfaceList)
}

// GetByAccountName gets an userinfo that contains the given accountname from the middleware
func (us *UserService) GetByAccountName(accountName string) error {
	user, err := us.UserRepo.GetByAccountName(accountName)
	if err != nil {
		return err
	}

	us.Users = append(us.Users, user)

	return err
}

// GetByEmail gets an userinfo that contains the given email from the middleware
func (us *UserService) GetByEmail(email string) error {
	user, err := us.UserRepo.GetByEmail(email)
	if err != nil {
		return err
	}

	us.Users = append(us.Users, user)

	return err
}

// GetByTelephone gets an userinfo that contains the given telephone from the middleware
func (us *UserService) GetByTelephone(telephone string) error {
	user, err := us.UserRepo.GetByTelephone(telephone)
	if err != nil {
		return err
	}

	us.Users = append(us.Users, user)

	return err
}

// GetByMobile gets an userinfo that contains the given mobile from the middleware
func (us *UserService) GetByMobile(mobile string) error {
	user, err := us.UserRepo.GetByMobile(mobile)
	if err != nil {
		return err
	}

	us.Users = append(us.Users, user)

	return err
}



// GetByName gets an userinfo that contains the given userName from the middleware
func (us *UserService) GetByName(userName string) error {
	var err error
	us.Users, err = us.UserRepo.GetByName(userName)

	return err
}

// GetUsers() return users of the service
func (us *UserService) GetUsers() []metadata.User {
	return us.Users
}
