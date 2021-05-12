package router

import (
	"github.com/gin-gonic/gin"
	"github.com/romberli/das/api/v1/healthcheck"
)

func RegisterHealthcheck(group *gin.RouterGroup) {
	healthcheckGroup := group.Group("/healthcheck")
	{
		healthcheckGroup.GET("/result/:id", healthcheck.GetResultByOperationID)
		healthcheckGroup.POST("/check/mysql-server/:id", healthcheck.Check)
		healthcheckGroup.POST("/check/host-info", healthcheck.CheckByHostInfo)
		healthcheckGroup.POST("/review", healthcheck.ReviewAccurate)
	}
}
