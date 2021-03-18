package metadata

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"

	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/das/pkg/resp"
)

const (
	// idJSON        = "id"
	// delFlagStruct = "DelFlag"
	appSystemNameStruct = "AppSystemName"
	levelStruct         = "Level"
)

// @Tags application system
// @Summary get all application systems
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8,"owner_group": "k"}]}"
// @Router /api/v1/metadata/app-system [get]
func GetAppSystem(c *gin.Context) {
	// init service
	fmt.Println(ownerGroupStruct)
	s := metadata.NewAppSystemServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetAppSystemAll, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetAppSystemAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetAppSystemAll)
}

// @Tags application system
// @Summary get application system by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8,"owner_group": "k"}]}"
// @Router /api/v1/metadata/app-system/:id [get]
func GetAppSystemByID(c *gin.Context) {
	// get param
	id := c.Param(idJSON)
	if id == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
		return
	}
	// init service
	s := metadata.NewAppSystemServiceWithDefault()
	// get entity
	err := s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetAppSystemByID, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, id, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataGetAppSystemByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetAppSystemByID, id)
}

// @Tags application system
// @Summary add a new application system
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8,"owner_group": "k"}]}"
// @Router /api/v1/metadata/app-system [post]
func AddAppSystem(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.AppSystemInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, ok := fields[appSystemNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appSystemNameStruct)
		return
	}

	_, ok = fields[levelStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, levelStruct)
		return
	}
	// init service
	s := metadata.NewAppSystemServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataAddAppSystem, appSystemNameStruct, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, appSystemNameStruct, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataAddAppSystem, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataAddAppSystem, appSystemNameStruct)
}

// @Tags application system
// @Summary update application system by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8,"owner_group": "k"}]}"
// @Router /api/v1/metadata/app-system/:id [post]
func UpdateAppSystemByID(c *gin.Context) {
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.AppSystemInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, appSystemNameExists := fields[appSystemNameStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !appSystemNameExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s", appSystemNameStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewAppSystemServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataUpdateAppSystem, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataUpdateAppSystem, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.DebugMetadataUpdateAppSystem, id)
}
