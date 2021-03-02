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
	mSNameStruct      = "MSName"
	systemTypeStruct  = "SystemType"
	mSHostIpStruct    = "HostIp"
	mSportNumStruct   = "PortNum"
	portNumSlowStruct = "PortNumSlow"
	baseUrlStruct     = "BaseUrl"
)

// @Tags monitor system
// @Summary get all monitor systems
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system [get]
func GetMonitorSystem(c *gin.Context) {
	// init service
	s := metadata.NewMSServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetMSAll, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetMSAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMSAll)
}

// @Tags monitor system
// @Summary get monitor system by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system/:id [get]
func GetMonitorSystemByID(c *gin.Context) {
	// get param
	id := c.Param(idJSON)
	if id == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
		return
	}
	// init service
	s := metadata.NewMSServiceWithDefault()
	// get entity
	err := s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetMSByID, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetMSByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMSByID, id)
}

// @Tags monitor system
// @Summary add a new monitor system
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system [post]
func AddMonitorSystem(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MSInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, ok := fields[mSNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mSNameStruct)
		return
	}
	// init service
	s := metadata.NewMSServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataAddMS, fmt.Sprintf("%s and %s and %s and %s and %s and %s", mSNameStruct, systemTypeStruct, mSHostIpStruct, mSportNumStruct, portNumSlowStruct, baseUrlStruct), err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, fmt.Sprintf("%s and %s and %s and %s and %s and %s", mSNameStruct, systemTypeStruct, mSHostIpStruct, mSportNumStruct, portNumSlowStruct, baseUrlStruct), err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataAddMS, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataAddMS, fmt.Sprintf("%s and %s and %s and %s and %s and %s", mSNameStruct, systemTypeStruct, mSHostIpStruct, mSportNumStruct, portNumSlowStruct, baseUrlStruct))
}

// @Tags monitor system
// @Summary update monitor system by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system/:id [post]
func UpdateMonitorSystemByID(c *gin.Context) {
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MSInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, mSNameExists := fields[mSNameStruct]
	_, systemTypeExists := fields[systemTypeStruct]
	_, hostIpExists := fields[mSHostIpStruct]
	_, portNumExists := fields[mSportNumStruct]
	_, portNumSlowExists := fields[portNumSlowStruct]
	_, baseUrlExists := fields[baseUrlStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !mSNameExists && !systemTypeExists && !hostIpExists && !portNumExists && !portNumSlowExists && !baseUrlExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s and %s and %s and %s", mSNameStruct, systemTypeStruct, mSHostIpStruct, mSportNumStruct, portNumSlowStruct, baseUrlStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewMSServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataUpdateMS, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataUpdateMS, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.DebugMetadataUpdateMS, id)
}
