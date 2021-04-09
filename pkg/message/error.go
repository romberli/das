package message

import (
	"github.com/romberli/go-util/config"
)

const (
	// server
	ErrPrintHelpInfo                    = 400001
	ErrEmptyLogFileName                 = 400002
	ErrNotValidLogFileName              = 400003
	ErrNotValidLogLevel                 = 400004
	ErrNotValidLogFormat                = 400005
	ErrNotValidLogMaxSize               = 400006
	ErrNotValidLogMaxDays               = 400007
	ErrNotValidLogMaxBackups            = 400008
	ErrNotValidServerPort               = 400009
	ErrNotValidPidFile                  = 400010
	ErrValidateConfig                   = 400011
	ErrInitDefaultConfig                = 400012
	ErrReadConfigFile                   = 400013
	ErrOverrideCommandLineArgs          = 400014
	ErrAbsoluteLogFilePath              = 400015
	ErrInitLogger                       = 400016
	ErrBaseDir                          = 400017
	ErrInitConfig                       = 400018
	ErrCheckServerPid                   = 400019
	ErrCheckServerRunningStatus         = 400020
	ErrServerIsRunning                  = 400021
	ErrStartAsForeground                = 400022
	ErrSavePidToFile                    = 400023
	ErrKillServerWithPid                = 400024
	ErrKillServerWithPidFile            = 400025
	ErrGetPidFromPidFile                = 400026
	ErrSetSid                           = 400027
	ErrRemovePidFile                    = 400028
	ErrNotValidDBAddr                   = 400029
	ErrNotValidDBName                   = 400030
	ErrNotValidDBUser                   = 400031
	ErrNotValidDBPass                   = 400032
	ErrNotValidDBPoolMaxConnections     = 400033
	ErrNotValidDBPoolInitConnections    = 400034
	ErrNotValidDBPoolMaxIdleConnections = 400035
	ErrNotValidDBPoolMaxIdleTime        = 400036
	ErrNotValidDBPoolKeepAliveInterval  = 400037
	ErrInitConnectionPool               = 400038
	ErrNotValidServerReadTimeout        = 400039
	ErrNotValidServerWriteTimeout       = 400040
	ErrNotValidServerAddr               = 400041
	ErrFieldNotExists                   = 400042
	ErrGetRawData                       = 400043
	ErrUnmarshalRawData                 = 400044
	ErrGenerateNewMapWithTag            = 400045
	ErrMarshalService                   = 400046
	ErrTypeConversion                   = 400047

	ErrMetadataGetEnvAll  = 400201
	ErrMetadataGetEnvByID = 400202
	ErrMetadataAddEnv     = 400203
	ErrMetadataUpdateEnv  = 400204

	ErrMetadataGetMySQLClusterAll  = 400701
	ErrMetadataGetMySQLClusterByID = 400702
	ErrMetadataAddMySQLCluster     = 400703
	ErrMetadataUpdateMySQLCluster  = 400704

	ErrMetadataGetMySQLServerAll  = 400801
	ErrMetadataGetMySQLServerByID = 400802
	ErrMetadataAddMySQLServer     = 400803
	ErrMetadataUpdateMySQLServer  = 400804

	ErrMetadataGetMiddlewareServerAll  = 400601
	ErrMetadataGetMiddlewareServerByID = 400602
	ErrMetadataAddMiddlewareServer     = 400603
	ErrMetadataUpdateMiddlewareServer  = 400604

	ErrMetadataGetUserAll  = 400901
	ErrMetadataGetUserByID = 400902
	ErrMetadataAddUser     = 400903
	ErrMetadataUpdateUser  = 400904

	ErrMetadataGetMonitorSystemAll  = 400601
	ErrMetadataGetMonitorSystemByID = 400602
	ErrMetadataAddMonitorSystem     = 400603
	ErrMetadataUpdateMonitorSystem  = 400604
)

func initErrorMessage() {
	// server
	Messages[ErrPrintHelpInfo] = config.NewErrMessage(DefaultMessageHeader, ErrPrintHelpInfo, "got message when printing help information")
	Messages[ErrEmptyLogFileName] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyLogFileName, "Log file name could not be an empty string")
	Messages[ErrNotValidLogFileName] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogFileName, "Log file name must be either unix or windows path format, %s is not valid")
	Messages[ErrNotValidLogLevel] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogLevel, "Log level must be one of [debug, info, warn, message, fatal], %s is not valid")
	Messages[ErrNotValidLogFormat] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogFormat, "Log level must be either text or json, %s is not valid")
	Messages[ErrNotValidLogMaxSize] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxSize, "Log max size must be between %d and %d, %d is not valid")
	Messages[ErrNotValidLogMaxDays] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxDays, "Log max days must be between %d and %d, %d is not valid")
	Messages[ErrNotValidLogMaxBackups] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxBackups, "Log max backups must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerPort] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerPort, "Server port must be between %d and %d, %d is not valid")
	Messages[ErrNotValidPidFile] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidPidFile, "pid file name must be either unix or windows path format, %s is not valid")
	Messages[ErrValidateConfig] = config.NewErrMessage(DefaultMessageHeader, ErrValidateConfig, "validate config failed")
	Messages[ErrInitDefaultConfig] = config.NewErrMessage(DefaultMessageHeader, ErrInitDefaultConfig, "init default configuration failed")
	Messages[ErrReadConfigFile] = config.NewErrMessage(DefaultMessageHeader, ErrReadConfigFile, "read config file failed")
	Messages[ErrOverrideCommandLineArgs] = config.NewErrMessage(DefaultMessageHeader, ErrOverrideCommandLineArgs, "override command line arguments failed")
	Messages[ErrAbsoluteLogFilePath] = config.NewErrMessage(DefaultMessageHeader, ErrAbsoluteLogFilePath, "get absolute path of log file failed. log file: %s")
	Messages[ErrInitLogger] = config.NewErrMessage(DefaultMessageHeader, ErrInitLogger, "initialize logger failed")
	Messages[ErrBaseDir] = config.NewErrMessage(DefaultMessageHeader, ErrBaseDir, "get base dir of %s failed")
	Messages[ErrInitConfig] = config.NewErrMessage(DefaultMessageHeader, ErrInitConfig, "init config failed")
	Messages[ErrCheckServerPid] = config.NewErrMessage(DefaultMessageHeader, ErrCheckServerPid, "check server pid file failed")
	Messages[ErrCheckServerRunningStatus] = config.NewErrMessage(DefaultMessageHeader, ErrCheckServerRunningStatus, "check server running status failed")
	Messages[ErrServerIsRunning] = config.NewErrMessage(DefaultMessageHeader, ErrServerIsRunning, "pid file exists or server is still running. pid file: %s")
	Messages[ErrStartAsForeground] = config.NewErrMessage(DefaultMessageHeader, ErrStartAsForeground, "got message when starting das as foreground")
	Messages[ErrSavePidToFile] = config.NewErrMessage(DefaultMessageHeader, ErrSavePidToFile, "got message when writing pid to file")
	Messages[ErrKillServerWithPid] = config.NewErrMessage(DefaultMessageHeader, ErrKillServerWithPid, "kill server failed. pid: %d")
	Messages[ErrKillServerWithPidFile] = config.NewErrMessage(DefaultMessageHeader, ErrKillServerWithPidFile, "kill server with pid file failed. pid: %d, pid file: %s")
	Messages[ErrGetPidFromPidFile] = config.NewErrMessage(DefaultMessageHeader, ErrGetPidFromPidFile, "get pid from pid file failed. pid file: %s")
	Messages[ErrSetSid] = config.NewErrMessage(DefaultMessageHeader, ErrSetSid, "set sid failed when daemonizing server")
	Messages[ErrRemovePidFile] = config.NewErrMessage(DefaultMessageHeader, ErrRemovePidFile, "remove pid file failed")
	Messages[ErrNotValidDBAddr] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBAddr, "database address must be formatted as host:port, %s is not valid")
	Messages[ErrNotValidDBName] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBName, "database name must be a string, %s is not valid")
	Messages[ErrNotValidDBUser] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBUser, "database user name must be a string, %s is not valid")
	Messages[ErrNotValidDBPass] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPass, "database password must be a string, %s is not valid")
	Messages[ErrNotValidDBPoolMaxConnections] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPoolMaxConnections, "database max connections must be between %d and %d, %d is not valid")
	Messages[ErrNotValidDBPoolInitConnections] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPoolInitConnections, "database initial connections must be between %d and %d, %d is not valid")
	Messages[ErrNotValidDBPoolMaxIdleConnections] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPoolMaxIdleConnections, "database max idle connections must be between %d and %d, %d is not valid")
	Messages[ErrNotValidDBPoolMaxIdleTime] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPoolMaxIdleTime, "database max idle time must be between %d and %d, %d is not valid")
	Messages[ErrNotValidDBPoolKeepAliveInterval] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPoolKeepAliveInterval, "database keep alive interval must be between %d and %d, %d is not valid")
	Messages[ErrInitConnectionPool] = config.NewErrMessage(DefaultMessageHeader, ErrInitConnectionPool, "init connection pool failed.")
	Messages[ErrNotValidServerReadTimeout] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerReadTimeout, "server read timeout must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerWriteTimeout] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerWriteTimeout, "server write timeout must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerAddr] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerAddr, "server addr must be formatted as host:port, %s is not valid")
	Messages[ErrFieldNotExists] = config.NewErrMessage(DefaultMessageHeader, ErrFieldNotExists, "field %s does not exists. field name: %s")
	Messages[ErrGetRawData] = config.NewErrMessage(DefaultMessageHeader, ErrGetRawData, "get raw data from http body failed.\n%s")
	Messages[ErrUnmarshalRawData] = config.NewErrMessage(DefaultMessageHeader, ErrUnmarshalRawData, "unmarshal raw data failed.\n%s")
	Messages[ErrGenerateNewMapWithTag] = config.NewErrMessage(DefaultMessageHeader, ErrGenerateNewMapWithTag, "generate new map with tag %s failed.\n%s")
	Messages[ErrMarshalService] = config.NewErrMessage(DefaultMessageHeader, ErrMarshalService, "marshal service failed.\n%s")
	Messages[ErrTypeConversion] = config.NewErrMessage(DefaultMessageHeader, ErrTypeConversion, "type conversion failed.\n%s")
	// metadata
	Messages[ErrMetadataGetEnvAll] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetEnvAll, "metadata: get all environment failed.\n%s")
	Messages[ErrMetadataGetEnvByID] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetEnvByID, "metadata: get environment by id failed. id: %s\n%s")
	Messages[ErrMetadataAddEnv] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataAddEnv, "metadata: add new environment failed. env_name: %s\n%s")
	Messages[ErrMetadataUpdateEnv] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataUpdateEnv, "metadata: update environment failed. id: %s\n%s")

	Messages[ErrMetadataGetMySQLClusterAll] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetMySQLClusterAll, "metadata: get all mysql cluster failed.\n%s")
	Messages[ErrMetadataGetMySQLClusterByID] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetMySQLClusterByID, "metadata: get mysql cluster by id failed. id: %s\n%s")
	Messages[ErrMetadataAddMySQLCluster] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataAddMySQLCluster, "metadata: add new mysql cluster failed. cluster_name: %s middleware_cluster_id: %s monitor_system_id: %s owner_id: %s owner_group: %s env_id: %s\n%s")
	Messages[ErrMetadataUpdateMySQLCluster] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataUpdateMySQLCluster, "metadata: update mysql cluster failed. id: %s\n%s")

	Messages[ErrMetadataGetMySQLServerAll] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetMySQLServerAll, "metadata: get all mysql server failed.\n%s")
	Messages[ErrMetadataGetMySQLServerByID] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetMySQLServerByID, "metadata: get mysql server by id failed. id: %s\n%s")
	Messages[ErrMetadataAddMySQLServer] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataAddMySQLServer, "metadata: add new mysql server failed. cluster_id: %s server_name: %s host_ip: %s port_num: %s deployment_type: %s version: %s\n%s")
	Messages[ErrMetadataUpdateMySQLServer] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataUpdateMySQLServer, "metadata: update mysql server failed. id: %s\n%s")

	Messages[ErrMetadataGetMiddlewareServerAll] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetMiddlewareServerAll, "metadata: get all middleware server failed.\n%s")
	Messages[ErrMetadataGetMiddlewareServerByID] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetMiddlewareServerByID, "metadata: get middleware server by id failed. id: %s\n%s")
	Messages[ErrMetadataAddMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataAddMiddlewareServer, "metadata: add new middleware server failed. env_name: %s\n%s")
	Messages[ErrMetadataUpdateMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataUpdateMiddlewareServer, "metadata: update middleware server failed. id: %s\n%s")

	Messages[ErrMetadataGetUserAll] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetUserAll, "metadata: get all user failed.\n%s")
	Messages[ErrMetadataGetUserByID] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetUserByID, "metadata: get user by id failed. id: %s\n%s")
	Messages[ErrMetadataAddUser] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataAddUser, "metadata: add new user failed. user_name: %s\n%s")
	Messages[ErrMetadataUpdateUser] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataUpdateUser, "metadata: update user failed. id: %s\n%s")

	Messages[ErrMetadataGetMonitorSystemAll] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetMonitorSystemAll, "metadata: get all monitor systems failed.\n%s")
	Messages[ErrMetadataGetMonitorSystemByID] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataGetMonitorSystemByID, "metadata: get monitor system by id failed. id: %s\n%s")
	Messages[ErrMetadataAddMonitorSystem] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataAddMonitorSystem, "metadata: add new monitor system failed. system_name: %s\n%s")
	Messages[ErrMetadataUpdateMonitorSystem] = config.NewErrMessage(DefaultMessageHeader, ErrMetadataUpdateMonitorSystem, "metadata: update monitor system failed. id: %s\n%s")
}
