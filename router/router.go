package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/romberli/das/api/v1/metadata"
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
	api := gr.Engine.Group("/api")
	v1 := api.Group("/v1")
	{
		// metadata
		metadataGroup := v1.Group("/metadata")
		{
			// env
			metadataGroup.GET("/env", metadata.GetEnv)
			metadataGroup.GET("/env/:id", metadata.GetEnvByID)
			metadataGroup.POST("/env", metadata.AddEnv)
			metadataGroup.POST("/env/:id", metadata.UpdateEnvByID)
		}
	}
}

func (gr *GinRouter) Run(addr ...string) error {
	return gr.Engine.Run(addr...)
}
