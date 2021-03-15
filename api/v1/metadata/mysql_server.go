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
	clusterIDStruct      = "ClusterID"
	hostIPStruct         = "HostIP"
	portNumStruct        = "PortNum"
	deploymentTypeStruct = "DeploymentType"
	versionStruct        = "Version"
)

// @Tags mysql server
// @Summary get all mysql servers
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"cluster_id":1,"deployment_type":1,"host_ip":"host_ip_init","port_num":3306,"version":"1.1.1","del_flag":0,"create_time":"2021-02-23T23:43:37.236228+08:00","last_update_time":"2021-02-23T23:43:37.236228+08:00","id":1}]}"
// @Router /api/v1/metadata/mysql-server [get]
func GetMySQLServer(c *gin.Context) {
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetMySQLServerAll, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetMySQLServerAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMySQLServerAll)
}

// @Tags mysql server
// @Summary get mysql server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"port_num":3306,"del_flag":0,"version":"1.1.1","create_time":"2021-02-23T23:43:37.236228+08:00","last_update_time":"2021-02-23T23:43:37.236228+08:00","id":1,"cluster_id":1,"host_ip":"host_ip_init","deployment_type":1}]}"
// @Router /api/v1/metadata/mysql-server/:id [get]
func GetMySQLServerByID(c *gin.Context) {
	// get param
	id := c.Param(idJSON)
	if id == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	err := s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetMySQLServerByID, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetMySQLServerByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMySQLServerByID, id)
}

// @Tags mysql server
// @Summary add a new mysql server
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"create_time":"2021-02-24T02:47:19.589172+08:00","del_flag":0,"last_update_time":"2021-02-24T02:47:19.589172+08:00","id":93,"cluster_id":0,"host_ip":"192.168.1.1","port_num":3306,"deployment_type":0,"version":""}]}"
// @Router /api/v1/metadata/mysql-server [post]
func AddMySQLServer(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MySQLServerInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	if _, ok := fields[hostIPStruct]; !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hostIPStruct)
		return
	}
	if _, ok := fields[portNumStruct]; !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, portNumStruct)
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataAddMySQLServer, hostIPStruct, portNumStruct, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, hostIPStruct, portNumStruct, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataAddMySQLServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataAddMySQLServer, hostIPStruct, portNumStruct)
}

// @Tags mysql server
// @Summary update mysql server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"last_update_time":"2021-02-24T02:47:19.589172+08:00","id":93,"cluster_id":0,"host_ip":"192.168.1.1","version":"","del_flag":1,"create_time":"2021-02-24T02:47:19.589172+08:00","port_num":3306,"deployment_type":0}]}"
// @Router /api/v1/metadata/mysql-server/:id [post]
func UpdateMySQLServerByID(c *gin.Context) {
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MySQLServerInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, hostIPExists := fields[hostIPStruct]
	_, portNumExists := fields[portNumStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !hostIPExists && !portNumExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s:%s and %s", hostIPStruct, portNumStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataUpdateMySQLServer, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataUpdateMySQLServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.DebugMetadataUpdateMySQLServer, id)
}
