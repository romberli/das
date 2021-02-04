package router

import (
	"github.com/gin-gonic/gin"

	"github.com/romberli/das/api/v1/metadata"
)

func RegisterMetadata(group *gin.RouterGroup) {
	metadataGroup := group.Group("/metadata")
	{
		// app system
		metadataGroup.GET("/app-system", metadata.GetAppSystem)
		metadataGroup.GET("/app-system/:id", metadata.GetAppSystemByID)
		metadataGroup.POST("/app-system", metadata.AddAppSystem)
		metadataGroup.POST("/app-system/:id", metadata.UpdateAppSystemByID)
		// db
		metadataGroup.GET("/db", metadata.GetDB)
		metadataGroup.GET("/db/:id", metadata.GetDBByID)
		metadataGroup.POST("/db", metadata.AddDB)
		metadataGroup.POST("/db/:id", metadata.UpdateDBByID)
		// env
		metadataGroup.GET("/env", metadata.GetEnv)
		metadataGroup.GET("/env/:id", metadata.GetEnvByID)
		metadataGroup.POST("/env", metadata.AddEnv)
		metadataGroup.POST("/env/:id", metadata.UpdateEnvByID)
		// middleware cluster
		metadataGroup.GET("/middleware-cluster", metadata.GetMiddlewareCluster)
		metadataGroup.GET("/middleware-cluster/:id", metadata.GetMiddlewareClusterByID)
		metadataGroup.POST("/middleware-cluster", metadata.AddMiddlewareCluster)
		metadataGroup.POST("/middleware-cluster/:id", metadata.UpdateMiddlewareClusterByID)
		// middleware server
		metadataGroup.GET("/middleware-server", metadata.GetMiddlewareServer)
		metadataGroup.GET("/middleware-server/:id", metadata.GetMiddlewareServerByID)
		metadataGroup.POST("/middleware-server", metadata.AddMiddlewareServer)
		metadataGroup.POST("/middleware-server/:id", metadata.UpdateMiddlewareServerByID)
		// monitor system
		metadataGroup.GET("/monitor-system", metadata.GetMonitorSystem)
		metadataGroup.GET("/monitor-system/:id", metadata.GetMonitorSystemByID)
		metadataGroup.POST("/monitor-system", metadata.AddMonitorSystem)
		metadataGroup.POST("/monitor-system/:id", metadata.UpdateMonitorSystemByID)
		// mysql cluster
		metadataGroup.GET("/mysql-cluster", metadata.GetMySQLCluster)
		metadataGroup.GET("/mysql-cluster/:id", metadata.GetMySQLClusterByID)
		metadataGroup.POST("/mysql-cluster", metadata.AddMySQLCluster)
		metadataGroup.POST("/mysql-cluster/:id", metadata.UpdateMySQLClusterByID)
		// mysql server
		metadataGroup.GET("/mysql-server", metadata.GetMySQLServer)
		metadataGroup.GET("/mysql-server/:id", metadata.GetMySQLServerByID)
		metadataGroup.POST("/mysql-server", metadata.AddMySQLServer)
		metadataGroup.POST("/mysql-server/:id", metadata.UpdateMySQLServerByID)
		// user
		metadataGroup.GET("/user", metadata.GetUser)
		metadataGroup.GET("/user/:id", metadata.GetUserByID)
		metadataGroup.POST("/user", metadata.AddUser)
		metadataGroup.POST("/user/:id", metadata.UpdateUserByID)
	}
}
