package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/romberli/das/docs"
)

type Router interface {
	http.Handler
	Register()
	Run(addr ...string) error
}

type GinRouter struct {
	Engine *gin.Engine
}

func NewGinRouter(r *gin.Engine) *GinRouter {
	return &GinRouter{r}
}

func (gr *GinRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	gr.Engine.ServeHTTP(w, req)
}

func (gr *GinRouter) Register() {
	// swagger
	gr.Swagger()

	api := gr.Engine.Group("/api")
	v1 := api.Group("/v1")
	{
		// metadata
		RegisterMetadata(v1)
	}
}

func (gr *GinRouter) Run(addr ...string) error {
	return gr.Engine.Run(addr...)
}

func (gr *GinRouter) Swagger() {
	swaggerGroup := gr.Engine.Group("/swagger")
	{
		url := ginSwagger.URL("/swagger/doc.json")
		swaggerGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
