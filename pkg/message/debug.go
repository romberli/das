package message

import (
	"github.com/romberli/go-util/config"
)

const (
	DebugMetadataGetEnvAll  = 100201
	DebugMetadataGetEnvByID = 100202
	DebugMetadataAddEnv     = 100203
	DebugMetadataUpdateEnv  = 100204

	DebugMetadataGetMiddlewareClusterAll  = 100501
	DebugMetadataGetMiddlewareClusterByID = 100502
	DebugMetadataAddMiddlewareCluster     = 100503
	DebugMetadataUpdateMiddlewareCluster  = 100504

	DebugMetadataGetMiddlewareServerAll  = 100601
	DebugMetadataGetMiddlewareServerByID = 100602
	DebugMetadataAddMiddlewareServer     = 100603
	DebugMetadataUpdateMiddlewareServer  = 100604
)

func initDebugMessage() {
	Messages[DebugMetadataGetEnvAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetEnvAll, "metadata: get all environments message: %s")
	Messages[DebugMetadataGetEnvByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetEnvByID, "metadata: get environment by id message: %s")
	Messages[DebugMetadataAddEnv] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddEnv, "metadata: add new environment message: %s")
	Messages[DebugMetadataUpdateEnv] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateEnv, "metadata: update environment message: %s")

	Messages[DebugMetadataGetMiddlewareClusterAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMiddlewareClusterAll, "metadata: get all middleware cluster message: %s")
	Messages[DebugMetadataGetMiddlewareClusterByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id message: %s")
	Messages[DebugMetadataAddMiddlewareCluster] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddMiddlewareCluster, "metadata: add new middleware cluster message: %s")
	Messages[DebugMetadataUpdateMiddlewareCluster] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster message: %s")

	Messages[DebugMetadataGetMiddlewareServerAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMiddlewareServerAll, "metadata: get all middleware server message: %s")
	Messages[DebugMetadataGetMiddlewareServerByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMiddlewareServerByID, "metadata: get middleware server by id message: %s")
	Messages[DebugMetadataAddMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddMiddlewareServer, "metadata: add new middleware server message: %s")
	Messages[DebugMetadataUpdateMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateMiddlewareServer, "metadata: update middleware server message: %s")

}
