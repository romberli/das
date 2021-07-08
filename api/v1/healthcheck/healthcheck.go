package healthcheck

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/romberli/das/internal/app/healthcheck"
	"github.com/romberli/das/pkg/message"
	msghealth "github.com/romberli/das/pkg/message/healthcheck"
	"github.com/romberli/das/pkg/resp"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
)

const (
	operationIDJSON = "operation_id"
	serverIDJSON    = "server_id"
	hostIPJSON      = "host_ip"
	portNumJSON     = "port_num"
	startTimeJSON   = "start_time"
	endTimeJSON     = "end_time"
	stepJSON        = "step"
	reviewJSON      = "review"
)

// @Tags healthcheck
// @Summary get result by operation id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/healthcheck/result/:id [get]
func GetResultByOperationID(c *gin.Context) {
	// get data
	operationIDStr := c.Param(operationIDJSON)
	if operationIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, operationIDJSON)
		return
	}
	operationID, err := strconv.Atoi(operationIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// get entities
	err = s.GetResultByOperationID(operationID)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckGetResultByOperationID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalJSON()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msghealth.DebugHealthcheckGetResultByOperationID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msghealth.InfoHealthcheckGetResultByOperationID, operationID)
}

// @Tags healthcheck
// @Summary check health of the database
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": "healthcheck started.}"
// @Router /api/v1/healthcheck/check [post]
func Check(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	dataMap := make(map[string]string)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	mysqlServerIDStr, mysqlServerIDExists := dataMap[serverIDJSON]
	if !mysqlServerIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, serverIDJSON)
		return
	}
	mysqlServerID, err := strconv.Atoi(mysqlServerIDStr)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheck, err.Error())
	}
	startTimeStr, startTimeExists := dataMap[startTimeJSON]
	if !startTimeExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, startTimeJSON)
		return
	}
	startTime, err := time.ParseInLocation(constant.TimeLayoutSecond, startTimeStr, time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, startTimeStr)
	}
	endTimeStr, endTimeExists := dataMap[endTimeJSON]
	if !endTimeExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, endTimeJSON)
		return
	}
	endTime, err := time.ParseInLocation(constant.TimeLayoutSecond, endTimeStr, time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, endTimeStr)
	}
	stepStr, stepExists := dataMap[stepJSON]
	if !stepExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, stepJSON)
		return
	}
	step, err := time.ParseDuration(stepStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeDuration, step)
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// check health
	err = s.Check(mysqlServerID, startTime, endTime, step)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheck)
		return
	}
	respMessage := "healthcheck started"
	log.Debug(message.NewMessage(msghealth.DebugHealthcheckCheck, respMessage).Error())
	resp.ResponseOK(c, respMessage, msghealth.InfoHealthcheckCheck)
}

// @Tags healthcheck
// @Summary check health of the database by host ip and port number
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": ""}"
// @Router /api/v1/healthcheck/check/host-info [post]
func CheckByHostInfo(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	dataMap := make(map[string]string)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheck, err.Error())
		return
	}
	hostIP, hostIPExists := dataMap[hostIPJSON]
	if !hostIPExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hostIPJSON)
		return
	}
	portNumStr, portNumExists := dataMap[portNumJSON]
	if !portNumExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, portNumJSON)
		return
	}
	portNum, err := strconv.Atoi(portNumStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	startTimeStr, startTimeExists := dataMap[startTimeJSON]
	if !startTimeExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, startTimeJSON)
		return
	}
	startTime, err := time.ParseInLocation(constant.TimeLayoutSecond, startTimeStr, time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, startTimeStr)
	}
	endTimeStr, endTimeExists := dataMap[endTimeJSON]
	if !endTimeExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, endTimeJSON)
		return
	}
	endTime, err := time.ParseInLocation(constant.TimeLayoutSecond, endTimeStr, time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, endTimeStr)
	}
	stepStr, stepExists := dataMap[stepJSON]
	if !stepExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, stepJSON)
		return
	}
	step, err := time.ParseDuration(stepStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeDuration, step)
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// get entities
	err = s.CheckByHostInfo(hostIP, portNum, startTime, endTime, step)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheckByHostInfo)
		return
	}
	respMessage := "healthcheck by host info started"
	log.Debug(message.NewMessage(msghealth.DebugHealthcheckCheckByHostInfo, respMessage).Error())
	resp.ResponseOK(c, respMessage, msghealth.InfoHealthcheckCheckByHostInfo)
}

// @Tags healthcheck
// @Summary update accurate review
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": ""}"
// @Router /api/v1/healthcheck/review [post]
func ReviewAccurate(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	dataMap := make(map[string]int)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheck, err.Error())
		return
	}
	operationID, operationIDExists := dataMap[operationIDJSON]
	if !operationIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, operationIDJSON)
		return
	}
	review, reviewExists := dataMap[reviewJSON]
	if !reviewExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, reviewJSON)
		return
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// review accurate
	err = s.ReviewAccurate(operationID, review)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckReviewAccurate)
		return
	}
	respMessage := "reviewed accurate"
	log.Debug(message.NewMessage(msghealth.DebugHealthcheckReviewAccurate, respMessage).Error())
	resp.ResponseOK(c, respMessage, msghealth.InfoHealthcheckReviewAccurate)
}
