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
	middlewareServerClusterIDStruct      = "ClusterID"
	middlewareServerNameStruct           = "ServerName"
	middlewareServerMiddlewareRoleStruct = "MiddlewareRole"
	middlewareServerHostIPStruct         = "HostIP"
	middlewareServerPortNumStruct        = "PortNum"
)

// @Tags middleware server
// @Summary get all middleware servers
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"server_name":"online","del_flag":0,"create_time":"2021-02-20T10:13:41.307685+08:00","port_num":1,"last_update_time":"2021-02-20T10:13:41.307685+08:00","id":1,"cluster_id_middleware":1,"middleware_role":1,"host_ip":"xxxx"},{"last_update_time":"2021-02-22T17:35:09.280743+08:00","id":25,"cluster_id_middleware":2,"middleware_role":2,"host_ip":"123.1.1.1","server_name":"ccc","port_num":2,"del_flag":0,"create_time":"2021-02-22T15:24:40.537777+08:00"},{"id":26,"middleware_role":1,"host_ip":"123.123.123.1","port_num":12,"cluster_id_middleware":3,"server_name":"connection_test","del_flag":0,"create_time":"2021-02-22T17:33:19.202738+08:00","last_update_time":"2021-02-22T17:33:19.202738+08:00"}]}"
// @Router /api/v1/metadata/middleware-server [get]
func GetMiddlewareServer(c *gin.Context) {
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetMiddlewareServerAll, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetMiddlewareServerAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMiddlewareServerAll)
}

// @Tags middleware server
// @Summary get middleware server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":1,"port_num":1,"del_flag":0,"create_time":"2021-02-20T10:13:41.307685+08:00","last_update_time":"2021-02-20T10:13:41.307685+08:00","cluster_id_middleware":1,"server_name":"online","middleware_role":1,"host_ip":"xxxx"}]}"
// @Router /api/v1/metadata/middleware-server/:id [get]
func GetMiddlewareServerByID(c *gin.Context) {
	// get param
	id := c.Param(idJSON)
	if id == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// get entity
	err := s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetMiddlewareServerByID, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetMiddlewareServerByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMiddlewareServerByID, id)
}

// @Tags middleware server
// @Summary add a new middleware server
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":30,"cluster_id_middleware":3,"middleware_role":1,"port_num":12,"last_update_time":"2021-02-23T09:25:38.656337+08:00","server_name":"connection_test","host_ip":"123.123.123.1","del_flag":0,"create_time":"2021-02-23T09:25:38.656337+08:00"}]}"
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
		resp.ResponseNOK(c, message.ErrMetadataAddMiddlewareServer, fields[middlewareServerNameStruct], err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataAddMiddlewareServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataAddMiddlewareServer, fields[middlewareServerNameStruct])
}

// @Tags middleware server
// @Summary update middleware server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"del_flag":1,"server_name":"ccc","id":30,"cluster_id_middleware":3,"middleware_role":1,"host_ip":"123.123.123.1","port_num":12,"create_time":"2021-02-23T09:25:38.656337+08:00","last_update_time":"2021-02-23T09:25:38.656337+08:00"}]}"
// @Router /api/v1/metadata/middleware-server/:id [post]
func UpdateMiddlewareServerByID(c *gin.Context) {
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
		resp.ResponseNOK(c, message.ErrMetadataUpdateMiddlewareServer, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataUpdateMiddlewareServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataUpdateMiddlewareServer, id)
}
