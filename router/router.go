package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
	api := gr.Engine.Group("/api")
	v1 := api.Group("/v1")
	{
		// swagger
		RegisterSwagger(v1)
		// metadata
		RegisterMetadata(v1)
	}
}

func (gr *GinRouter) Run(addr ...string) error {
	return gr.Engine.Run(addr...)
}

func (gr *GinRouter) Swagger() {

}
