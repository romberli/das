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

	InfoMetadataGetMYSQLClusterAll  = 200701
	InfoMetadataGetMYSQLClusterByID = 200702
	InfoMetadataAddMYSQLCluster     = 200703
	InfoMetadataUpdateMYSQLCluster  = 200704

	InfoMetadataGetMYSQLServerAll  = 200801
	InfoMetadataGetMYSQLServerByID = 200802
	InfoMetadataAddMYSQLServer     = 200803
	InfoMetadataUpdateMYSQLServer  = 200804
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

	Messages[InfoMetadataGetMYSQLClusterAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMYSQLClusterAll, "metadata: get mysql cluster all completed")
	Messages[InfoMetadataGetMYSQLClusterByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMYSQLClusterByID, "metadata: get mysql cluster by id completed. id: %s")
	Messages[InfoMetadataAddMYSQLCluster] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMYSQLCluster, "metadata: add new mysql cluster completed. env_name: %s")
	Messages[InfoMetadataUpdateMYSQLCluster] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMYSQLCluster, "metadata: update mysql cluster completed. id: %s")

	Messages[InfoMetadataGetMYSQLServerAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMYSQLServerAll, "metadata: get mysql server all completed")
	Messages[InfoMetadataGetMYSQLServerByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMYSQLServerByID, "metadata: get mysql server by id completed. id: %s")
	Messages[InfoMetadataAddMYSQLServer] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMYSQLServer, "metadata: add new mysql server completed. env_name: %s")
	Messages[InfoMetadataUpdateMYSQLServer] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMYSQLServer, "metadata: update mysql server completed. id: %s")
}
