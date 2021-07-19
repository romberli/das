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
	msgmeta "github.com/romberli/das/pkg/message/metadata"

	"github.com/romberli/das/pkg/message"
	"github.com/romberli/das/pkg/resp"
)

const (
	dbIDJSON          = "id"
	dbEnvIDJSON       = "env_id"
	dbAppIDJSON       = "app_id"
	dbDBNameJSON      = "db_name"
	dbClusterIDJSON   = "cluster_id"
	dbClusterTypeJSON = "cluster_type"

	dbDBNameStruct      = "DBName"
	dbClusterIDStruct   = "ClusterID"
	dbClusterTypeStruct = "ClusterType"
	dbOwnerIDStruct     = "OwnerID"
	dbEnvIDStruct       = "EnvID"
	dbAppIDListStruct   = "AppIDList"
)

// @Tags database
// @Summary get all databases
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "owner_id": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db [get]
func GetDB(c *gin.Context) {
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBAll, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBAll)
}

// @Tags database
// @Summary get database by env_id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "owner_id": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/env/:env_id [get]
func GetDBByEnv(c *gin.Context) {
	// get param
	envIDStr := c.Param(dbEnvIDJSON)
	if envIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbEnvIDJSON)
		return
	}
	envID, err := strconv.Atoi(envIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return

	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetByEnv(envID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBByEnv, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBByEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBByEnv, envID)
}

// @Tags database
// @Summary get database by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "owner_id": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/get/:id [get]
func GetDBByID(c *gin.Context) {
	// get param
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBByID, id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBByID, id)
}

// @Tags database
// @Summary get database by db name and cluster info
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "owner_id": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/name-and-cluster-info[get]
func GetDBByNameAndClusterInfo(c *gin.Context) {
	var dbInfo *metadata.DBInfo
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	err = json.Unmarshal(data, dbInfo)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetByNameAndClusterInfo(dbInfo.DBName, dbInfo.ClusterID, dbInfo.ClusterType)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBByNameAndClusterInfo, dbInfo.DBName, dbInfo.ClusterID, dbInfo.ClusterType, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBByNameAndClusterInfo, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBByNameAndClusterInfo, dbInfo.DBName, dbInfo.ClusterID, dbInfo.ClusterType, err.Error())
}

// @Tags db
// @Summary get app id list
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1, 2]}"
// @Router /api/v1/metadata/db/apps/:id [get]
func GetAppIDList(c *gin.Context) {
	// get params
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetAppIDList(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppIDList, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbAppIDListStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppIDList, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppIDList, id)
}

// @Tags database
// @Summary add a new database
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "owner_id": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db [post]
func AddDB(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.DBInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, dbNameExists := fields[dbDBNameStruct]
	_, clusterIDExists := fields[dbClusterIDStruct]
	_, clusterTypeExists := fields[dbClusterTypeStruct]
	_, envIDExists := fields[dbEnvIDStruct]
	if !dbNameExists || !clusterIDExists || !clusterTypeExists || !envIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s",
			dbDBNameStruct, dbClusterIDStruct, dbClusterTypeStruct, dbEnvIDStruct))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddDB,
			fields[dbDBNameStruct], fields[dbClusterIDStruct], fields[dbClusterTypeStruct], fields[dbEnvIDStruct], err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddDB, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddDB, fields[dbDBNameStruct], fields[dbClusterIDStruct], fields[dbClusterTypeStruct], fields[dbEnvIDStruct])
}

// @Tags database
// @Summary update database by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "owner_id": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/update/:id [post]
func UpdateDBByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.DBInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, dbNameExists := fields[dbDBNameStruct]
	_, clusterIDExists := fields[dbClusterIDStruct]
	_, clusterTypeExists := fields[dbClusterTypeStruct]
	_, ownerIDExists := fields[dbOwnerIDStruct]
	_, envIDExists := fields[dbEnvIDStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !dbNameExists && !clusterIDExists && !clusterTypeExists && !ownerIDExists && !envIDExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists,
			fmt.Sprintf("%s, %s, %s, %s, %s, %s",
				dbDBNameStruct, dbClusterIDStruct, dbClusterTypeStruct, dbOwnerIDStruct, dbEnvIDStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateDB, id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateDB, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataUpdateDB, id)
}

// @Tags database
// @Summary delete database by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "owner_id": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/delete/:id [post]
func DeleteDBByID(c *gin.Context) {
	// get params
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entity
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteDB, id, err.Error())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteDB, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteDB, id)
}

// @Tags database
// @Summary add application map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1, 2]}"
// @Router /api/v1/metadata/db/add-app/:id [post]
func DBAddApp(c *gin.Context) {
	// get params
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
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
	dataMap := make(map[string]int)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBAddApp, id, err.Error())
		return
	}
	appID, appIDExists := dataMap[dbAppIDJSON]
	if !appIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbAppIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entities
	err = s.AddApp(id, appID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBAddApp, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbAppIDListStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDBAddApp, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDBAddApp, id, appID)
}

// @Tags database
// @Summary delete application map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1]}"
// @Router /api/v1/metadata/db/delete-app/:id [post]
func DBDeleteApp(c *gin.Context) {
	// get params
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
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
	dataMap := make(map[string]int)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBDeleteApp, id, err.Error())
		return
	}
	appID, appIDExists := dataMap[dbAppIDJSON]
	if !appIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbAppIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entities
	err = s.DeleteApp(id, appID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBDeleteApp, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbAppIDListStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDBAddApp, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDBDeleteApp, id, appID)
}
