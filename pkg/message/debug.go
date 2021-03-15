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

	DebugMetadataGetMiddlewareClusterAll  = 100501
	DebugMetadataGetMiddlewareClusterByID = 100502
	DebugMetadataAddMiddlewareCluster     = 100503
	DebugMetadataUpdateMiddlewareCluster  = 100504

	DebugMetadataGetMiddlewareServerAll  = 100601
	DebugMetadataGetMiddlewareServerByID = 100602
	DebugMetadataAddMiddlewareServer     = 100603
	DebugMetadataUpdateMiddlewareServer  = 100604

	DebugMetadataGetAppSystemAll  = 100301
	DebugMetadataGetAppSystemByID = 100302
	DebugMetadataAddAppSystem     = 100303
	DebugMetadataUpdateAppSystem  = 100304

	DebugMetadataGetUserAll  = 100901
	DebugMetadataGetUserByID = 100902
	DebugMetadataAddUser     = 100903
	DebugMetadataUpdateUser  = 100904

	DebugMetadataGetDbAll  = 100401
	DebugMetadataGetDbByID = 100402
	DebugMetadataAddDb     = 100403
	DebugMetadataUpdateDb  = 100404

	DebugMetadataGetMSAll  = 100601
	DebugMetadataGetMSByID = 100602
	DebugMetadataAddMS     = 100603
	DebugMetadataUpdateMS  = 100604
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

	Messages[DebugMetadataGetMiddlewareClusterAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMiddlewareClusterAll, "metadata: get all middleware cluster message: %s")
	Messages[DebugMetadataGetMiddlewareClusterByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id message: %s")
	Messages[DebugMetadataAddMiddlewareCluster] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddMiddlewareCluster, "metadata: add new middleware cluster message: %s")
	Messages[DebugMetadataUpdateMiddlewareCluster] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster message: %s")

	Messages[DebugMetadataGetMiddlewareServerAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMiddlewareServerAll, "metadata: get all middleware server message: %s")
	Messages[DebugMetadataGetMiddlewareServerByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMiddlewareServerByID, "metadata: get middleware server by id message: %s")
	Messages[DebugMetadataAddMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddMiddlewareServer, "metadata: add new middleware server message: %s")
	Messages[DebugMetadataUpdateMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateMiddlewareServer, "metadata: update middleware server message: %s")

	Messages[DebugMetadataGetAppSystemAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetAppSystemAll, "metadata: get all appsystem message: %s")
	Messages[DebugMetadataGetAppSystemByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetAppSystemByID, "metadata: get appsystem by id message: %s")
	Messages[DebugMetadataAddAppSystem] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddAppSystem, "metadata: add new appsystem message: %s")
	Messages[DebugMetadataUpdateAppSystem] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateAppSystem, "metadata: update appsystem message: %s")

	Messages[DebugMetadataGetUserAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetUserAll, "metadata: get all user message: %s")
	Messages[DebugMetadataGetUserByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetUserByID, "metadata: get user by id message: %s")
	Messages[DebugMetadataAddUser] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddUser, "metadata: add new user message: %s")
	Messages[DebugMetadataUpdateUser] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateUser, "metadata: update user message: %s")

	Messages[DebugMetadataGetDbAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetDbAll, "metadata: get all databases message: %s")
	Messages[DebugMetadataGetDbByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetDbByID, "metadata: get database by id message: %s")
	Messages[DebugMetadataAddDb] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddDb, "metadata: add new database message: %s")
	Messages[DebugMetadataUpdateDb] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateDb, "metadata: update database message: %s")

	Messages[DebugMetadataGetMSAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMSAll, "metadata: get all monitor systems message: %s")
	Messages[DebugMetadataGetMSByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMSByID, "metadata: get monitor system by id message: %s")
	Messages[DebugMetadataAddMS] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddMS, "metadata: add new monitor system message: %s")
	Messages[DebugMetadataUpdateMS] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateMS, "metadata: update monitor system message: %s")
}
