package metadata

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"

	"github.com/romberli/log"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Repository = (*UserRepo)(nil)

type UserRepo struct {
	Database middleware.Pool
}

// NewUserRepo returns *UserRepo with given middleware.Pool
func NewUserRepo(db middleware.Pool) *UserRepo {
	return &UserRepo{db}
}

// NewUserRepo returns *UserRepo with global mysql pool
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

func (ur *UserRepo) Transaction() (middleware.Transaction, error) {
	return ur.Database.Transaction()
}

// GetAll returns all available entities
func (ur *UserRepo) GetAll() ([]dependency.Entity, error) {
	sql := `
		select id, user_name, del_flag, create_time, last_update_time ,department_name,employee_id,domin_account,email,telephone,mobile,role
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
	// init []dependency.Entity
	entityList := make([]dependency.Entity, result.RowNumber())
	for i := range entityList {
		entityList[i] = userInfoList[i]
	}

	return entityList, nil
}

func (ur *UserRepo) GetByID(id string) (dependency.Entity, error) {
	sql := `
		select id, user_name, del_flag, create_time, last_update_time, department_name,employee_id,domin_account,email,telephone,mobile,role
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

// GetID checks identity of given entity from the middleware
func (ur *UserRepo) GetID(entity dependency.Entity) (string, error) {
	sql := `select id from t_meta_user_info where del_flag = 0 and user_name = ?;`
	log.Debugf("metadata UserRepo.GetID() select sql: %s", sql)
	result, err := ur.Execute(sql, entity.(*UserInfo).UserName)
	if err != nil {
		return constant.EmptyString, err
	}

	return result.GetString(constant.ZeroInt, constant.ZeroInt)
}

// Create creates data with given entity in the middleware
func (ur *UserRepo) Create(entity dependency.Entity) (dependency.Entity, error) {
	sql := `insert into t_meta_user_info(user_name,department_name, employee_id, domain_account, email , telephone , mobile, role) values(?,?,?,?,?,?,?,?);`
	log.Debugf("metadata UserRepo.Create() insert sql: %s", sql)
	// execute
	_, err := ur.Execute(sql, entity.(*UserInfo).UserName, entity.(*UserInfo).DepartmentName, entity.(*UserInfo).EmployeeID, entity.(*UserInfo).DomainAccount, entity.(*UserInfo).Email, entity.(*UserInfo).Telephone, entity.(*UserInfo).Mobile, entity.(*UserInfo).Role)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := ur.GetID(entity)
	if err != nil {
		return nil, err
	}
	// get entity
	return ur.GetByID(id)
}

// Update updates data with given entity in the middleware
func (ur *UserRepo) Update(entity dependency.Entity) error {
	sql := `update t_meta_user_info set user_name = ?, del_flag = ?, department_name = ?, employee_id = ?, domain_account = ?, email = ?, telephone = ?, mobile = ?, role = ? where id = ?;`
	log.Debugf("metadata UserRepo.Update() update sql: %s", sql)
	userInfo := entity.(*UserInfo)
	_, err := ur.Execute(sql, userInfo.UserName, userInfo.DelFlag, userInfo.DepartmentName, userInfo.EmployeeID, userInfo.DomainAccount, userInfo.Email, userInfo.Telephone, userInfo.Mobile, userInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (ur *UserRepo) Delete(id string) error {
	sql := `update t_meta_user_info set del_flag = 1 where id = ?;`
	log.Debugf("metadata UserRepo.Delete() update sql: %s", sql)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = ur.Execute(sql, idInt)

	return err
}
