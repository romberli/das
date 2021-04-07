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
	clusterNameStruct         = "ClusterName"
	middlewareClusterIDStruct = "MiddlewareClusterID"
	monitorSystemIDStruct     = "MonitorSystemID"
	ownerIDStruct             = "OwnerID"
	envIDStruct               = "EnvID"
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
		resp.ResponseNOK(c, message.ErrMetadataGetMySQLClusterAll, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, err.Error())
		return
	}
	fmt.Println("ok")
	// response

	jsonStr := string(jsonBytes)
	fmt.Println(message.DebugMetadataGetMySQLClusterAll, jsonStr)
	log.Debug(message.NewMessage(message.DebugMetadataGetMySQLClusterAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMySQLClusterAll)
}

func GetMySQLClusterByEnv(c *gin.Context) {

}

// @Tags mysql cluster
// @Summary get mysql cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"owner_id":1,"owner_group":"2,3","del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","id":1,"monitor_system_id":1,"env_id":1,"last_update_time":"2021-02-23T20:57:24.603009+08:00","cluster_name":"cluster_name_init","middleware_cluster_id":1}]}"
// @Router /api/v1/metadata/mysql-cluster/:id [get]
func GetMySQLClusterByID(c *gin.Context) {
	// get param
	id := c.Param(idJSON)
	if id == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err := s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetMySQLClusterByID, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetMySQLClusterByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMySQLClusterByID, id)
}

func GetMySQLClusterByName(c *gin.Context) {

}

func GetMySQLServerIDList(c *gin.Context) {

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
	if _, ok := fields[clusterNameStruct]; !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, clusterNameStruct)
		return
	}
	if _, ok := fields[middlewareClusterIDStruct]; !ok {
		fields[middlewareClusterIDStruct] = constant.DefaultRandomInt
	}
	if _, ok := fields[monitorSystemIDStruct]; !ok {
		fields[monitorSystemIDStruct] = constant.DefaultRandomInt
	}
	if _, ok := fields[ownerIDStruct]; !ok {
		fields[ownerIDStruct] = constant.DefaultRandomInt
	}
	if _, ok := fields[envIDStruct]; !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, envIDStruct)
		return
	}

	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataAddMySQLCluster, clusterNameStruct, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataAddMySQLCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataAddMySQLCluster, clusterNameStruct)
}

// @Tags mysql cluster
// @Summary update mysql cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":154,"middleware_cluster_id":0,"owner_id":0,"env_id":0,"create_time":"2021-02-24T02:33:50.936279+08:00","cluster_name":"api_test","monitor_system_id":0,"owner_group":"","del_flag":1,"last_update_time":"2021-02-24T02:33:50.936279+08:00"}]}"
// @Router /api/v1/metadata/mysql-cluster/:id [post]
func UpdateMySQLClusterByID(c *gin.Context) {
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MySQLClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, clusterNameExists := fields[clusterNameStruct]
	_, middlewareClusterIDExists := fields[middlewareClusterIDStruct]
	_, monitorSystemIDExists := fields[monitorSystemIDStruct]
	_, ownerIDExists := fields[ownerIDStruct]
	_, envIDExists := fields[envIDStruct]
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
				clusterNameStruct,
				middlewareClusterIDStruct,
				monitorSystemIDStruct,
				ownerIDStruct,
				envIDStruct,
				delFlagStruct))
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataUpdateMySQLCluster, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, err.Error())
		return
	}
	// resp
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataUpdateMySQLCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.DebugMetadataUpdateMySQLCluster, id)
}

func DeleteMySQLClusterByID(c *gin.Context) {

}
