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

var (
	// info
	InfoServerStart      = "DAS-200001: das started successfully. pid: %d, pid file: %s"
	InfoServerStop       = "DAS-200002: das stopped successfully. pid: %d, pid file: %s"
	InfoServerIsRunning  = "DAS-200003: das is running. pid: %d"
	InfoServerNotRunning = "DAS-200004: das is not running.pid: %d"

	// error
	ErrPrintHelpInfo            = "DAS-400001: got error when printing help information.\n%s"
	ErrEmptyLogFileName         = "DAS-400002: Log file name could not be an empty string."
	ErrNotValidLogFileName      = "DAS-400003: Log file name must be either unix or windows path format, %s is not valid."
	ErrNotValidLogLevel         = "DAS-400004: Log level must be one of [debug, info, warn, error, fatal], %s is not valid."
	ErrNotValidLogFormat        = "DAS-400005: Log level must be either text or json, %s is not valid."
	ErrNotValidLogMaxSize       = "DAS-400006: Log max size must be between %d and %d, %d is not valid."
	ErrNotValidLogMaxDays       = "DAS-400007: Log max days must be between %d and %d, %d is not valid."
	ErrNotValidLogMaxBackups    = "DAS-400008: Log max backups must be between %d and %d, %d is not valid."
	ErrNotValidServerPort       = "DAS-400009: Server port must be between %d and %d, %d is not valid."
	ErrNotValidPidFile          = "DAS-400010: pid file name must be either unix or windows path format, %s is not valid."
	ErrValidateConfig           = "DAS-400011: validate config failed.\n%s"
	ErrInitDefaultConfig        = "DAS-400012: init default configuration failed.\n%s"
	ErrReadConfigFile           = "DAS-400013: read config file failed.\n%s"
	ErrOverrideCommandLineArgs  = "DAS-400014: override command line arguments failed.\n%s"
	ErrAbsoluteLogFilePath      = "DAS-400015: get absolute path of log file failed. log file: %s\n%s"
	ErrInitLogger               = "DAS-400016: initialize logger failed.\n%s"
	ErrBaseDir                  = "DAS-400017: get base dir of %s failed.\n%s"
	ErrInitConfig               = "DAS-400018: init config failed.\n%s"
	ErrCheckServerPid           = "DAS-400019: check server pid file failed.\n%s"
	ErrCheckServerRunningStatus = "DAS-400020: check server running status failed.\n%s"
	ErrServerIsRunning          = "DAS-400021: pid file exists or server is still running. pid file: %s."
	ErrStartAsForeground        = "DAS-400022: got error when starting das as foreground.\n%s"
	ErrSavePidToFile            = "DAS-400023: got error when writing pid to file.\n%s"
	ErrKillServerWithPid        = "DAS-400024: kill server failed. pid: %d\n%s"
	ErrKillServerWithPidFile    = "DAS-400025: kill server with pid file failed. pid: %d\n, pid file: %s.\n%s"
	ErrGetPidFromPidFile        = "DAS-400026: get pid from pid file failed. pid file: %s\n%s"
	ErrSetSid                   = "DAS-400027: set sid failed when daemonizing server.\n%s"
	ErrRemovePidFile            = "DAS-400028: remove pid file failed.\n%s"
)
