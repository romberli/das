package router

import (
	"github.com/gin-gonic/gin"
	"github.com/romberli/das/api/v1/sqladvisor"
)

func RegisterSQLAdvisor(group *gin.RouterGroup) {
	sqladvisorGroup := group.Group("/sqladvisor")
	{
		sqladvisorGroup.GET("/fingerprint/:sql", sqladvisor.GetFingerprint)
		sqladvisorGroup.GET("/sql-id/mysql-server/:id", sqladvisor.GetSQLID)
		sqladvisorGroup.POST("/advise", sqladvisor.Advise)
	}
}
