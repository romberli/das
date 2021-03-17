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

	InfoMetadataGetMiddlewareClusterAll  = 200501
	InfoMetadataGetMiddlewareClusterByID = 200502
	InfoMetadataAddMiddlewareCluster     = 200503
	InfoMetadataUpdateMiddlewareCluster  = 200504

	InfoMetadataGetMiddlewareServerAll  = 200601
	InfoMetadataGetMiddlewareServerByID = 200602
	InfoMetadataAddMiddlewareServer     = 200603
	InfoMetadataUpdateMiddlewareServer  = 200604

	InfoMetadataGetAppSystemAll  = 200301
	InfoMetadataGetAppSystemByID = 200302
	InfoMetadataAddAppSystem     = 200303
	InfoMetadataUpdateAppSystem  = 200304

	InfoMetadataGetUserAll  = 200901
	InfoMetadataGetUserByID = 200902
	InfoMetadataAddUser     = 200903
	InfoMetadataUpdateUser  = 200904

	InfoMetadataGetDbAll  = 200401
	InfoMetadataGetDbByID = 200402
	InfoMetadataAddDb     = 200403
	InfoMetadataUpdateDb  = 200404

	InfoMetadataGetMSAll  = 200601
	InfoMetadataGetMSByID = 200602
	InfoMetadataAddMS     = 200603
	InfoMetadataUpdateMS  = 200604
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

	Messages[InfoMetadataGetMiddlewareClusterAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMiddlewareClusterAll, "metadata: get middleware cluster all completed. id: %s")
	Messages[InfoMetadataGetMiddlewareClusterByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id completed. id: %s")
	Messages[InfoMetadataAddMiddlewareCluster] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMiddlewareCluster, "metadata: add new middleware cluster completed. cluster_name: %s")
	Messages[InfoMetadataUpdateMiddlewareCluster] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster completed. id: %s")

	Messages[InfoMetadataGetMiddlewareServerAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMiddlewareServerAll, "metadata: get middleware server all completed. id: %s")
	Messages[InfoMetadataGetMiddlewareServerByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMiddlewareServerByID, "metadata: get middleware server by id completed. id: %s")
	Messages[InfoMetadataAddMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMiddlewareServer, "metadata: add new middleware server completed. server_name: %s")
	Messages[InfoMetadataUpdateMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMiddlewareServer, "metadata: update middleware server completed. id: %s")

	Messages[InfoMetadataGetAppSystemAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetAppSystemAll, "metadata: get appsystem all completed")
	Messages[InfoMetadataGetAppSystemByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetAppSystemByID, "metadata: get appsystem by id completed. id: %s")
	Messages[InfoMetadataAddAppSystem] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddAppSystem, "metadata: add new appsystem completed. system_name: %s")
	Messages[InfoMetadataUpdateAppSystem] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateAppSystem, "metadata: update appsystem completed. id: %s")

	Messages[InfoMetadataGetUserAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetUserAll, "metadata: get user all completed")
	Messages[InfoMetadataGetUserByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetUserByID, "metadata: get user by id completed. id: %s")
	Messages[InfoMetadataAddUser] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddUser, "metadata: add new user completed. env_name: %s")
	Messages[InfoMetadataUpdateUser] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateUser, "metadata: update user completed. id: %s")

	Messages[InfoMetadataGetDbAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetDbAll, "metadata: get database all completed")
	Messages[InfoMetadataGetDbByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetDbByID, "metadata: get database by id completed. id: %s")
	Messages[InfoMetadataAddDb] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddDb, "metadata: add new database completed. db_name and owner_id and env_id: %s")
	Messages[InfoMetadataUpdateDb] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateDb, "metadata: update database completed. id: %s")

	Messages[InfoMetadataGetMSAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMSAll, "metadata: get monitor systems all completed")
	Messages[InfoMetadataGetMSByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMSByID, "metadata: get monitor system by id completed. id: %s")
	Messages[InfoMetadataAddMS] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMS, "metadata: add new monitor system completed. system_name and system_type and host_ip and port_num and port_num_slow and base_url: %s")
	Messages[InfoMetadataUpdateMS] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMS, "metadata: update monitor system completed. id: %s")
}
