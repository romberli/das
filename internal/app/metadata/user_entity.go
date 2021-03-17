package metadata

import (
	"strconv"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Entity = (*UserInfo)(nil)

type UserInfo struct {
	dependency.Repository
	ID             int       `middleware:"id" json:"id"`
	UserName       string    `middleware:"user_name" json:"user_name"`
	DepartmentName string    `middleware:"department_name" json:"department_name"`
	EmployeeID     int       `middleware:"employee_id" json:"employee_id"`
	AccountName    string    `middleware:"account_name" json:"account_name"`
	Email          string    `middleware:"email" json:"email"`
	Telephone      string    `middleware:"telephone" json:"telephone"`
	Mobile         string    `middleware:"mobile" json:"mobile"`
	Role           int       `middleware:"role" json:"role"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewUserInfo returns a new UserInfo
func NewUserInfo(repo *UserRepo, id int, userName string, departmentName string, employeeID int, accountName string, email string, telephone string, mobile string, role int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *UserInfo {
	return &UserInfo{
		repo,
		id,
		userName,
		departmentName,
		employeeID,
		accountName,
		email,
		telephone,
		mobile,
		role,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewUserInfo returns a new UserInfo with default UserRepo
func NewUserInfoWithGlobal(id int, userName string, delFlag int, createTime time.Time, lastUpdateTime time.Time, departmentName string, employeeID int, accountName string, email string, telephone string, mobile string, role int) *UserInfo {
	return &UserInfo{
		NewUserRepoWithGlobal(),
		id,
		userName,
		departmentName,
		employeeID,
		accountName,
		email,
		telephone,
		mobile,
		role,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

func NewEmptyUserInfoWithGlobal() *UserInfo {
	return &UserInfo{Repository: NewUserRepoWithGlobal()}
}

// NewUserInfoWithDefault returns a new UserInfo with default UserRepo
func NewUserInfoWithDefault(userName string, departmentName string, employeeID int, accountName string, email string, telephone string, mobile string, role int) *UserInfo {
	return &UserInfo{
		Repository:     NewUserRepoWithGlobal(),
		UserName:       userName,
		DepartmentName: departmentName,
		EmployeeID:     employeeID,
		AccountName:    accountName,
		Email:          email,
		Telephone:      telephone,
		Mobile:         mobile,
		Role:           role,
	}
}

// Identity returns ID of entity
func (ui *UserInfo) Identity() string {
	return strconv.Itoa(ui.ID)
}

// IsDeleted checks if delete flag had been set
func (ui *UserInfo) IsDeleted() bool {
	return ui.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (ui *UserInfo) GetCreateTime() time.Time {
	return ui.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (ui *UserInfo) GetLastUpdateTime() time.Time {
	return ui.LastUpdateTime
}

// GetDepartmentName returns last updated time of entity
func (ui *UserInfo) GetDepartmentName() string {
	return ui.DepartmentName
}

// GetEmployeeID returns last updated time of entity
func (ui *UserInfo) GetEmployeeID() int {
	return ui.EmployeeID
}

// GetDomainAccount returns last updated time of entity
func (ui *UserInfo) GetDomainAccount() string {
	return ui.AccountName
}

// GetEmail returns last updated time of entity
func (ui *UserInfo) GetEmail() string {
	return ui.Email
}

// GetTelephone returns last updated time of entity
func (ui *UserInfo) GetTelephone() string {
	return ui.Telephone
}

// GetRole returns last updated time of entity
func (ui *UserInfo) GetRole() int {
	return ui.Role
}

// Get returns value of given field
func (ui *UserInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(ui, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (ui *UserInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(ui, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (ui *UserInfo) Delete() {
	ui.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (ui *UserInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(ui, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (ui *UserInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ui, fields...)
}
