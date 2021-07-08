package metadata

import (
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
	middlewareServerIDJSON        = "id"
	middlewareServerClusterIDJSON = "cluster_id"
	middlewareServerHostIPJSON    = "host_ip"
	middlewareServerPortNumJSON   = "port_num"

	middlewareServerClusterIDStruct      = "ClusterID"
	middlewareServerNameStruct           = "ServerName"
	middlewareServerMiddlewareRoleStruct = "MiddlewareRole"
	middlewareServerHostIPStruct         = "HostIP"
	middlewareServerPortNumStruct        = "PortNum"
)

// @Tags middleware server
// @Summary get all middleware servers
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"port_num":1,"last_update_time":"2021-04-11T23:16:10.281222+08:00","server_name":"test001","middleware_role":1,"host_ip":"3","del_flag":0,"create_time":"2021-04-07T17:51:00.270268+08:00","id":1,"cluster_id":13},{"last_update_time":"2021-04-09T16:20:03.063295+08:00","id":2,"cluster_id":13,"server_name":"test002","del_flag":0,"create_time":"2021-04-09T16:20:03.063295+08:00","middleware_role":2,"host_ip":"2","port_num":2}]}"
// @Router /api/v1/metadata/middleware-server [get]
func GetMiddlewareServer(c *gin.Context) {
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareServerAll, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareServerAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareServerAll)
}

// @Tags middleware server
// @Summary get middleware servers by cluster id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"server_name":"test001","middleware_role":1,"host_ip":"3","del_flag":0,"create_time":"2021-04-07T17:51:00.270268+08:00","last_update_time":"2021-04-11T23:16:10.281222+08:00","id":1,"cluster_id":13,"port_num":1},{"host_ip":"2","port_num":2,"del_flag":0,"last_update_time":"2021-04-09T16:20:03.063295+08:00","id":2,"cluster_id":13,"server_name":"test002","middleware_role":2,"create_time":"2021-04-09T16:20:03.063295+08:00"}]}"
// @Router /api/v1/metadata/middleware-server/cluster-id/:cluster_id [get]
func GetMiddlewareServerByClusterID(c *gin.Context) {
	// get param
	clusterIDStr := c.Param(middlewareServerClusterIDJSON)
	if clusterIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerClusterIDJSON)
		return
	}
	clusterID, err := strconv.Atoi(clusterIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// get entity
	err = s.GetByClusterID(clusterID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareSeverByClusterID, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareSeverByClusterID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareSeverByClusterID)
}

// @Tags middleware server
// @Summary get middleware server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":1,"middleware_role":1,"del_flag":0,"last_update_time":"2021-04-11T23:16:10.281222+08:00","cluster_id":13,"server_name":"test001","host_ip":"3","port_num":1,"create_time":"2021-04-07T17:51:00.270268+08:00"}]}"
// @Router /api/v1/metadata/middleware-server/get/:id [get]
func GetMiddlewareServerByID(c *gin.Context) {
	// get param
	idStr := c.Param(middlewareServerIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareServerByID, id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareServerByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareServerByID, id)
}

// @Tags middleware server
// @Summary get middleware server by host info
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"cluster_id":13,"server_name":"test001","port_num":1,"last_update_time":"2021-04-11T23:16:10.281222+08:00","id":1,"middleware_role":1,"host_ip":"3","del_flag":0,"create_time":"2021-04-07T17:51:00.270268+08:00"}]}"
// @Router /api/v1/metadata/middleware-server/host-info [get]
func GetMiddlewareServerByHostInfo(c *gin.Context) {
	// get params
	middleServerHostIP := c.Query(middlewareServerHostIPJSON)
	if middleServerHostIP == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerHostIPJSON)
		return
	}
	middleServerPortNumStr := c.Query(middlewareServerPortNumJSON)
	if middleServerPortNumStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerPortNumJSON)
		return
	}
	middleServerPortNum, err := strconv.Atoi(middleServerPortNumStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	hostIP := middleServerHostIP + "-" + middleServerPortNumStr
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// get entity
	err = s.GetByHostInfo(middleServerHostIP, middleServerPortNum)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareServerByHostInfo, hostIP)
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareServerByHostInfo, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareServerByHostInfo, hostIP)
}

// @Tags middleware server
// @Summary add a new middleware server
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":31,"middleware_role":1,"port_num":12,"create_time":"2021-04-12T10:59:11.559227+08:00","last_update_time":"2021-04-12T10:59:11.559227+08:00","cluster_id":13,"server_name":"test003","host_ip":"123.123.123.1","del_flag":0}]}"
// @Router /api/v1/metadata/middleware-server [post]
func AddMiddlewareServer(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MiddlewareServerInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, ok := fields[middlewareServerClusterIDStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerClusterIDStruct)
		return
	}
	_, ok = fields[middlewareServerNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerNameStruct)
		return
	}
	_, ok = fields[middlewareServerMiddlewareRoleStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerMiddlewareRoleStruct)
		return
	}
	_, ok = fields[middlewareServerHostIPStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerHostIPStruct)
		return
	}
	_, ok = fields[middlewareServerPortNumStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerPortNumStruct)
		return
	}
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddMiddlewareServer, fields[middlewareServerNameStruct], err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddMiddlewareServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddMiddlewareServer, fields[middlewareServerNameStruct])
}

// @Tags middleware server
// @Summary update middleware server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"server_name":"newTest","host_ip":"123.123.123.1","last_update_time":"2021-04-12T10:59:11.559227+08:00","id":31,"cluster_id":13,"port_num":12,"del_flag":1,"create_time":"2021-04-12T10:59:11.559227+08:00","middleware_role":1}]}"
// @Router /api/v1/metadata/middleware-server/update/:id [post]
func UpdateMiddlewareServerByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(middlewareServerIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerIDJSON)
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MiddlewareServerInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, middlewareServerClusterIDExists := fields[middlewareServerClusterIDStruct]
	_, middlewareServerNameExists := fields[middlewareServerNameStruct]
	_, middlewareServerMiddlewareRoleExists := fields[middlewareServerMiddlewareRoleStruct]
	_, middlewareServerHostIPExists := fields[middlewareServerHostIPStruct]
	_, middlewareServerPortNumExists := fields[middlewareServerPortNumStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !middlewareServerClusterIDExists && !middlewareServerNameExists && !middlewareServerMiddlewareRoleExists && !middlewareServerHostIPExists && !middlewareServerPortNumExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s", middlewareServerNameStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateMiddlewareServer, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateMiddlewareServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataUpdateMiddlewareServer, id)
}

// @Tags middleware server
// @Summary delete middleware server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/metadata/middleware-server/delete/:id [post]
func DeleteMiddlewareServerByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(middlewareServerIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// update entities
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteMiddlewareServer, fields[middlewareClusterNameStruct], err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteMiddlewareServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteMiddlewareServer, fields[middlewareClusterNameStruct])
}
