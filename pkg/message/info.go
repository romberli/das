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

	InfoMetadataGetMySQLClusterAll  = 200701
	InfoMetadataGetMySQLClusterByID = 200702
	InfoMetadataAddMySQLCluster     = 200703
	InfoMetadataUpdateMySQLCluster  = 200704

	InfoMetadataGetMySQLServerAll  = 200801
	InfoMetadataGetMySQLServerByID = 200802
	InfoMetadataAddMySQLServer     = 200803
	InfoMetadataUpdateMySQLServer  = 200804

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

	Messages[InfoMetadataGetMySQLClusterAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMySQLClusterAll, "metadata: get mysql cluster all completed")
	Messages[InfoMetadataGetMySQLClusterByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMySQLClusterByID, "metadata: get mysql cluster by id completed. id: %s")
	Messages[InfoMetadataAddMySQLCluster] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMySQLCluster, "metadata: add new mysql cluster completed. env_name: %s")
	Messages[InfoMetadataUpdateMySQLCluster] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMySQLCluster, "metadata: update mysql cluster completed. id: %s")

	Messages[InfoMetadataGetMySQLServerAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMySQLServerAll, "metadata: get mysql server all completed")
	Messages[InfoMetadataGetMySQLServerByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMySQLServerByID, "metadata: get mysql server by id completed. id: %s")
	Messages[InfoMetadataAddMySQLServer] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMySQLServer, "metadata: add new mysql server completed. env_name: %s")
	Messages[InfoMetadataUpdateMySQLServer] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMySQLServer, "metadata: update mysql server completed. id: %s")

	Messages[InfoMetadataGetUserAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetUserAll, "metadata: get user all completed")
	Messages[InfoMetadataGetUserByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetUserByID, "metadata: get user by id completed. id: %s")
	Messages[InfoMetadataAddUser] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddUser, "metadata: add new user completed. user_name: %s")
	Messages[InfoMetadataUpdateUser] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateUser, "metadata: update user completed. id: %s")

	Messages[InfoMetadataGetMonitorSystemAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMonitorSystemAll, "metadata: get monitor systems all completed")
	Messages[InfoMetadataGetMonitorSystemByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMonitorSystemByID, "metadata: get monitor system by id completed. id: %s")
	Messages[InfoMetadataAddMonitorSystem] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMonitorSystem, "metadata: add new monitor system completed. host_ip and port_num: %s")
	Messages[InfoMetadataUpdateMonitorSystem] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMonitorSystem, "metadata: update monitor system completed. id: %s")
}
