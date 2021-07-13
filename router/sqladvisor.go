package router

import (
	"github.com/gin-gonic/gin"
	"github.com/romberli/das/api/v1/sqladvisor"
)

func RegisterSQLAdvisor(group *gin.RouterGroup) {
	sqladvisorGroup := group.Group("/sqladvisor")
	{
		sqladvisorGroup.GET("/fingerprint", sqladvisor.GetFingerprint)
		sqladvisorGroup.GET("/sql-id", sqladvisor.GetSQLID)
		sqladvisorGroup.POST("/advise/:db_id", sqladvisor.Advise)
	}
}
