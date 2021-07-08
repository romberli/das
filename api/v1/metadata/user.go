package metadata

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	msgmeta "github.com/romberli/das/pkg/message/metadata"
	"github.com/romberli/das/pkg/resp"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
)

const (
	userNameJSON    = "user_name"
	employeeIDJSON  = "employee_id"
	accountNameJSON = "account_name"
	emailJSON       = "email"
	telephoneJSON   = "telephone"
	mobileJSON      = "mobile"

	userNameStruct       = "UserName"
	departmentNameStruct = "DepartmentName"
	employeeIDStruct     = "EmployeeID"
	accountNameStruct    = "AccountName"
	emailStruct          = "Email"
	telephoneStruct      = "Telephone"
	roleStruct           = "Role"
	mobileStruct         = "Mobile"
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUserAll, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// responseF
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetUserAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetUserAll)
}

// @Tags user
// @Summary get user by Name
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/user/user-name/:user_name [get]
func GetUserByName(c *gin.Context) {
	// get param
	userName := c.Param(userNameJSON)
	if userName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, userNameJSON)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByName(userName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUserByName, userName, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetUserByName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetUserByName, userName)
}

// @Tags user
// @Summary get user by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/user/get/:id [get]
func GetUserByID(c *gin.Context) {
	// get param
	idStr := c.Param(idJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUserByID, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetUserByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetUserByID, id)
}

// @Tags user
// @Summary get user by EmployeeID
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/user/employee-id/:employee_id [get]
func GetUserByEmployeeID(c *gin.Context) {
	// get param
	employeeID := c.Param(employeeIDJSON)
	if employeeID == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, employeeIDStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByEmployeeID(employeeID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetEmployeeID, employeeID, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetEmployeeID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetEmployeeID, employeeID)
}

// @Tags user
// @Summary get user by AccountName
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/user/account-name/:account_name [get]
func GetUserByAccountName(c *gin.Context) {
	// get param
	accountName := c.Param(accountNameJSON)
	if accountName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, accountNameStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByAccountName(accountName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAccountName, accountName, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAccountName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAccountName, accountName)
}

// @Tags user
// @Summary get user by Email
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata"/user/email/:email [get]
func GetUserByEmail(c *gin.Context) {
	// get param
	email := c.Param(emailJSON)
	if email == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, emailStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByEmail(email)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetEmail, email, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetEmail, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetEmail, email)
}

// @Tags user
// @Summary get user by Telephone
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/user/telephone/:telephone [get]
func GetUserByTelephone(c *gin.Context) {
	// get param
	telephone := c.Param(telephoneJSON)
	if telephone == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, telephoneStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByTelephone(telephone)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetTelephone, telephone, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetTelephone, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetTelephone, telephone)
}

// @Tags user
// @Summary get user by Mobile
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/user/mobile/:mobile [get]
func GetUserByMobile(c *gin.Context) {
	// get param
	mobile := c.Param(mobileJSON)
	if mobile == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mobileStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByMobile(mobile)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMobile, mobile, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMobile, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMobile, mobile)
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
	_, ok = fields[emailStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, emailStruct)
		return
	}
	_, ok = fields[departmentNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, departmentNameStruct)
		return
	}
	// _, ok = fields[employeeIDStruct]
	// if !ok {
	// 	resp.ResponseNOK(c, message.ErrFieldNotExists, employeeIDStruct)
	// 	return
	// }
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddUser, fields[userNameStruct], err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddUser, fields[userNameStruct])
}

// @Tags user
// @Summary update user by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/user/update/:id [post]
func UpdateUserByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(idJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
	}
	id, err := strconv.Atoi(idStr)
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
	_, mobileExists := fields[mobileStruct]
	_, telephoneExists := fields[telephoneStruct]
	_, roleExists := fields[roleStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !userNameExists && !departmentNameExists && !employeeIDExists && !accountNameExists && !emailExists && !telephoneExists && !roleExists && !delFlagExists && !mobileExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s", userNameStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// update UserRepo
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateUser, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, id, err.Error())
		return
	}
	// resp
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.DebugMetadataUpdateUser, id)
}

// @Tags user
// @Summary delete user by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/metadata/user/delete/:id [get]
func DeleteUserByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(idJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// update entities
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteUserByID, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteUserByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteUserByID, fields[userNameStruct])
}
