package metadata

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initDebugDBMessage()
	initInfoDBMessage()
	initErrorDBMessage()
}

const (
	//debug
	DebugMetadataGetEnvAll     = 100301
	DebugMetadataGetEnvByID    = 100302
	DebugMetadataAddEnv        = 100303
	DebugMetadataUpdateEnv     = 100304
	DebugMetadataGetEnvByName  = 100305
	DebugMetadataDeleteEnvByID = 100306
	//error
	ErrMetadataGetEnvAll     = 400301
	ErrMetadataGetEnvByID    = 400302
	ErrMetadataAddEnv        = 400303
	ErrMetadataUpdateEnv     = 400304
	ErrMetadataGetEnvByName  = 400305
	ErrMetadataDeleteEnvByID = 400306
	//info
	InfoMetadataGetEnvAll     = 200301
	InfoMetadataGetEnvByID    = 200302
	InfoMetadataAddEnv        = 200303
	InfoMetadataUpdateEnv     = 200304
	InfoMetadataGetEnvByName  = 200305
	InfoMetadataDeleteEnvByID = 200306
)

func initDebugEnvMessage() {
	//debug
	message.Messages[DebugMetadataGetEnvAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEnvAll, "metadata: get all environments message: %s")
	message.Messages[DebugMetadataGetEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEnvByID, "metadata: get environment by id message: %s")
	message.Messages[DebugMetadataAddEnv] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddEnv, "metadata: add new environment message: %s")
	message.Messages[DebugMetadataUpdateEnv] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateEnv, "metadata: update environment message: %s")
	message.Messages[DebugMetadataGetEnvByName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEnvByName, "metadata: get environment by name message: %s")
	message.Messages[DebugMetadataDeleteEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteEnvByID, "metadata: delete environment by ID message: %s")
}

func initErrorEnvMessage() {
	message.Messages[ErrMetadataGetEnvAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEnvAll, "metadata: get all environment failed.\n%s")
	message.Messages[ErrMetadataGetEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEnvByID, "metadata: get environment by id failed. id: %s\n%s")
	message.Messages[ErrMetadataAddEnv] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddEnv, "metadata: add new environment failed. env_name: %s\n%s")
	message.Messages[ErrMetadataUpdateEnv] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateEnv, "metadata: update environment failed. id: %s\n%s")
	message.Messages[ErrMetadataGetEnvByName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEnvByName, "metadata: get environment by name failed. id: %s\n%s")
	message.Messages[ErrMetadataDeleteEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteEnvByID, "metadata: delete environment by ID failed. id: %s\n%s")
}

func initInfoEnvMessage() {
	message.Messages[InfoMetadataGetEnvAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEnvAll, "metadata: get environment all completed")
	message.Messages[InfoMetadataGetEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEnvByID, "metadata: get environment by id completed. id: %s")
	message.Messages[InfoMetadataAddEnv] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddEnv, "metadata: add new environment completed. env_name: %s")
	message.Messages[InfoMetadataUpdateEnv] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateEnv, "metadata: update environment completed. id: %s")
	message.Messages[InfoMetadataGetEnvByName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEnvByName, "metadata: get environment by name completed. id: %s")
	message.Messages[InfoMetadataDeleteEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteEnvByID, "metadata: delete environment by ID completed. id: %s")
}
