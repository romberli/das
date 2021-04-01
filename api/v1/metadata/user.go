package metadata

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/das/pkg/resp"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
)

const (
	userNameStruct       = "UserName"
	departmentNameStruct = "DepartmentName"
	employeeIDStruct     = "EmployeeID"
	accountNameStruct    = "AccountName"
	emailStruct          = "Email"
	telephoneStruct      = "Telephone"
	roleStruct           = "Role"
)

// @Tags user
// @Summary get all users
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/user [get]

func GetUser(c *gin.Context) {
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetUserAll, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataGetUserAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetUserAll)
}

func GetUserByName(c *gin.Context) {

}

// @Tags user
// @Summary get user by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/user/:id [get]
func GetUserByID(c *gin.Context) {
	// get param
	id := c.Param(idJSON)
	if id == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get entity
	err := s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetUserByID, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataGetUserByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetUserByID, id)
}

func GetUserByEmployeeID(c *gin.Context) {

}

func GetUserByAccountName(c *gin.Context) {

}

func GetUserByEmail(c *gin.Context) {

}

func GetUserByTelephone(c *gin.Context) {

}

func GetUserByMobile(c *gin.Context) {

}

// @Tags user
// @Summary add a new user
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/user [post]
func AddUser(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.UserInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, ok := fields[userNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, userNameStruct)
		return
	}
	_, ok = fields[departmentNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, departmentNameStruct)
		return
	}
	_, ok = fields[employeeIDStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, employeeIDStruct)
		return
	}
	_, ok = fields[accountNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, accountNameStruct)
		return
	}
	_, ok = fields[roleStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, roleStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataAddUser, fields[userNameStruct], err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataAddUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataAddUser, fields[userNameStruct])
}

// @Tags user
// @Summary update user by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/user/:id [post]
func UpdateUserByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	id := c.Param(idJSON)
	if id == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.UserInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, userNameExists := fields[userNameStruct]
	_, departmentNameExists := fields[departmentNameStruct]
	_, employeeIDExists := fields[employeeIDStruct]
	_, accountNameExists := fields[accountNameStruct]
	_, emailExists := fields[emailStruct]
	_, telephoneExists := fields[telephoneStruct]
	_, roleExists := fields[roleStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !userNameExists && !departmentNameExists && !employeeIDExists && !accountNameExists && !emailExists && !telephoneExists && !roleExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s", userNameStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataUpdateUser, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, id, err.Error())
		return
	}
	// resp
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataUpdateUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.DebugMetadataUpdateUser, id)
}

func DeleteUserByID(c *gin.Context) {

}
