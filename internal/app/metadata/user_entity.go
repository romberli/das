package metadata

import (
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
)

var _ metadata.User = (*UserInfo)(nil)

// UserInfo create userinfo struct
type UserInfo struct {
	metadata.UserRepo
	ID             int       `middleware:"id" json:"id"`
	UserName       string    `middleware:"user_name" json:"user_name"`
	DepartmentName string    `middleware:"department_name" json:"department_name"`
	EmployeeID     string    `middleware:"employee_id" json:"employee_id"`
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
func NewUserInfo(repo metadata.UserRepo, id int, userName string, departmentName string, employeeID string, accountName string, email string, telephone string, mobile string, role int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *UserInfo {
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

// NewUserInfoWithGlobal returns a new UserInfo with default UserRepo
func NewUserInfoWithGlobal(id int, userName string, departmentName string, employeeID string, accountName string, email string, telephone string, mobile string, role int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *UserInfo {
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

// NewEmptyUserInfoWithGlobal return userinfo
func NewEmptyUserInfoWithGlobal() *UserInfo {
	return &UserInfo{UserRepo: NewUserRepoWithGlobal()}
}

// NewUserInfoWithDefault returns a new UserInfo with default UserRepo
func NewUserInfoWithDefault(userName string, departmentName string, accountName string, email string, role int) *UserInfo {
	return &UserInfo{
		UserRepo:       NewUserRepoWithGlobal(),
		UserName:       userName,
		DepartmentName: departmentName,
		EmployeeID:     constant.DefaultRandomString,
		AccountName:    accountName,
		Email:          email,
		Telephone:      constant.DefaultRandomString,
		Mobile:         constant.DefaultRandomString,
		Role:           role,
	}
}

// NewUserInfoWithMapAndRandom returns a new *UserInfoInfo with given map
func NewUserInfoWithMapAndRandom(fields map[string]interface{}) (*UserInfo, error) {
	ui := &UserInfo{}
	err := common.SetValuesWithMapAndRandom(ui, fields)
	if err != nil {
		return nil, err
	}

	return ui, nil
}

// Identity returns the identity
func (ui *UserInfo) Identity() int {
	return ui.ID
}

// // IsDeleted checks if delete flag had been set
// func (ui *UserInfo) IsDeleted() bool {
// 	return ui.DelFlag != constant.ZeroInt
// }

// GetCreateTime returns the create time
func (ui *UserInfo) GetCreateTime() time.Time {
	return ui.CreateTime
}

// GetLastUpdateTime returns the last update time
func (ui *UserInfo) GetLastUpdateTime() time.Time {
	return ui.LastUpdateTime
}

// GetDepartmentName returns the department name
func (ui *UserInfo) GetDepartmentName() string {
	return ui.DepartmentName
}

// GetMobile returns the mobile
func (ui *UserInfo) GetMobile() string {
	return ui.Mobile
}

// GetUserName returns the user name
func (ui *UserInfo) GetUserName() string {
	return ui.UserName
}

// GetEmployeeID returns the employee id
func (ui *UserInfo) GetEmployeeID() string {
	return ui.EmployeeID
}

// GetAccountName returns the account name
func (ui *UserInfo) GetAccountName() string {
	return ui.AccountName
}

// GetEmail returns the email
func (ui *UserInfo) GetEmail() string {
	return ui.Email
}

// GetEmail returns the telephone
func (ui *UserInfo) GetTelephone() string {
	return ui.Telephone
}

// GetRole returns the role
func (ui *UserInfo) GetRole() int {
	return ui.Role
}

// GetDelFlag returns the delete flag
func (ui *UserInfo) GetDelFlag() int {
	return ui.DelFlag
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

// Delete sets DelFlag to 1
func (ui *UserInfo) Delete() {
	ui.DelFlag = 1
}

// MarshalJSON marshals User to json string
func (ui *UserInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(ui, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only specified field of the User to json string
func (ui *UserInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ui, fields...)
}
