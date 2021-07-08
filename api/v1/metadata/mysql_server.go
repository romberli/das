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
	msIDJSON        = "id"
	msClusterIDJSON = "cluster_id"
	msHostIPJSON    = "host_ip"
	msPortNumJSON   = "port_num"

	msClusterIDStruct      = "ClusterID"
	msServerNameStruct     = "ServerName"
	msHostIPStruct         = "HostIP"
	msPortNumStruct        = "PortNum"
	msDeploymentTypeStruct = "DeploymentType"
	msVersionStruct        = "Version"
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServerAll, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServerAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServerAll)
}

// TODO: Modify Swagger comment
// @Tags mysql server
// @Summary get mysql servers by cluster id
// @Produce  application/json
// @Success
// @Router /api/v1/metadata/mysql-server/cluster-id/:cluster_id [get]
func GetMySQLServerByClusterID(c *gin.Context) {
	// get param
	clusterIDStr := c.Param(msClusterIDJSON)
	if clusterIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, msIDJSON)
		return
	}
	clusterID, err := strconv.Atoi(clusterIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return

	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	err = s.GetByClusterID(clusterID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServerByClusterID, clusterID, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServerByClusterID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServerByClusterID, clusterID)
}

// @Tags mysql server
// @Summary get mysql server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"port_num":3306,"del_flag":0,"version":"1.1.1","create_time":"2021-02-23T23:43:37.236228+08:00","last_update_time":"2021-02-23T23:43:37.236228+08:00","id":1,"cluster_id":1,"host_ip":"host_ip_init","deployment_type":1}]}"
// @Router /api/v1/metadata/mysql-server/get/:id [get]
func GetMySQLServerByID(c *gin.Context) {
	// get param
	idStr := c.Param(msIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, msIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return

	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServerByID, id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServerByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServerByID, id)
}

// TODO: Modify Swagger comment
// @Tags mysql server
// @Summary get mysql servers by host info
// @Produce  application/json
// @Success
// @Router /api/v1/metadata/mysql-server/:id [get]
func GetMySQLServerByHostInfo(c *gin.Context) {
	// get param
	hostIP := c.Query(msHostIPJSON)
	if hostIP == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, msHostIPJSON)
		return
	}
	portNumStr := c.Query(msPortNumJSON)
	if portNumStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, msPortNumJSON)
		return
	}
	portNum, err := strconv.Atoi(portNumStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return

	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	err = s.GetByHostInfo(hostIP, portNum)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServerByHostInfo, hostIP, portNum, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServerByHostInfo, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServerByHostInfo, hostIP, portNum)
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
	_, clusterIDExists := fields[msClusterIDStruct]
	_, serverNameExists := fields[msServerNameStruct]
	_, hostIPExists := fields[msHostIPStruct]
	_, portNumExists := fields[msPortNumStruct]
	_, deploymentTypeExists := fields[msDeploymentTypeStruct]
	if !clusterIDExists || !serverNameExists || !hostIPExists || !portNumExists || !deploymentTypeExists {
		fmt.Println("wtf")
		resp.ResponseNOK(
			c, message.ErrFieldNotExists,
			fmt.Sprintf(
				"%s and %s and %s and %s and %s",
				msClusterIDStruct,
				msServerNameStruct,
				msHostIPStruct,
				msPortNumStruct,
				msDeploymentTypeStruct))
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddMySQLServer,
			fields[msServerNameStruct],
			fields[msClusterIDStruct],
			fields[msHostIPStruct],
			fields[msPortNumStruct],
			fields[msDeploymentTypeStruct],
			err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddMySQLServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddMySQLServer,
		fields[msServerNameStruct],
		fields[msClusterIDStruct],
		fields[msHostIPStruct],
		fields[msPortNumStruct],
		fields[msDeploymentTypeStruct],
	)
}

// @Tags mysql server
// @Summary update mysql server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"last_update_time":"2021-02-24T02:47:19.589172+08:00","id":93,"cluster_id":0,"host_ip":"192.168.1.1","version":"","del_flag":1,"create_time":"2021-02-24T02:47:19.589172+08:00","port_num":3306,"deployment_type":0}]}"
// @Router /api/v1/metadata/mysql-server/:id [post]
func UpdateMySQLServerByID(c *gin.Context) {
	var fields map[string]interface{}
	// get param
	idStr := c.Param(msIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, msIDJSON)
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MySQLServerInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, clusterIDExists := fields[msClusterIDStruct]
	_, serverNameExists := fields[msServerNameStruct]
	_, hostIPExists := fields[msHostIPStruct]
	_, portNumExists := fields[msPortNumStruct]
	_, deploymentTypeExists := fields[msDeploymentTypeStruct]
	_, versionExists := fields[msVersionStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !clusterIDExists &&
		!serverNameExists &&
		!hostIPExists &&
		!portNumExists &&
		!deploymentTypeExists &&
		!versionExists &&
		!delFlagExists {
		resp.ResponseNOK(
			c, message.ErrFieldNotExists,
			fmt.Sprintf("%s, %s, %s, %s, %s, %s and %s",
				fields[msClusterIDStruct],
				fields[msServerNameStruct],
				fields[msHostIPStruct],
				fields[msPortNumStruct],
				fields[msDeploymentTypeStruct],
				fields[msVersionStruct],
				fields[delFlagStruct]))
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateMySQLServer, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// resp
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateMySQLServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.DebugMetadataUpdateMySQLServer, fields[msServerNameStruct])
}

// TODO: Modify Swagger comment
// @Tags mysql server
// @Summary get mysql servers by host info
// @Produce  application/json
// @Success
// @Router /api/v1/metadata/mysql-server/:id [get]
func DeleteMySQLServerByID(c *gin.Context) {
	var fields map[string]interface{}

	// get param
	idStr := c.Param(msIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, msIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return

	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// insert into middleware
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteMySQLServer,
			id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteMySQLServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteMySQLServer, fields[msServerNameStruct])
}
