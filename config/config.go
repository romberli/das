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
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
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
	viper.SetDefault(LogFileNameKey, defaultLogFile)
	viper.SetDefault(LogLevelKey, log.DefaultLogLevel)
	viper.SetDefault(LogFormatKey, log.DefaultLogFormat)
	viper.SetDefault(LogMaxSizeKey, log.DefaultLogMaxSize)
	viper.SetDefault(LogMaxDaysKey, log.DefaultLogMaxDays)
	viper.SetDefault(LogMaxBackupsKey, log.DefaultLogMaxBackups)
	// server
	viper.SetDefault(ServerPortKey, DefaultServerPort)
	defaultPidFile := filepath.Join(baseDir, fmt.Sprintf("%s.pid", DefaultCommandName))
	viper.SetDefault(ServerPidFileKey, defaultPidFile)
}

// ValidateConfig validates if the configuration is valid
func ValidateConfig() (err error) {
	// validate daemon section
	err = ValidateDaemon()
	if err != nil {
		return err
	}

	// validate log section
	err = ValidateLog()
	if err != nil {
		return err
	}

	// validate server section
	err = ValidateServer()
	if err != nil {
		return err
	}

	return nil
}

// ValidateDaemon validates if daemon section is valid
func ValidateDaemon() error {
	_, err := cast.ToBoolE(viper.Get(DaemonKey))

	return err
}

// ValidateLog validates if log section is valid.
func ValidateLog() error {
	var valid bool

	// validate log.FileName
	logFileName, err := cast.ToStringE(viper.Get(LogFileNameKey))
	if err != nil {
		return err
	}
	logFileName = strings.TrimSpace(logFileName)
	if logFileName == constant.EmptyString {
		return errors.New(fmt.Sprintf(ErrEmptyLogFileName))
	}
	isAbs := filepath.IsAbs(logFileName)
	if !isAbs {
		logFileName, err = filepath.Abs(logFileName)
		if err != nil {
			return err
		}
	}
	valid, _ = govalidator.IsFilePath(logFileName)
	if !valid {
		return errors.New(fmt.Sprintf(ErrNotValidLogFileName, logFileName))
	}

	// validate log.level
	logLevel, err := cast.ToStringE(viper.Get(LogLevelKey))
	if err != nil {
		return err
	}
	valid, err = common.ElementInSlice(logLevel, ValidLogLevels)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New(fmt.Sprintf(ErrNotValidLogLevel, logLevel))
	}

	// validate log.format
	logFormat, err := cast.ToStringE(viper.Get(LogFormatKey))
	if err != nil {
		return err
	}
	valid, err = common.ElementInSlice(logFormat, ValidLogFormat)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New(fmt.Sprintf(ErrNotValidLogFormat, logFormat))
	}

	// validate log.maxSize
	logMaxSize, err := cast.ToIntE(viper.Get(LogMaxSizeKey))
	if err != nil {
		return err
	}
	if logMaxSize < MinLogMaxSize || logMaxSize > MaxLogMaxSize {
		return errors.New(fmt.Sprintf(ErrNotValidLogMaxSize, MinLogMaxSize, MaxLogMaxSize, logMaxSize))
	}

	// validate log.maxDays
	logMaxDays, err := cast.ToIntE(viper.Get(LogMaxDaysKey))
	if err != nil {
		return err
	}
	if logMaxDays < MinLogMaxDays || logMaxDays > MaxLogMaxDays {
		return errors.New(fmt.Sprintf(ErrNotValidLogMaxDays, MinLogMaxDays, MaxLogMaxDays, logMaxDays))
	}

	// validate log.maxBackups
	logMaxBackups, err := cast.ToIntE(viper.Get(LogMaxBackupsKey))
	if err != nil {
		return err
	}
	if logMaxBackups < MinLogMaxDays || logMaxBackups > MaxLogMaxDays {
		return errors.New(fmt.Sprintf(ErrNotValidLogMaxBackups, MinLogMaxBackups, MaxLogMaxBackups, logMaxDays))
	}

	return nil
}

// ValidateServer validates if server section is valid
func ValidateServer() error {
	// validate server.port
	serverPort, err := cast.ToIntE(viper.Get(ServerPortKey))
	if err != nil {
		return err
	}
	if !govalidator.IsPort(strconv.Itoa(serverPort)) {
		return errors.New(fmt.Sprintf(ErrNotValidServerPort, constant.MinPort, constant.MaxPort, serverPort))
	}

	// validate server.pidFile
	serverPidFile, err := cast.ToStringE(viper.Get(ServerPidFileKey))
	if err != nil {
		return err
	}
	isAbs := filepath.IsAbs(serverPidFile)
	if !isAbs {
		serverPidFile, err = filepath.Abs(serverPidFile)
		if err != nil {
			return err
		}
	}
	ok, _ := govalidator.IsFilePath(serverPidFile)
	if !ok {
		return errors.New(fmt.Sprintf(ErrNotValidPidFile, serverPidFile))
	}

	return nil
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
