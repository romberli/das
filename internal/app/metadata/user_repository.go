package metadata

import (
	"errors"
	"fmt"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"

	"github.com/romberli/log"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/metadata"
)

var _ metadata.UserRepo = (*UserRepo)(nil)

// UserRepo struct
type UserRepo struct {
	Database middleware.Pool
}

// NewUserRepo returns *UserRepo with given middleware.Pool
func NewUserRepo(db middleware.Pool) *UserRepo {
	return &UserRepo{db}
}

// NewUserRepoWithGlobal returns *UserRepo with global mysql pool
func NewUserRepoWithGlobal() *UserRepo {
	return NewUserRepo(global.MySQLPool)
}

// Execute implements dependency.Repository interface,
// it executes command with arguments on database
func (ur *UserRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := ur.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata UserRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// GetByName get userinfo by userName
func (ur *UserRepo) GetByName(userName string) ([]metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	and user_name = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, userName)

	result, err := ur.Execute(sql, userName)
	if err != nil {
		return nil, err
	}
	// init []*UserInfo
	userInfoList := make([]*UserInfo, result.RowNumber())
	for i := range userInfoList {
		userInfoList[i] = NewEmptyUserInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(userInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.User
	userList := make([]metadata.User, result.RowNumber())
	for i := range userList {
		userList[i] = userInfoList[i]
	}

	return userList, nil

}

// Transaction ..
func (ur *UserRepo) Transaction() (middleware.Transaction, error) {
	return ur.Database.Transaction()
}

// GetAll returns all available entities
func (ur *UserRepo) GetAll() ([]metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	order by id;
	`
	log.Debugf("metadata UserRepo.GetAll() sql: \n%s", sql)

	result, err := ur.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*UserInfo
	userInfoList := make([]*UserInfo, result.RowNumber())
	for i := range userInfoList {
		userInfoList[i] = NewEmptyUserInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(userInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.User
	userList := make([]metadata.User, result.RowNumber())
	for i := range userList {
		userList[i] = userInfoList[i]
	}

	return userList, nil
}

// GetByMobile get userinfo by mobile
func (ur *UserRepo) GetByMobile(mobile string) (metadata.User, error) {
	sql := `
		select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
		from t_meta_user_info
		where del_flag = 0
		and mobile = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, mobile)

	result, err := ur.Execute(sql, mobile)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByMobile(): data does not exists, id: %s", mobile))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByMobile(): duplicate key exists, id: %s", mobile))
	}
}

// GetByTelephone get userinfo by telephone
func (ur *UserRepo) GetByTelephone(telephone string) (metadata.User, error) {
	sql := `
		select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
		from t_meta_user_info
		where del_flag = 0
		and telephone = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, telephone)

	result, err := ur.Execute(sql, telephone)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByTelephone(): data does not exists, id: %s", telephone))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByTelephone(): duplicate key exists, id: %s", telephone))
	}
}

// GetByID get userinfo by id
func (ur *UserRepo) GetByID(id int) (metadata.User, error) {
	sql := `
		select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
		from t_meta_user_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := ur.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByID(): data does not exists, id: %s", id))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByID(): duplicate key exists, id: %s", id))
	}
}

// GetByAccountName get userinfo by accountname
func (ur *UserRepo) GetByAccountName(accountName string) (metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	and account_name = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, accountName)

	result, err := ur.Execute(sql, accountName)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByAccountName(): data does not exists, id: %s", accountName))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByAccountName(): duplicate key exists, id: %s", accountName))
	}
}

// GetByEmployeeID get userinfo by employeeID
func (ur *UserRepo) GetByEmployeeID(employeeID string) (metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	and employee_id = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, employeeID)

	result, err := ur.Execute(sql, employeeID)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByEmployeeID(): data does not exists, id: %s", employeeID))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByEmployeeID(): duplicate key exists, id: %s", employeeID))
	}
}

// GetByEmail get userinfo by email
func (ur *UserRepo) GetByEmail(email string) (metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	and email = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, email)

	result, err := ur.Execute(sql, email)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByEmail(): data does not exists, id: %s", email))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByEmail(): duplicate key exists, id: %s", email))
	}
}

// GetID checks iduser of given user from the middleware
func (ur *UserRepo) GetID(employeeID string) (int, error) {
	sql := `select id from t_meta_user_info where del_flag = 0 and employee_id = ?;`
	log.Debugf("metadata UserRepo.GetID() select sql: %s", sql)
	result, err := ur.Execute(sql, employeeID)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given user in the middleware
func (ur *UserRepo) Create(user metadata.User) (metadata.User, error) {
	sql := `insert into t_meta_user_info(user_name, department_name, employee_id, account_name, email , telephone , mobile, role) values(?,?,?,?,?,?,?,?);`
	log.Debugf("metadata UserRepo.Create() insert sql: %s", sql)
	// execute
	userInfo := user.(*UserInfo)
	_, err := ur.Execute(sql, userInfo.UserName, userInfo.DepartmentName, userInfo.EmployeeID, userInfo.AccountName, userInfo.Email, userInfo.Telephone, userInfo.Mobile, userInfo.Role)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := ur.GetID(user.GetEmployeeID())
	if err != nil {
		return nil, err
	}
	// get user
	return ur.GetByID(id)
}

// Update updates data with given user in the middleware
func (ur *UserRepo) Update(user metadata.User) error {
	sql := `update t_meta_user_info set user_name = ?, del_flag = ?, department_name = ?, employee_id = ?, account_name = ?, email = ?, telephone = ?, mobile = ?, role = ? where id = ?;`
	log.Debugf("metadata UserRepo.Update() update sql: %s", sql)
	userInfo := user.(*UserInfo)
	_, err := ur.Execute(sql, userInfo.UserName, userInfo.DelFlag, userInfo.DepartmentName, userInfo.EmployeeID, userInfo.AccountName, userInfo.Email, userInfo.Telephone, userInfo.Mobile, userInfo.Role, userInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (ur *UserRepo) Delete(id int) error {
	sql := `update t_meta_user_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata UserRepo.Delete() update sql: %s", sql)

	_, err := ur.Execute(sql, id)

	return err
}
