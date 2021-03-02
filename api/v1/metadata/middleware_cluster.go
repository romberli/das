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
	middlewareClusterNameStruct = "ClusterName"
)

// @Tags middleware cluster
// @Summary get all middleware clusters
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":1,"cluster_name":"online","env_id":1,"del_flag":0,"create_time":"2021-02-20T09:25:48.192194+08:00","last_update_time":"2021-02-20T09:25:48.192194+08:00"},{"cluster_name":"newTest","env_id":8,"del_flag":0,"create_time":"2021-02-20T09:18:27.459081+08:00","last_update_time":"2021-02-20T09:18:27.4636+08:00","id":12},{"last_update_time":"2021-02-22T17:09:10.588805+08:00","id":30,"cluster_name":"connection_test","env_id":1,"del_flag":0,"create_time":"2021-02-22T17:09:10.588805+08:00"}]}"
// @Router /api/v1/metadata/middleware-cluster [get]
func GetMiddlewareCluster(c *gin.Context) {
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetMiddlewareClusterAll, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetMiddlewareClusterAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMiddlewareClusterAll)
}

// @Tags middleware cluster
// @Summary get middleware cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"create_time":"2021-02-20T09:25:48.192194+08:00","last_update_time":"2021-02-20T09:25:48.192194+08:00","id":1,"cluster_name":"online","env_id":1,"del_flag":0}]}"
// @Router /api/v1/metadata/middleware-cluster/:id [get]
func GetMiddlewareClusterByID(c *gin.Context) {
	// get param
	id := c.Param(idJSON)
	if id == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err := s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetMiddlewareClusterByID, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetMiddlewareClusterByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMiddlewareClusterByID, id)
}

// @Tags middleware cluster
// @Summary add a new middleware cluster
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":31,"cluster_name":"connection_test","env_id":1,"del_flag":0,"create_time":"2021-02-23T09:21:40.017914+08:00","last_update_time":"2021-02-23T09:21:40.017914+08:00"}]}"
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
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataAddMiddlewareCluster, middlewareClusterNameStruct, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, middlewareClusterNameStruct, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataAddMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataAddMiddlewareCluster, middlewareClusterNameStruct)
}

// @Tags middleware cluster
// @Summary update middleware cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"cluster_name":"ccc","env_id":1,"del_flag":1,"create_time":"2021-02-22T15:23:21.279984+08:00","last_update_time":"2021-02-23T09:22:44.068998+08:00","id":29}]}"
// @Router /api/v1/metadata/middleware-cluster/:id [post]
func UpdateMiddlewareClusterByID(c *gin.Context) {
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
	// unmarshal data=
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MiddlewareClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, middlewareClusterNameExists := fields[middlewareClusterNameStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !middlewareClusterNameExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s", middlewareClusterNameStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataUpdateMiddlewareCluster, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataUpdateMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataUpdateMiddlewareCluster, id)
}
