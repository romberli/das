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
	idJSON        = "id"
	delFlagStruct = "DelFlag"
	envNameStruct = "EnvName"
)

// @Tags	environment
// @Summary	get all environments
// @Accept	application/json
// @Produce application/json
// @Success	200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router	/api/v1/metadata/env [get]
func GetEnv(c *gin.Context) {
	// init service
	s := metadata.NewEnvServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetEnvAll, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetEnvAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetEnvAll)
}

// @Tags	environment
// @Summary get environment by id
// @Accept	application/json
// @Produce application/json
// @Param	id path int true "environment id"
// @Success	200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router	/api/v1/metadata/env/:id [get]
func GetEnvByID(c *gin.Context) {
	// get param
	id := c.Param(idJSON)
	if id == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
		return
	}
	// init service
	s := metadata.NewEnvServiceWithDefault()
	// get entity
	err := s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetEnvByID, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetEnvByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetEnvByID, id)
}

// @Tags	environment
// @Summary add a new environment
// @Accept	application/json
// @Produce application/json
// @Param	env_name body string true "environment name"
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router	/api/v1/metadata/env [post]
func AddEnv(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.EnvInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, ok := fields[envNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, envNameStruct)
		return
	}
	// init service
	s := metadata.NewEnvServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataAddEnv, envNameStruct, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, envNameStruct, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataAddEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataAddEnv, envNameStruct)
}

// @Tags	environment
// @Summary	update environment by id
// @Accept	application/json
// @Produce application/json
// @Param	id path int true "environment id"
// @Success	200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router	/api/v1/metadata/env/:id [post]
func UpdateEnvByID(c *gin.Context) {
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.EnvInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, envNameExists := fields[envNameStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !envNameExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s", envNameStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewEnvServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataUpdateEnv, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataUpdateEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.DebugMetadataUpdateEnv, id)
}
