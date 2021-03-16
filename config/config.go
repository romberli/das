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
	ValidLogLevels = []string{"debug", "info", "warn", "warning", "error", "fatal"}
	ValidLogFormat = []string{"text", "json"}
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
	viper.SetDefault(DBMySQLAddrKey, fmt.Sprintf("%s:%d", constant.DefaultLocalHostIP, constant.DefaultMySQLPort))
	viper.SetDefault(DBMySQLNameKey, DefaultDBName)
	viper.SetDefault(DBMySQLUserKey, DefaultDBUser)
	viper.SetDefault(DBMySQLPassKey, DefaultDBPass)
	viper.SetDefault(DBPoolMaxConnectionsKey, mysql.DefaultMaxConnections)
	viper.SetDefault(DBPoolInitConnectionsKey, mysql.DefaultInitConnections)
	viper.SetDefault(DBPoolMaxIdleConnectionsKey, mysql.DefaultMaxIdleConnections)
	viper.SetDefault(DBPoolMaxIdleTimeKey, mysql.DefaultMaxIdleTime)
	viper.SetDefault(DBPoolKeepAliveIntervalKey, mysql.DefaultKeepAliveInterval)
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
	valid, err = common.ElementInSlice(logLevel, ValidLogLevels)
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
	valid, err = common.ElementInSlice(logFormat, ValidLogFormat)
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

func ValidateDatabase() error {
	merr := &multierror.Error{}

	// validate db.addr
	dbAddr, err := cast.ToStringE(viper.Get(DBMySQLAddrKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	addr := strings.Split(dbAddr, ":")
	if len(addr) != 2 {
		merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBAddr].Renew(dbAddr))
	} else {
		if !govalidator.IsIPv4(addr[0]) {
			merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBAddr].Renew(dbAddr))
		}
		if !govalidator.IsPort(addr[1]) {
			merr = multierror.Append(merr, message.Messages[message.ErrNotValidDBAddr].Renew(dbAddr))
		}
	}
	// validate db.name
	_, err = cast.ToStringE(viper.Get(DBMySQLNameKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.user
	_, err = cast.ToStringE(viper.Get(DBMySQLUserKey))
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate db.pass
	_, err = cast.ToStringE(viper.Get(DBMySQLPassKey))
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

	return merr
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
