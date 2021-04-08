package message

import (
	"github.com/romberli/go-util/config"
)

const (
	// server
	InfoServerStart      = 200001
	InfoServerStop       = 200002
	InfoServerIsRunning  = 200003
	InfoServerNotRunning = 200004

	InfoMetadataGetEnvAll  = 200201
	InfoMetadataGetEnvByID = 200202
	InfoMetadataAddEnv     = 200203
	InfoMetadataUpdateEnv  = 200204

	InfoMetadataGetMiddlewareClusterAll  = 200501
	InfoMetadataGetMiddlewareClusterByID = 200502
	InfoMetadataAddMiddlewareCluster     = 200503
	InfoMetadataUpdateMiddlewareCluster  = 200504

	InfoMetadataGetMiddlewareServerAll  = 200601
	InfoMetadataGetMiddlewareServerByID = 200602
	InfoMetadataAddMiddlewareServer     = 200603
	InfoMetadataUpdateMiddlewareServer  = 200604

	InfoMetadataGetUserAll  = 200901
	InfoMetadataGetUserByID = 200902
	InfoMetadataAddUser     = 200903
	InfoMetadataUpdateUser  = 200904

	InfoMetadataGetMonitorSystemAll  = 200601
	InfoMetadataGetMonitorSystemByID = 200602
	InfoMetadataAddMonitorSystem     = 200603
	InfoMetadataUpdateMonitorSystem  = 200604
)

func initInfoMessage() {
	// server
	Messages[InfoServerStart] = config.NewErrMessage(DefaultMessageHeader, InfoServerStart, "das started successfully. pid: %d, pid file: %s")
	Messages[InfoServerStop] = config.NewErrMessage(DefaultMessageHeader, InfoServerStop, "das stopped successfully. pid: %d, pid file: %s")
	Messages[InfoServerIsRunning] = config.NewErrMessage(DefaultMessageHeader, InfoServerIsRunning, "das is running. pid: %d")
	Messages[InfoServerNotRunning] = config.NewErrMessage(DefaultMessageHeader, InfoServerNotRunning, "das is not running. pid: %d")

	Messages[InfoMetadataGetEnvAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetEnvAll, "metadata: get environment all completed")
	Messages[InfoMetadataGetEnvByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetEnvByID, "metadata: get environment by id completed. id: %s")
	Messages[InfoMetadataAddEnv] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddEnv, "metadata: add new environment completed. env_name: %s")
	Messages[InfoMetadataUpdateEnv] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateEnv, "metadata: update environment completed. id: %s")

	Messages[InfoMetadataGetMiddlewareClusterAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMiddlewareClusterAll, "metadata: get middleware cluster all completed. id: %s")
	Messages[InfoMetadataGetMiddlewareClusterByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id completed. id: %s")
	Messages[InfoMetadataAddMiddlewareCluster] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMiddlewareCluster, "metadata: add new middleware cluster completed. cluster_name: %s")
	Messages[InfoMetadataUpdateMiddlewareCluster] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster completed. id: %s")

	Messages[InfoMetadataGetMiddlewareServerAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMiddlewareServerAll, "metadata: get middleware server all completed. id: %s")
	Messages[InfoMetadataGetMiddlewareServerByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMiddlewareServerByID, "metadata: get middleware server by id completed. id: %s")
	Messages[InfoMetadataAddMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMiddlewareServer, "metadata: add new middleware server completed. server_name: %s")
	Messages[InfoMetadataUpdateMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMiddlewareServer, "metadata: update middleware server completed. id: %s")

	Messages[InfoMetadataGetUserAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetUserAll, "metadata: get user all completed")
	Messages[InfoMetadataGetUserByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetUserByID, "metadata: get user by id completed. id: %s")
	Messages[InfoMetadataAddUser] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddUser, "metadata: add new user completed. user_name: %s")
	Messages[InfoMetadataUpdateUser] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateUser, "metadata: update user completed. id: %s")

	Messages[InfoMetadataGetMonitorSystemAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMonitorSystemAll, "metadata: get monitor systems all completed")
	Messages[InfoMetadataGetMonitorSystemByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMonitorSystemByID, "metadata: get monitor system by id completed. id: %s")
	Messages[InfoMetadataAddMonitorSystem] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMonitorSystem, "metadata: add new monitor system completed. host_ip and port_num: %s")
	Messages[InfoMetadataUpdateMonitorSystem] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMonitorSystem, "metadata: update monitor system completed. id: %s")
}
