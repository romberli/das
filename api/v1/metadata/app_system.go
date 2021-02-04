package metadata

import (
	"github.com/gin-gonic/gin"
)

// @Tags application system
// @Summary get all application systems
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/app-system [get]
func GetAppSystem(c *gin.Context) {

}

// @Tags application system
// @Summary get application system by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/app-system/:id [get]
func GetAppSystemByID(c *gin.Context) {

}

// @Tags application system
// @Summary add a new application system
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/app-system [post]
func AddAppSystem(c *gin.Context) {

}

// @Tags application system
// @Summary update application system by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/app-system/:id [post]
func UpdateAppSystemByID(c *gin.Context) {

}
