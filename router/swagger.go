package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/romberli/das/docs"
)

func RegisterSwagger(group *gin.RouterGroup) {
	swaggerGroup := group.Group("/swagger")
	{
		url := ginSwagger.URL(fmt.Sprintf("%s/doc.json", swaggerGroup.BasePath()))
		swaggerGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
