package global

import (
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/viper"
)

var MySQLPool *mysql.Pool

func InitMySQLPool() (err error) {
	dbAddr := viper.GetString("db.mysql.addr")
	dbName := viper.GetString("db.mysql.name")
	dbUser := viper.GetString("db.mysql.user")
	dbPass := viper.GetString("db.mysql.pass")
	maxConnections := viper.GetInt("db.pool.maxConnections")
	initConnections := viper.GetInt("db.pool.initConnections")
	maxIdleConnections := viper.GetInt("db.pool.maxIdleConnections")
	maxIdleTime := viper.GetInt("db.pool.maxIdleTime")
	keepAliveInterval := viper.GetInt("db.pool.keepAliveInterval")

	config := mysql.NewMySQLConfig(dbAddr, dbName, dbUser, dbPass)
	poolConfig := mysql.NewPoolConfigWithConfig(config, maxConnections, initConnections, maxIdleConnections, maxIdleTime, keepAliveInterval)
	log.Debugf("pool config: %v", poolConfig)
	MySQLPool, err = mysql.NewMySQLPoolWithPoolConfig(poolConfig)

	return err
}
