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
	hOperationIDJSON = "operation_id"
	hServerIDJSON    = "server_id"
	hHostIPJSON      = "host_ip"
	hPortNumJSON     = "port_num"
	hStartTimeJSON   = "start_time"
	hEndTimeJSON     = "end_time"
	hStepJSON        = "step"
	hReviewJSON      = "review"
)

// @Tags healthcheck
// @Summary get result by operation id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/healthcheck/result/:id [get]
func GetResultByOperationID(c *gin.Context) {
	// get data
	operationIDStr := c.Param(hOperationIDJSON)
	if operationIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hOperationIDJSON)
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
		resp.ResponseNOK(c, message.ErrMarshalService, err.Error())
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
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheck, err.Error())
		return
	}
	mysqlServerIDStr, mysqlServerIDExists := dataMap[hServerIDJSON]
	if !mysqlServerIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hServerIDJSON)
		return
	}
	mysqlServerID, err := strconv.Atoi(mysqlServerIDStr)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheck, err.Error())
	}
	startTimeStr, startTimeExists := dataMap[hStartTimeJSON]
	if !startTimeExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hStartTimeJSON)
		return
	}
	startTime, err := time.ParseInLocation(constant.TimeLayoutSecond, startTimeStr, time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, startTimeStr)
	}
	endTimeStr, endTimeExists := dataMap[hEndTimeJSON]
	if !endTimeExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hEndTimeJSON)
		return
	}
	endTime, err := time.ParseInLocation(constant.TimeLayoutSecond, endTimeStr, time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, endTimeStr)
	}
	stepStr, stepExists := dataMap[hStepJSON]
	if !stepExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hStepJSON)
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
	hostIP, hostIPExists := dataMap[hHostIPJSON]
	if !hostIPExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hHostIPJSON)
		return
	}
	portNumStr, portNumExists := dataMap[hPortNumJSON]
	if !portNumExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hPortNumJSON)
		return
	}
	portNum, err := strconv.Atoi(portNumStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	startTimeStr, startTimeExists := dataMap[hStartTimeJSON]
	if !startTimeExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hStartTimeJSON)
		return
	}
	startTime, err := time.ParseInLocation(constant.TimeLayoutSecond, startTimeStr, time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, startTimeStr)
	}
	endTimeStr, endTimeExists := dataMap[hEndTimeJSON]
	if !endTimeExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hEndTimeJSON)
		return
	}
	endTime, err := time.ParseInLocation(constant.TimeLayoutSecond, endTimeStr, time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, endTimeStr)
	}
	stepStr, stepExists := dataMap[hStepJSON]
	if !stepExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hStepJSON)
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
	operationID, operationIDExists := dataMap[hOperationIDJSON]
	if !operationIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hOperationIDJSON)
		return
	}
	review, reviewExists := dataMap[hReviewJSON]
	if !reviewExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, hReviewJSON)
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
