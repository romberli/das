package metadata

import (
	"github.com/gin-gonic/gin"
)

// @Tags mysql server
// @Summary get all mysql servers
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/mysql-server [get]
func GetMySQLServer(c *gin.Context) {

}

// @Tags mysql server
// @Summary get mysql server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/mysql-server/:id [get]
func GetMySQLServerByID(c *gin.Context) {

}

// @Tags mysql server
// @Summary add a new mysql server
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/mysql-server [post]
func AddMySQLServer(c *gin.Context) {

}

// @Tags mysql server
// @Summary update mysql server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/mysql-server/:id [post]
func UpdateMySQLServerByID(c *gin.Context) {

}
