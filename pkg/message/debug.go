package message

import (
	"github.com/romberli/go-util/config"
)

const (
	DebugMetadataGetEnvAll  = 100201
	DebugMetadataGetEnvByID = 100202
	DebugMetadataAddEnv     = 100203
	DebugMetadataUpdateEnv  = 100204

	DebugMetadataGetMYSQLClusterAll  = 100701
	DebugMetadataGetMYSQLClusterByID = 100702
	DebugMetadataAddMYSQLCluster     = 100703
	DebugMetadataUpdateMYSQLCluster  = 100704

	DebugMetadataGetMYSQLServerAll  = 100801
	DebugMetadataGetMYSQLServerByID = 100802
	DebugMetadataAddMYSQLServer     = 100803
	DebugMetadataUpdateMYSQLServer  = 100804
)

func initDebugMessage() {
	Messages[DebugMetadataGetEnvAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetEnvAll, "metadata: get all environments message: %s")
	Messages[DebugMetadataGetEnvByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetEnvByID, "metadata: get environment by id message: %s")
	Messages[DebugMetadataAddEnv] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddEnv, "metadata: add new environment message: %s")
	Messages[DebugMetadataUpdateEnv] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateEnv, "metadata: update environment message: %s")

	Messages[DebugMetadataGetMYSQLClusterAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMYSQLClusterAll, "metadata: get all mysql clusters message: %s")
	Messages[DebugMetadataGetMYSQLClusterByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMYSQLClusterByID, "metadata: get mysql cluster by id message: %s")
	Messages[DebugMetadataAddMYSQLCluster] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddMYSQLCluster, "metadata: add new mysql cluster message: %s")
	Messages[DebugMetadataUpdateMYSQLCluster] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateMYSQLCluster, "metadata: update mysql cluster message: %s")

	Messages[DebugMetadataGetMYSQLServerAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMYSQLServerAll, "metadata: get all mysql servers message: %s")
	Messages[DebugMetadataGetMYSQLServerByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMYSQLServerByID, "metadata: get mysql server by id message: %s")
	Messages[DebugMetadataAddMYSQLServer] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddMYSQLServer, "metadata: add new mysql server message: %s")
	Messages[DebugMetadataUpdateMYSQLServer] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateMYSQLServer, "metadata: update mysql server message: %s")
}
