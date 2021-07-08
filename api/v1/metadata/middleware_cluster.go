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
	middlewareClusterIDJSON    = "id"
	middlewareClusterNameJSON  = "cluster_name"
	middlewareClusterEnvIDJSON = "env_id"

	middlewareClusterNameStruct    = "ClusterName"
	middlewareClusterOwnerIDStruct = "OwnerID"
	middlewareClusterEnvIDStruct   = "EnvID"
)

// @Tags middleware cluster
// @Summary get all middleware clusters
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":13,"cluster_name":"test001","owner_id":1,"env_id":1,"del_flag":0,"create_time":"2021-04-09T10:55:43.920406+08:00","last_update_time":"2021-04-09T10:55:43.920406+08:00"}]}"
// @Router /api/v1/metadata/middleware-cluster [get]
func GetMiddlewareCluster(c *gin.Context) {
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterAll, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterAll)
}

// @Tags middleware cluster
// @Summary get middleware cluster by env
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"del_flag":0,"create_time":"2021-04-09T10:55:43.920406+08:00","last_update_time":"2021-04-09T10:55:43.920406+08:00","id":13,"cluster_name":"test001","owner_id":1,"env_id":1}]}"
// @Router /api/v1/metadata/middleware-cluster/env/:env_id [get]
func GetMiddlewareClusterByEnv(c *gin.Context) {
	// get param
	envIDStr := c.Param(middlewareClusterEnvIDJSON)
	if envIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
		return
	}
	envID, err := strconv.Atoi(envIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err = s.GetByEnv(envID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterByEnv, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterByEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterByEnv, envID)
}

// @Tags middleware cluster
// @Summary get middleware cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":13,"cluster_name":"test001","owner_id":1,"env_id":1,"del_flag":0,"create_time":"2021-04-09T10:55:43.920406+08:00","last_update_time":"2021-04-09T10:55:43.920406+08:00"}]}"
// @Router /api/v1/metadata/middleware-cluster/get/:id [get]
func GetMiddlewareClusterByID(c *gin.Context) {
	// get param
	idStr := c.Param(middlewareClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterByID, id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterByID, id)
}

// @Tags middleware cluster
// @Summary get middleware cluster by name
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":13,"cluster_name":"test001","owner_id":1,"env_id":1,"del_flag":0,"create_time":"2021-04-09T10:55:43.920406+08:00","last_update_time":"2021-04-09T10:55:43.920406+08:00"}]}"
// @Router /api/v1/metadata/middleware-cluster/cluster-name/:cluster_name [get]
func GetMiddlewareClusterByName(c *gin.Context) {
	// get params
	middlewareClusterName := c.Param(middlewareClusterNameJSON)
	if middlewareClusterName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterNameJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err := s.GetByName(middlewareClusterName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterByName, middlewareClusterName, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterByName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterByName, middlewareClusterName)
}

// @Tags application
// @Summary get middleware server id list by cluster id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1,2]}"
// @Router /api/vi/metadata/middleware-server/:id [get]
func GetMiddlewareServerIDList(c *gin.Context) {
	// get params
	idStr := c.Param(middlewareClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	_, err = s.GetMiddlewareServerIDList(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareServerIDList, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := json.Marshal(s.MiddlewareServerList)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareServerIDList, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareServerIDList, id)

}

// @Tags middleware cluster
// @Summary add a new middleware cluster
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"del_flag":0,"create_time":"2021-04-09T16:02:25.541701+08:00","last_update_time":"2021-04-09T16:02:25.541701+08:00","id":14,"cluster_name":"rest_test","owner_id":1,"env_id":1}]}"
// @Router /api/v1/metadata/middleware-cluster [post]
func AddMiddlewareCluster(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MiddlewareClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, ok := fields[middlewareClusterNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterNameStruct)
		return
	}
	_, ok = fields[middlewareClusterEnvIDStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterEnvIDStruct)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddMiddlewareCluster, fields[middlewareClusterNameStruct], err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddMiddlewareCluster, fields[middlewareClusterNameStruct])
}

// @Tags middleware cluster
// @Summary update middleware cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":13,"cluster_name":"new_test","owner_id":1,"env_id":1,"del_flag":1,"create_time":"2021-04-09T10:55:43.920406+08:00","last_update_time":"2021-04-09T10:55:43.920406+08:00"}]}"
// @Router /api/v1/metadata/middleware-cluster/update/:id [post]
func UpdateMiddlewareClusterByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(middlewareClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MiddlewareClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, middlewareClusterNameExists := fields[middlewareClusterNameStruct]
	_, middlewareClusterOwnerIDExists := fields[middlewareClusterOwnerIDStruct]
	_, middlewareClusterEnvIDExists := fields[middlewareClusterEnvIDStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !middlewareClusterNameExists && !middlewareClusterEnvIDExists && !middlewareClusterOwnerIDExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s %s %s and %s", middlewareClusterNameStruct, middlewareClusterOwnerIDStruct, middlewareClusterEnvIDStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateMiddlewareCluster, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataUpdateMiddlewareCluster, id)
}

// @Tags application
// @Summary delete middleware cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/metadata/app/delete/:id [post]
func DeleteMiddlewareClusterByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(middlewareClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// update entities
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteMiddlewareCluster, fields[middlewareClusterNameStruct], err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteMiddlewareCluster, fields[middlewareClusterNameStruct])
}
