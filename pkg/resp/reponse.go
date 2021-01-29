package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romberli/log"

	"github.com/romberli/das/pkg/message"
)

// ResponseNOK responses with given code and values,
// if code is between 400000 and 500000, it will log error and resp 500 to client
// otherwise, it will log info and resp 200 to client
func ResponseNOK(c *gin.Context, code int, values ...interface{}) {
	msg := message.NewMessage(code, values...).Error()
	log.Error(msg)

	c.String(http.StatusInternalServerError, msg)
}

func ResponseOK(c *gin.Context, respMessage string, code int, values ...interface{}) {
	msg := message.NewMessage(code, values...).Error()
	log.Info(msg)

	c.String(http.StatusOK, respMessage)
}
