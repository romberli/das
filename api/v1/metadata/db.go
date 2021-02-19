package metadata

import (
	"github.com/gin-gonic/gin"
)

// @Tags database
// @Summary get all databases
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db [get]
func GetDB(c *gin.Context) {

}

// @Tags database
// @Summary get database by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/:id [get]
func GetDBByID(c *gin.Context) {

}

// @Tags database
// @Summary add a new database
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db [post]
func AddDB(c *gin.Context) {

}

// @Tags database
// @Summary update database by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/:id [post]
func UpdateDBByID(c *gin.Context) {

}
