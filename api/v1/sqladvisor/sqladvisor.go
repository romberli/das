package sqladvisor

import (
	"github.com/gin-gonic/gin"
)

// @Tags sqladvisor
// @Summary get sql fingerprint
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"sql_id": 66, "sql_text": "select * from t01"}"
// @Router /api/v1/sqladvisor/fingerprint/:sql [get]
func GetFingerprint(c *gin.Context) {

}

// @Tags sqladvisor
// @Summary get sql id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"sql_id": 66, "sql_text": "select * from t01"}"
// @Router /api/v1/sqladvisor/sql-id/:sql [get]
func GetSQLID(c *gin.Context) {

}

// @Tags sqladvisor
// @Summary get advice
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"sql_text": "select * from t01", "advice": "xxx"}"
// @Router /api/v1/sqladvisor/advise [post]
func Advise(c *gin.Context) {

}
