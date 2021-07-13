package router

import (
	"github.com/gin-gonic/gin"
	"github.com/romberli/das/api/v1/healthcheck"
)

// RegisterHealthcheck is the sub-router of das for healthcheck
func RegisterHealthcheck(group *gin.RouterGroup) {
	healthcheckGroup := group.Group("/healthcheck")
	{
		healthcheckGroup.GET("/result/:operation_id", healthcheck.GetResultByOperationID)
		healthcheckGroup.POST("/check", healthcheck.Check)
		healthcheckGroup.POST("/check/host-info", healthcheck.CheckByHostInfo)
		healthcheckGroup.POST("/review", healthcheck.ReviewAccurate)
	}
}
