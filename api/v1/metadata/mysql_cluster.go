package metadata

import (
	"github.com/gin-gonic/gin"
)

// @Tags mysql cluster
// @Summary get all mysql clusters
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/mysql-cluster [get]
func GetMySQLCluster(c *gin.Context) {

}

// @Tags mysql cluster
// @Summary get mysql cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/mysql-cluster/:id [get]
func GetMySQLClusterByID(c *gin.Context) {

}

// @Tags mysql cluster
// @Summary add a new mysql cluster
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/mysql-cluster [post]
func AddMySQLCluster(c *gin.Context) {

}

// @Tags mysql cluster
// @Summary update mysql cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/mysql-cluster/:id [post]
func UpdateMySQLClusterByID(c *gin.Context) {

}
