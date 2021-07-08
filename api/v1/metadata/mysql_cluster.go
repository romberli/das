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
	mcIDJSON          = "id"
	mcEnvIDJSON       = "env_id"
	mcClusterNameJSON = "name"

	mcClusterNameStruct         = "ClusterName"
	mcMiddlewareClusterIDStruct = "MiddlewareClusterID"
	mcMonitorSystemIDStruct     = "MonitorSystemID"
	mcOwnerIDStruct             = "OwnerID"
	mcEnvIDStruct               = "EnvID"
	mcMySQLServerIDListStruct   = "MySQLServerIDList"
)

// @Tags mysql cluster
// @Summary get all mysql clusters
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"middleware_cluster_id":1,"monitor_system_id":1,"env_id":1,"owner_group":"2,3","del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","last_update_time":"2021-02-23T20:57:24.603009+08:00","id":1,"cluster_name":"cluster_name_init","owner_id":1},{"monitor_system_id":1,"owner_id":1,"owner_group":"2,3","env_id":1,"create_time":"2021-02-23T04:14:23.707238+08:00","last_update_time":"2021-02-23T04:14:23.707238+08:00","id":2,"cluster_name":"newTest","middleware_cluster_id":1,"del_flag":0}]}"
// @Router /api/v1/metadata/mysql-cluster [get]
func GetMySQLCluster(c *gin.Context) {
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterAll, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	fmt.Println("ok")
	// response

	jsonStr := string(jsonBytes)
	fmt.Println(msgmeta.DebugMetadataGetMySQLClusterAll, jsonStr)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClusterAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClusterAll)
}

func GetMySQLClusterByEnv(c *gin.Context) {
	// get param
	envIDStr := c.Param(mcEnvIDJSON)
	if envIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mcEnvIDJSON)
		return
	}
	envID, err := strconv.Atoi(envIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return

	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetByEnv(envID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterByEnv, envID, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClusterByEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClusterByEnv, envID)
}

// @Tags mysql cluster
// @Summary get mysql cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"owner_id":1,"owner_group":"2,3","del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","id":1,"monitor_system_id":1,"env_id":1,"last_update_time":"2021-02-23T20:57:24.603009+08:00","cluster_name":"cluster_name_init","middleware_cluster_id":1}]}"
// @Router /api/v1/metadata/mysql-cluster/:id [get]
func GetMySQLClusterByID(c *gin.Context) {
	// get param
	idStr := c.Param(mcIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mcIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return

	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterByID, id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClusterByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClusterByID, id)
}

func GetMySQLClusterByName(c *gin.Context) {
	// get param
	clusterName := c.Param(mcClusterNameJSON)
	if clusterName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mcClusterNameJSON)
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err := s.GetByName(clusterName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterByName, clusterName, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClusterByName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClusterByName, clusterName)
}

func GetMySQLServerIDList(c *gin.Context) {
	// get params
	idStr := c.Param(mcIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mcIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetMySQLServerIDList(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServerIDList, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(mcMySQLServerIDListStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(jsonStr)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServerIDList, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServerIDList, id)
}

// @Tags mysql cluster
// @Summary add a new mysql cluster
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"cluster_name":"api_test","monitor_system_id":0,"owner_group":"","del_flag":0,"create_time":"2021-02-24T02:33:50.936279+08:00","last_update_time":"2021-02-24T02:33:50.936279+08:00","middleware_cluster_id":0,"owner_id":0,"env_id":0,"id":154}]}"
// @Router /api/v1/metadata/mysql-cluster [post]
func AddMySQLCluster(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MySQLClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	if _, ok := fields[mcClusterNameStruct]; !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mcClusterNameStruct)
		return
	}
	if _, ok := fields[mcEnvIDStruct]; !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mcEnvIDStruct)
		return
	}

	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddMySQLCluster,
			fields[mcClusterNameStruct],
			fields[mcEnvIDStruct],
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddMySQLCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddMySQLCluster,
		fields[mcClusterNameStruct],
		fields[mcEnvIDStruct],
	)
}

// @Tags mysql cluster
// @Summary update mysql cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":154,"middleware_cluster_id":0,"owner_id":0,"env_id":0,"create_time":"2021-02-24T02:33:50.936279+08:00","cluster_name":"api_test","monitor_system_id":0,"owner_group":"","del_flag":1,"last_update_time":"2021-02-24T02:33:50.936279+08:00"}]}"
// @Router /api/v1/metadata/mysql-cluster/:id [post]
func UpdateMySQLClusterByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(mcIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mcIDJSON)
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MySQLClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, clusterNameExists := fields[mcClusterNameStruct]
	_, middlewareClusterIDExists := fields[mcMiddlewareClusterIDStruct]
	_, monitorSystemIDExists := fields[mcMonitorSystemIDStruct]
	_, ownerIDExists := fields[mcOwnerIDStruct]
	_, envIDExists := fields[mcEnvIDStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !clusterNameExists &&
		!middlewareClusterIDExists &&
		!monitorSystemIDExists &&
		!ownerIDExists &&
		!envIDExists &&
		!delFlagExists {
		resp.ResponseNOK(
			c, message.ErrFieldNotExists,
			fmt.Sprintf("%s, %s, %s, %s, %s and %s",
				mcClusterNameStruct,
				mcMiddlewareClusterIDStruct,
				mcMonitorSystemIDStruct,
				mcOwnerIDStruct,
				mcEnvIDStruct,
				delFlagStruct))
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateMySQLCluster, id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateMySQLCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.DebugMetadataUpdateMySQLCluster, fields[mcClusterNameStruct])
}

func DeleteMySQLClusterByID(c *gin.Context) {
	var fields map[string]interface{}

	// get param
	idStr := c.Param(mcIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mcIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return

	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// insert into middleware
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteMySQLCluster,
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteMySQLCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteMySQLCluster, fields[mcClusterNameStruct])
}
