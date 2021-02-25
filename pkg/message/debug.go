package message

import (
	"github.com/romberli/go-util/config"
)

const (
	DebugMetadataGetEnvAll  = 100201
	DebugMetadataGetEnvByID = 100202
	DebugMetadataAddEnv     = 100203
	DebugMetadataUpdateEnv  = 100204

	DebugMetadataGetAppSystemAll  = 100301
	DebugMetadataGetAppSystemByID = 100302
	DebugMetadataAddAppSystem     = 100303
	DebugMetadataUpdateAppSystem  = 100304

	DebugMetadataGetUserAll  = 100901
	DebugMetadataGetUserByID = 100902
	DebugMetadataAddUser     = 100903
	DebugMetadataUpdateUser  = 100904
)

func initDebugMessage() {
	Messages[DebugMetadataGetEnvAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetEnvAll, "metadata: get all environments message: %s")
	Messages[DebugMetadataGetEnvByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetEnvByID, "metadata: get environment by id message: %s")
	Messages[DebugMetadataAddEnv] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddEnv, "metadata: add new environment message: %s")
	Messages[DebugMetadataUpdateEnv] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateEnv, "metadata: update environment message: %s")

	Messages[DebugMetadataGetAppSystemAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetAppSystemAll, "metadata: get all appsystem message: %s")
	Messages[DebugMetadataGetAppSystemByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetAppSystemByID, "metadata: get appsystem by id message: %s")
	Messages[DebugMetadataAddAppSystem] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddAppSystem, "metadata: add new appsystem message: %s")
	Messages[DebugMetadataUpdateAppSystem] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateAppSystem, "metadata: update appsystem message: %s")

	Messages[DebugMetadataGetUserAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetUserAll, "metadata: get all user message: %s")
	Messages[DebugMetadataGetUserByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetUserByID, "metadata: get user by id message: %s")
	Messages[DebugMetadataAddUser] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddUser, "metadata: add new user message: %s")
	Messages[DebugMetadataUpdateUser] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateUser, "metadata: update user message: %s")
}
