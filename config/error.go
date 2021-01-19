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
	"github.com/romberli/go-util/config"
)

const (
	// info
	InfoServerStart      = 200001
	InfoServerStop       = 200002
	InfoServerIsRunning  = 200003
	InfoServerNotRunning = 200004

	// error
	ErrPrintHelpInfo            = 400001
	ErrEmptyLogFileName         = 400002
	ErrNotValidLogFileName      = 400003
	ErrNotValidLogLevel         = 400004
	ErrNotValidLogFormat        = 400005
	ErrNotValidLogMaxSize       = 400006
	ErrNotValidLogMaxDays       = 400007
	ErrNotValidLogMaxBackups    = 400008
	ErrNotValidServerPort       = 400009
	ErrNotValidPidFile          = 400010
	ErrValidateConfig           = 400011
	ErrInitDefaultConfig        = 400012
	ErrReadConfigFile           = 400013
	ErrOverrideCommandLineArgs  = 400014
	ErrAbsoluteLogFilePath      = 400015
	ErrInitLogger               = 400016
	ErrBaseDir                  = 400017
	ErrInitConfig               = 400018
	ErrCheckServerPid           = 400019
	ErrCheckServerRunningStatus = 400020
	ErrServerIsRunning          = 400021
	ErrStartAsForeground        = 400022
	ErrSavePidToFile            = 400023
	ErrKillServerWithPid        = 400024
	ErrKillServerWithPidFile    = 400025
	ErrGetPidFromPidFile        = 400026
	ErrSetSid                   = 400027
	ErrRemovePidFile            = 400028
)

var Messages = map[int]*config.ErrMessage{
	// info
	InfoServerStart:      config.NewErrMessage(DefaultErrorHeader, InfoServerStart, "das started successfully. pid: %d, pid file: %s"),
	InfoServerStop:       config.NewErrMessage(DefaultErrorHeader, InfoServerStop, "das stopped successfully. pid: %d, pid file: %s"),
	InfoServerIsRunning:  config.NewErrMessage(DefaultErrorHeader, InfoServerIsRunning, "das is running. pid: %d"),
	InfoServerNotRunning: config.NewErrMessage(DefaultErrorHeader, InfoServerNotRunning, "das is not running.pid: %d"),

	// error
	ErrPrintHelpInfo:            config.NewErrMessage(DefaultErrorHeader, ErrPrintHelpInfo, "got error when printing help information"),
	ErrEmptyLogFileName:         config.NewErrMessage(DefaultErrorHeader, ErrEmptyLogFileName, "Log file name could not be an empty string"),
	ErrNotValidLogFileName:      config.NewErrMessage(DefaultErrorHeader, ErrNotValidLogFileName, "Log file name must be either unix or windows path format, %s is not valid"),
	ErrNotValidLogLevel:         config.NewErrMessage(DefaultErrorHeader, ErrNotValidLogLevel, "Log level must be one of [debug, info, warn, error, fatal], %s is not valid"),
	ErrNotValidLogFormat:        config.NewErrMessage(DefaultErrorHeader, ErrNotValidLogFormat, "Log level must be either text or json, %s is not valid"),
	ErrNotValidLogMaxSize:       config.NewErrMessage(DefaultErrorHeader, ErrNotValidLogMaxSize, "Log max size must be between %d and %d, %d is not valid"),
	ErrNotValidLogMaxDays:       config.NewErrMessage(DefaultErrorHeader, ErrNotValidLogMaxDays, "Log max days must be between %d and %d, %d is not valid"),
	ErrNotValidLogMaxBackups:    config.NewErrMessage(DefaultErrorHeader, ErrNotValidLogMaxBackups, "Log max backups must be between %d and %d, %d is not valid"),
	ErrNotValidServerPort:       config.NewErrMessage(DefaultErrorHeader, ErrNotValidServerPort, "Server port must be between %d and %d, %d is not valid"),
	ErrNotValidPidFile:          config.NewErrMessage(DefaultErrorHeader, ErrNotValidPidFile, "pid file name must be either unix or windows path format, %s is not valid"),
	ErrValidateConfig:           config.NewErrMessage(DefaultErrorHeader, ErrValidateConfig, "validate config failed"),
	ErrInitDefaultConfig:        config.NewErrMessage(DefaultErrorHeader, ErrInitDefaultConfig, "init default configuration failed"),
	ErrReadConfigFile:           config.NewErrMessage(DefaultErrorHeader, ErrReadConfigFile, "read config file failed"),
	ErrOverrideCommandLineArgs:  config.NewErrMessage(DefaultErrorHeader, ErrOverrideCommandLineArgs, "override command line arguments failed"),
	ErrAbsoluteLogFilePath:      config.NewErrMessage(DefaultErrorHeader, ErrAbsoluteLogFilePath, "get absolute path of log file failed. log file: %s"),
	ErrInitLogger:               config.NewErrMessage(DefaultErrorHeader, ErrInitLogger, "initialize logger failed"),
	ErrBaseDir:                  config.NewErrMessage(DefaultErrorHeader, ErrBaseDir, "get base dir of %s failed"),
	ErrInitConfig:               config.NewErrMessage(DefaultErrorHeader, ErrInitConfig, "init config failed"),
	ErrCheckServerPid:           config.NewErrMessage(DefaultErrorHeader, ErrCheckServerPid, "check server pid file failed"),
	ErrCheckServerRunningStatus: config.NewErrMessage(DefaultErrorHeader, ErrCheckServerRunningStatus, "check server running status failed"),
	ErrServerIsRunning:          config.NewErrMessage(DefaultErrorHeader, ErrServerIsRunning, "pid file exists or server is still running. pid file: %s"),
	ErrStartAsForeground:        config.NewErrMessage(DefaultErrorHeader, ErrStartAsForeground, "got error when starting das as foreground"),
	ErrSavePidToFile:            config.NewErrMessage(DefaultErrorHeader, ErrSavePidToFile, "got error when writing pid to file"),
	ErrKillServerWithPid:        config.NewErrMessage(DefaultErrorHeader, ErrKillServerWithPid, "kill server failed. pid: %d"),
	ErrKillServerWithPidFile:    config.NewErrMessage(DefaultErrorHeader, ErrKillServerWithPidFile, "kill server with pid file failed. pid: %d, pid file: %s"),
	ErrGetPidFromPidFile:        config.NewErrMessage(DefaultErrorHeader, ErrGetPidFromPidFile, "get pid from pid file failed. pid file: %s"),
	ErrSetSid:                   config.NewErrMessage(DefaultErrorHeader, ErrSetSid, "set sid failed when daemonizing server"),
	ErrRemovePidFile:            config.NewErrMessage(DefaultErrorHeader, ErrRemovePidFile, "remove pid file failed"),
}
