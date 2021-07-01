/*
Copyright © 2020 Romber Li <romber2001@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/hashicorp/go-multierror"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/romberli/das/pkg/message"
)

var (
	ValidLogLevels  = []string{"debug", "info", "warn", "warning", "error", "fatal"}
	ValidLogFormats = []string{"text", "json"}
)

// SetDefaultConfig set default configuration, it is the lowest priority
func SetDefaultConfig(baseDir string) {
	// daemon
	viper.SetDefault(DaemonKey, DefaultDaemon)
	// log
	defaultLogFile := filepath.Join(baseDir, DefaultLogDir, log.DefaultLogFileName)
	viper.SetDefault(LogFileKey, defaultLogFile)
	viper.SetDefault(LogLevelKey, log.DefaultLogLevel)
	viper.SetDefault(LogFormatKey, log.DefaultLogFormat)
	viper.SetDefault(LogMaxSizeKey, log.DefaultLogMaxSize)
	viper.SetDefault(LogMaxDaysKey, log.DefaultLogMaxDays)
	viper.SetDefault(LogMaxBackupsKey, log.DefaultLogMaxBackups)
	// server
	viper.SetDefault(ServerAddrKey, DefaultServerAddr)
	defaultPidFile := filepath.Join(baseDir, fmt.Sprintf("%s.pid", DefaultCommandName))
	viper.SetDefault(ServerPidFileKey, defaultPidFile)
	viper.SetDefault(ServerReadTimeoutKey, DefaultServerReadTimeout)
	viper.SetDefault(ServerWriteTimeoutKey, DefaultServerWriteTimeout)
	// database
	viper.SetDefault(DBDASMySQLAddrKey, fmt.Sprintf("%s:%d", constant.DefaultLocalHostIP, constant.DefaultMySQLPort))
	viper.SetDefault(DBDASMySQLNameKey, DefaultDBDASMySQLName)
	viper.SetDefault(DBDASMySQLUserKey, DefaultDBUser)
	viper.SetDefault(DBDASMySQLPassKey, DefaultDBPass)
	viper.SetDefault(DBPoolMaxConnectionsKey, mysql.DefaultMaxConnections)
	viper.SetDefault(DBPoolInitConnectionsKey, mysql.DefaultInitConnections)
	viper.SetDefault(DBPoolMaxIdleConnectionsKey, mysql.DefaultMaxIdleConnections)
	viper.SetDefault(DBPoolMaxIdleTimeKey, mysql.DefaultMaxIdleTime)
	viper.SetDefault(DBPoolKeepAliveIntervalKey, mysql.DefaultKeepAliveInterval)
	viper.SetDefault(DBMonitorPrometheusUserKey, DefaultDBMonitorPrometheusUser)
	viper.SetDefault(DBMonitorPrometheusPassKey, DefaultDBMonitorPrometheusPass)
	viper.SetDefault(DBMonitorClickhouseUserKey, DefaultDBMonitorClickhouseUser)
	viper.SetDefault(DBMonitorClickhousePassKey, DefaultDBMonitorClickhousePass)
	viper.SetDefault(DBMonitorMySQLUserKey, DefaultDBMonitorMySQLUser)
	viper.SetDefault(DBMonitorMySQLPassKey, DefaultDBMonitorMySQLPass)
	viper.SetDefault(DBApplicationMySQLUserKey, DefaultDBApplicationMySQLUser)
	viper.SetDefault(DBApplicationMySQLPassKey, DefaultDBApplicationMySQLPass)
	viper.SetDefault(DBSoarMySQLAddrKey, fmt.Sprintf("%s:%d", constant.DefaultLocalHostIP, constant.DefaultMySQLPort))
	viper.SetDefault(DBSoarMySQLNameKey, DefaultDBDASMySQLName)
	viper.SetDefault(DBSoarMySQLUserKey, DefaultDBUser)
	viper.SetDefault(DBSoarMySQLPassKey, DefaultDBPass)
	// sqladvisor
	viper.SetDefault(SQLAdvisorSoarBin, DefaultSQLAdvisorSoarBin)
	viper.SetDefault(SQLAdvisorSoarConfig, DefaultSQLAdvisorSoarConfig)
	viper.SetDefault(SQLAdvisorSoarSamplingKey, false)
	viper.SetDefault(SQLAdvisorSoarProfilingKey, false)
	viper.SetDefault(SQLAdvisorSoarTraceKey, false)
	viper.SetDefault(SQLAdvisorSoarExplainKey, false)
}

// ValidateConfig validates if the configuration is valid
func ValidateConfig() (err error) {
	merr := &multierror.Error{}

	// validate daemon section
	err = ValidateDaemon()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	// validate log section
	err = ValidateLog()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	// validate server section
	err = ValidateServer()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	// validate database section
	err = ValidateDatabase()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	// validate soar section
	err = ValidateSQLAdvisor()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	return merr.ErrorOrNil()
}

// ValidateDaemon validates if daemon section is valid
func ValidateDaemon() error {
	_, err := cast.ToBoolE(viper.Get(DaemonKey))

	return err
}

// ValidateLog validates if log section is valid.
func ValidateLog() error {
	var valid bool

	merr := &multierror.Error{}

	// validate log.FileName
	logFileName, err := cast.ToStringE(viper.Get(LogFileKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	logFileName = strings.TrimSpace(logFileName)
	if logFileName == constant.EmptyString {
		merr = multierror.Append(merr, message.Messages[message.ErrEmptyLogFileName])
	}
	isAbs := filepath.IsAbs(logFileName)
	if !isAbs {
		logFileName, err = filepath.Abs(logFileName)
		if err != nil {
			merr = multierror.Append(merr, err)
		}
	}
	valid, _ = govalidator.IsFilePath(logFileName)
	if !valid {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidLogFileName].Renew(logFileName))
	}

	// validate log.level
	logLevel, err := cast.ToStringE(viper.Get(LogLevelKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	valid, err = common.ElementInSlice(ValidLogLevels, logLevel)
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if !valid {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidLogLevel].Renew(logLevel))
	}

	// validate log.format
	logFormat, err := cast.ToStringE(viper.Get(LogFormatKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	valid, err = common.ElementInSlice(ValidLogFormats, logFormat)
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if !valid {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidLogFormat].Renew(logFormat))
	}

	// validate log.maxSize
	logMaxSize, err := cast.ToIntE(viper.Get(LogMaxSizeKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if logMaxSize < MinLogMaxSize || logMaxSize > MaxLogMaxSize {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidLogMaxSize].Renew(MinLogMaxSize, MaxLogMaxSize, logMaxSize))
	}

	// validate log.maxDays
	logMaxDays, err := cast.ToIntE(viper.Get(LogMaxDaysKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if logMaxDays < MinLogMaxDays || logMaxDays > MaxLogMaxDays {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidLogMaxDays].Renew(MinLogMaxDays, MaxLogMaxDays, logMaxDays))
	}

	// validate log.maxBackups
	logMaxBackups, err := cast.ToIntE(viper.Get(LogMaxBackupsKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if logMaxBackups < MinLogMaxDays || logMaxBackups > MaxLogMaxDays {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidLogMaxBackups].Renew(MinLogMaxBackups, MaxLogMaxBackups, logMaxBackups))
	}

	return merr.ErrorOrNil()
}

// ValidateServer validates if server section is valid
func ValidateServer() error {
	merr := &multierror.Error{}

	// validate server.addr
	serverAddr, err := cast.ToStringE(viper.Get(ServerAddrKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	serverAddrList := strings.Split(serverAddr, ":")

	switch len(serverAddrList) {
	case 2:
		port := serverAddrList[1]
		if !govalidator.IsPort(port) {
			merr = multierror.Append(merr, message.Messages[message.ErrNotValidServerPort].Renew(constant.MinPort, constant.MaxPort, port))
		}
	default:
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidServerAddr].Renew(serverAddr))
	}

	// validate server.pidFile
	serverPidFile, err := cast.ToStringE(viper.Get(ServerPidFileKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	isAbs := filepath.IsAbs(serverPidFile)
	if !isAbs {
		serverPidFile, err = filepath.Abs(serverPidFile)
		if err != nil {
			merr = multierror.Append(merr, err)
		}
	}
	ok, _ := govalidator.IsFilePath(serverPidFile)
	if !ok {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidPidFile].Renew(serverPidFile))
	}

	// validate server.readTimeout
	serverReadTimeout, err := cast.ToIntE(viper.Get(ServerReadTimeoutKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if serverReadTimeout < MinServerReadTimeout || serverReadTimeout > MaxServerReadTimeout {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidServerPort].Renew(MinServerReadTimeout, MaxServerWriteTimeout, serverReadTimeout))
	}

	// validate server.writeTimeout
	serverWriteTimeout, err := cast.ToIntE(viper.Get(ServerWriteTimeoutKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if serverWriteTimeout < MinServerWriteTimeout || serverWriteTimeout > MaxServerWriteTimeout {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidServerPort].Renew(MinServerReadTimeout, MaxServerWriteTimeout, serverWriteTimeout))
	}

	return merr.ErrorOrNil()
}

// ValidateDatabase validates if database section is valid
func ValidateDatabase() error {
	merr := &multierror.Error{}

	// validate db.das.mysql.addr
	dbDASAddr, err := cast.ToStringE(viper.Get(DBDASMySQLAddrKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	dasAddr := strings.Split(dbDASAddr, ":")
	if len(dasAddr) != 2 {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBAddr].Renew(dbDASAddr))
	} else {
		if !govalidator.IsIPv4(dasAddr[0]) {
			merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBAddr].Renew(dbDASAddr))
		}
		if !govalidator.IsPort(dasAddr[1]) {
			merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBAddr].Renew(dbDASAddr))
		}
	}
	// validate db.das.mysql.name
	_, err = cast.ToStringE(viper.Get(DBDASMySQLNameKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.das.mysql.user
	_, err = cast.ToStringE(viper.Get(DBDASMySQLUserKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.das.mysql.pass
	_, err = cast.ToStringE(viper.Get(DBDASMySQLPassKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.pool.maxConnections
	maxConnections, err := cast.ToIntE(viper.Get(DBPoolMaxConnectionsKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if maxConnections < MinDBPoolMaxConnections || maxConnections > MaxDBPoolMaxConnections {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBPoolMaxConnections].Renew(MinDBPoolMaxConnections, MaxDBPoolMaxConnections, maxConnections))
	}
	// validate db.pool.initConnections
	initConnections, err := cast.ToIntE(viper.Get(DBPoolInitConnectionsKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if initConnections < MinDBPoolInitConnections || initConnections > MaxDBPoolInitConnections {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBPoolInitConnections].Renew(MinDBPoolInitConnections, MaxDBPoolInitConnections, initConnections))
	}
	// validate db.pool.maxIdleConnections
	maxIdleConnections, err := cast.ToIntE(viper.Get(DBPoolMaxIdleConnectionsKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if maxIdleConnections < MinDBPoolMaxIdleConnections || maxIdleConnections > MaxDBPoolMaxIdleConnections {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBPoolMaxIdleConnections].Renew(MinDBPoolMaxIdleConnections, MaxDBPoolMaxIdleConnections, maxIdleConnections))
	}
	// validate db.pool.maxIdleTime
	maxIdleTime, err := cast.ToIntE(viper.Get(DBPoolMaxIdleTimeKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if maxIdleTime < MinDBPoolMaxIdleTime || maxIdleTime > MaxDBPoolMaxIdleTime {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBPoolMaxIdleTime].Renew(MinDBPoolMaxIdleTime, MaxDBPoolMaxIdleTime, maxIdleTime))
	}
	// validate db.pool.keepAliveInterval
	keepAliveInterval, err := cast.ToIntE(viper.Get(DBPoolKeepAliveIntervalKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if keepAliveInterval < MinDBPoolKeepAliveInterval || keepAliveInterval > MaxDBPoolKeepAliveInterval {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBPoolKeepAliveInterval].Renew(MinDBPoolKeepAliveInterval, MaxDBPoolKeepAliveInterval, keepAliveInterval))
	}
	// validate db.application.mysql.user
	_, err = cast.ToStringE(viper.Get(DBApplicationMySQLUserKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.application.mysql.pass
	_, err = cast.ToStringE(viper.Get(DBApplicationMySQLPassKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.monitor.prometheus.user
	_, err = cast.ToStringE(viper.Get(DBMonitorPrometheusUserKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.monitor.prometheus.pass
	_, err = cast.ToStringE(viper.Get(DBMonitorPrometheusPassKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.monitor.clickhouse.user
	_, err = cast.ToStringE(viper.Get(DBMonitorClickhouseUserKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.monitor.clickhouse.pass
	_, err = cast.ToStringE(viper.Get(DBMonitorClickhousePassKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.monitor.mysql.user
	_, err = cast.ToStringE(viper.Get(DBMonitorMySQLUserKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.monitor.mysql.pass
	_, err = cast.ToStringE(viper.Get(DBMonitorMySQLPassKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.soar.mysql.addr
	dbSoarAddr, err := cast.ToStringE(viper.Get(DBDASMySQLAddrKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	soarAddr := strings.Split(dbSoarAddr, ":")
	if len(soarAddr) != 2 {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBAddr].Renew(dbSoarAddr))
	} else {
		if !govalidator.IsIPv4(soarAddr[0]) {
			merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBAddr].Renew(dbSoarAddr))
		}
		if !govalidator.IsPort(soarAddr[1]) {
			merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBAddr].Renew(dbSoarAddr))
		}
	}
	// validate db.soar.mysql.name
	_, err = cast.ToStringE(viper.Get(DBSoarMySQLNameKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.soar.mysql.user
	_, err = cast.ToStringE(viper.Get(DBSoarMySQLUserKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.soar.mysql.pass
	_, err = cast.ToStringE(viper.Get(DBSoarMySQLPassKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	return merr.ErrorOrNil()
}

// ValidateSQLAdvisor validates if sqladvisor section is valid
func ValidateSQLAdvisor() error {
	merr := &multierror.Error{}

	// validate sqladvisor.soar.bin
	soarBin, err := cast.ToStringE(viper.Get(SQLAdvisorSoarBin))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	soarBin = strings.TrimSpace(soarBin)
	if soarBin == constant.EmptyString {
		merr = multierror.Append(merr, message.Messages[message.ErrEmptySoarBin])
	}
	isAbs := filepath.IsAbs(soarBin)
	if !isAbs {
		soarBin, err = filepath.Abs(soarBin)
		if err != nil {
			merr = multierror.Append(merr, err)
		}
	}
	valid, _ := govalidator.IsFilePath(soarBin)
	if !valid {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidSoarBin].Renew(soarBin))
	}

	// validate sqladvisor.soar.config
	config, err := cast.ToStringE(viper.Get(SQLAdvisorSoarConfig))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	config = strings.TrimSpace(config)
	if soarBin == constant.EmptyString {
		merr = multierror.Append(merr, message.Messages[message.ErrEmptySoarConfig])
	}
	isAbs = filepath.IsAbs(config)
	if !isAbs {
		config, err = filepath.Abs(config)
		if err != nil {
			merr = multierror.Append(merr, err)
		}
	}
	valid, _ = govalidator.IsFilePath(config)
	if !valid {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidSoarConfig].Renew(config))
	}

	// validate sqladvisor.soar.sampling
	_, err = cast.ToBoolE(viper.Get(SQLAdvisorSoarSamplingKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate sqladvisor.soar.profiling
	_, err = cast.ToBoolE(viper.Get(SQLAdvisorSoarProfilingKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate sqladvisor.soar.trace
	_, err = cast.ToBoolE(viper.Get(SQLAdvisorSoarTraceKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate sqladvisor.soar.explain
	_, err = cast.ToBoolE(viper.Get(SQLAdvisorSoarExplainKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	return merr.ErrorOrNil()
}

// TrimSpaceOfArg trims spaces of given argument
func TrimSpaceOfArg(arg string) string {
	args := strings.SplitN(arg, "=", 2)

	switch len(args) {
	case 1:
		return strings.TrimSpace(args[0])
	case 2:
		argName := strings.TrimSpace(args[0])
		argValue := strings.TrimSpace(args[1])
		return fmt.Sprintf("%s=%s", argName, argValue)
	default:
		return arg
	}
}
