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
	monitorSystemNameStruct        = "SystemName"
	monitorSystemTypeStruct        = "SystemType"
	monitorSystemHostIpStruct      = "HostIp"
	monitorSystemPortNumStruct     = "PortNum"
	monitorSystemPortNumSlowStruct = "PortNumSlow"
	monitorSystemBaseUrlStruct     = "BaseUrl"
)

// @Tags monitor system
// @Summary get all monitor systems
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system [get]
func GetMonitorSystem(c *gin.Context) {
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetMonitorSystemAll, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetMonitorSystemAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMonitorSystemAll)
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
	s := metadata.NewMonitorSystemServiceWithDefault()
	// get entity
	err := s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetMonitorSystemByID, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataGetMonitorSystemByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetMonitorSystemByID, id)
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MonitorSystemInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, systemNameExists := fields[monitorSystemNameStruct]
	_, systemTypeExists := fields[monitorSystemTypeStruct]
	_, hostIpExists := fields[monitorSystemHostIpStruct]
	_, portNumExists := fields[monitorSystemPortNumStruct]
	_, portNumSlowExists := fields[monitorSystemPortNumSlowStruct]
	_, baseUrlExists := fields[monitorSystemBaseUrlStruct]
	if !systemNameExists && !systemTypeExists && !hostIpExists && !portNumExists && !portNumSlowExists && !baseUrlExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s and %s and %s", monitorSystemNameStruct, monitorSystemTypeStruct, monitorSystemHostIpStruct, monitorSystemPortNumStruct, monitorSystemPortNumSlowStruct, monitorSystemBaseUrlStruct))
		return
	}
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataAddMS, fmt.Sprintf("%s and %s and %s and %s and %s and %s", monitorSystemNameStruct, monitorSystemTypeStruct, monitorSystemHostIpStruct, monitorSystemPortNumStruct, monitorSystemPortNumSlowStruct, monitorSystemBaseUrlStruct), err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, fmt.Sprintf("%s and %s and %s and %s and %s and %s", monitorSystemNameStruct, monitorSystemTypeStruct, monitorSystemHostIpStruct, monitorSystemPortNumStruct, monitorSystemPortNumSlowStruct, monitorSystemBaseUrlStruct), err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataAddMonitorSystem, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataAddMonitorSystem, fmt.Sprintf("%s and %s and %s and %s and %s and %s", monitorSystemNameStruct, monitorSystemTypeStruct, monitorSystemHostIpStruct, monitorSystemPortNumStruct, monitorSystemPortNumSlowStruct, monitorSystemBaseUrlStruct))
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MonitorSystemInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, systemNameExists := fields[monitorSystemNameStruct]
	_, systemTypeExists := fields[monitorSystemTypeStruct]
	_, hostIpExists := fields[monitorSystemHostIpStruct]
	_, portNumExists := fields[monitorSystemPortNumStruct]
	_, portNumSlowExists := fields[monitorSystemPortNumSlowStruct]
	_, baseUrlExists := fields[monitorSystemBaseUrlStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !systemNameExists && !systemTypeExists && !hostIpExists && !portNumExists && !portNumSlowExists && !baseUrlExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s and %s and %s and %s", monitorSystemNameStruct, monitorSystemTypeStruct, monitorSystemHostIpStruct, monitorSystemPortNumStruct, monitorSystemPortNumSlowStruct, monitorSystemBaseUrlStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataUpdateMonitorSystem, id, err.Error())
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
	log.Debug(message.NewMessage(message.DebugMetadataUpdateMonitorSystem, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.DebugMetadataUpdateMonitorSystem, id)
}
