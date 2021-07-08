package metadata

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"

	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	msgmeta "github.com/romberli/das/pkg/message/metadata"
	"github.com/romberli/das/pkg/resp"
)

const (
	appIDJSON      = "id"
	appAppNameJSON = "app_name"
	appDBIDJSON    = "db_id"

	appAppNameStruct  = "AppName"
	appLevelStruct    = "Level"
	appDelFlagStruct  = "DelFlag"
	appDBIDListStruct = "DBIDList"
)

// @Tags application
// @Summary get all applications
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8,"owner_group": "k"}]}"
// @Router /api/v1/metadata/app [get]
func GetApp(c *gin.Context) {
	// init service
	s := metadata.NewAppServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppAll, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppAll)
}

// @Tags application
// @Summary get application by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8,"owner_group": "k"}]}"
// @Router /api/v1/metadata/app/:id [get]
func GetAppByID(c *gin.Context) {
	// get param
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppByID, id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppByID, id)
}

// @Tags application
// @Summary get application by system name
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "app_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8,"owner_group": "k"}]}"
// @Router /api/v1/metadata/app/app-name/:name [get]
func GetAppByName(c *gin.Context) {
	// get params
	appName := c.Param(appAppNameJSON)
	if appName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appAppNameJSON)
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// get entity
	err := s.GetAppByName(appName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppByName, appName, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppByName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppByName, appName)
}

// @Tags application
// @Summary get db id list
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1, 2]}"
// @Router /api/vi/metadata/app/dbs/:id [get]
func GetDBIDList(c *gin.Context) {
	// get params
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// get entity
	err = s.GetDBIDList(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBIDList, id, err.Error())
		return
	}

	b := s.DBIDList
	fmt.Println(b)
	// marshal service
	jsonBytes, err := s.MarshalWithFields(appDBIDListStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBIDList, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBIDList, id)
}

// @Tags application
// @Summary add a new application
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8,"owner_group": "k"}]}"
// @Router /api/v1/metadata/app [post]
func AddApp(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.AppInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, ok := fields[appAppNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appAppNameStruct)
		return
	}
	_, ok = fields[appLevelStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appLevelStruct)
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddApp, fields[appAppNameStruct], err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddApp, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddApp, fields[appAppNameStruct])
}

// @Tags application
// @Summary update application by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8,"owner_group": "k"}]}"
// @Router /api/v1/metadata/app/:id [post]
func UpdateAppByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.AppInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, AppNameExists := fields[appAppNameStruct]
	_, levelExists := fields[appLevelStruct]
	_, delFlagExists := fields[appDelFlagStruct]
	if !AppNameExists && !delFlagExists && !levelExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s", appAppNameStruct, appDelFlagStruct))
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// update entities
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateApp, id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateApp, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataUpdateApp, fields[appAppNameStruct])
}

// @Tags application
// @Summary delete app by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8,"owner_group": "k"}]}"
// @Router /api/v1/metadata/app/delete/:id [post]
func DeleteAppByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// update entities
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteApp, id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteApp, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteApp, fields[appAppNameStruct])
}

// @Tags application
// @Summary add database map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8,"owner_group": "k"}]}"
// @Router /api/v1/metadata/app/add-db/:id [post]
func AppAddDB(c *gin.Context) {
	// get params
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}

	dataMap := make(map[string]int)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppAddDB, id, err.Error())
		return
	}
	dbID, dbIDExists := dataMap[appDBIDJSON]
	if !dbIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appDBIDJSON)
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// update entities
	err = s.AddDB(id, dbID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppAddDB, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(appDBIDListStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAppAddDB, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAppAddDB, id, dbID)
}

// @Tags application
// @Summary delete database map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1]}"
// @Router /api/v1/metadata/app/delete-db/:id [post]
func AppDeleteDB(c *gin.Context) {
	// get params
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	dataMap := make(map[string]int)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppDeleteDB, id, err.Error())
		return
	}
	dbID, dbIDExists := dataMap[appDBIDJSON]
	if !dbIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appDBIDJSON)
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// update entities
	err = s.DeleteDB(id, dbID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppDeleteDB, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(appDBIDListStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAppDeleteDB, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAppDeleteDB, id, dbID)
}
