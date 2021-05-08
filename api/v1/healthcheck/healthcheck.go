package healthcheck

import (
	"github.com/gin-gonic/gin"
)

// @Tags healthcheck
// @Summary get result by operation id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/healthcheck/result/:id [get]
func GetResultByOperationID(c *gin.Context) {

}

// @Tags healthcheck
// @Summary check health of the database
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": "healthcheck started.}"
// @Router /api/v1/healthcheck/check [post]
func Check(c *gin.Context) {

}

// @Tags healthcheck
// @Summary check health of the database by host ip and port number
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": ""}"
// @Router /api/v1/healthcheck/check/host-info [post]
func CheckByHostInfo(c *gin.Context) {

}

// @Tags healthcheck
// @Summary update accurate review
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": ""}"
// @Router /api/v1/healthcheck/review [post]
func ReviewAccurate(c *gin.Context) {

}
