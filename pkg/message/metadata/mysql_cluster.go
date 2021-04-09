package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugMySQLCLusterMessage()
	initInfoMySQLCLusterMessage()
	initErrorMySQLCLusterMessage()
}

const (
	// debug
	DebugMetadataGetMySQLClusterAll    = 100701
	DebugMetadataGetMySQLClusterByEnv  = 100702
	DebugMetadataGetMySQLClusterByID   = 100703
	DebugMetadataGetMySQLClusterByName = 100704
	DebugMetadataGetMySQLServerIDList  = 100705
	DebugMetadataAddMySQLCluster       = 100706
	DebugMetadataUpdateMySQLCluster    = 100707
	DebugMetadataDeleteMySQLCluster    = 100708
	// debug
	InfoMetadataGetMySQLClusterAll    = 200701
	InfoMetadataGetMySQLClusterByEnv  = 200702
	InfoMetadataGetMySQLClusterByID   = 200703
	InfoMetadataGetMySQLClusterByName = 200704
	InfoMetadataGetMySQLServerIDList  = 200705
	InfoMetadataAddMySQLCluster       = 200706
	InfoMetadataUpdateMySQLCluster    = 200707
	InfoMetadataDeleteMySQLCluster    = 200708
	// error
	ErrMetadataGetMySQLClusterAll    = 400701
	ErrMetadataGetMySQLClusterByEnv  = 400702
	ErrMetadataGetMySQLClusterByID   = 400703
	ErrMetadataGetMySQLClusterByName = 400704
	ErrMetadataGetMySQLServerIDList  = 400705
	ErrMetadataAddMySQLCluster       = 400706
	ErrMetadataUpdateMySQLCluster    = 400707
	ErrMetadataDeleteMySQLCluster    = 400708
)

func initDebugMySQLCLusterMessage() {
	message.Messages[DebugMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterAll,
		"metadata: get all mysql clusters message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByEnv,
		"metadata: get mysql cluster by env message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByID,
		"metadata: get mysql cluster by id message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByName,
		"metadata: get mysql cluster by name message: %s")
	message.Messages[DebugMetadataGetMySQLServerIDList] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLServerIDList,
		"metadata: get mysql server id from mysql cluster message: %s")
	message.Messages[DebugMetadataAddMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataAddMySQLCluster,
		"metadata: add new mysql cluster message: %s")
	message.Messages[DebugMetadataUpdateMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataUpdateMySQLCluster,
		"metadata: update mysql cluster message: %s")
	message.Messages[DebugMetadataDeleteMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataDeleteMySQLCluster,
		"metadata: delete mysql cluster message: %s")
}

func initInfoMySQLCLusterMessage() {
	message.Messages[InfoMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterAll,
		"metadata: get mysql cluster all completed")
	message.Messages[InfoMetadataGetMySQLClusterByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterByEnv,
		"metadata: get mysql cluster by env completed. env_id: %s")
	message.Messages[InfoMetadataGetMySQLClusterByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterByID,
		"metadata: get mysql cluster by id completed. id: %s")
	message.Messages[InfoMetadataGetMySQLClusterByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterByName,
		"metadata: get mysql cluster by name completed. cluster_name: %s")
	message.Messages[InfoMetadataGetMySQLServerIDList] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLServerIDList,
		"metadata: get mysql server id from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataAddMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataAddMySQLCluster,
		"metadata: add new mysql cluster completed. cluster_name: %s, env_id: %s")
	message.Messages[InfoMetadataUpdateMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataUpdateMySQLCluster,
		"metadata: update mysql cluster completed. id: %s")
	message.Messages[InfoMetadataDeleteMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataDeleteMySQLCluster,
		"metadata: delete mysql cluster completed. id: %s")
}

func initErrorMySQLCLusterMessage() {
	message.Messages[ErrMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterAll,
		"metadata: get all mysql cluster failed.\n%s")
	message.Messages[ErrMetadataGetMySQLClusterByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterByEnv,
		"metadata: get mysql cluster by env failed. env_id: %d\n%s")
	message.Messages[ErrMetadataGetMySQLClusterByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterByID,
		"metadata: get mysql cluster by id failed. id: %d\n%s")
	message.Messages[ErrMetadataGetMySQLClusterByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterByName,
		"metadata: get mysql cluster by name failed. cluster_name: %s\n%s")
	message.Messages[ErrMetadataGetMySQLServerIDList] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLServerIDList,
		"metadata: get mysql server id from mysql cluster failed. id: %d\n%s")
	message.Messages[ErrMetadataAddMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataAddMySQLCluster,
		"metadata: add new mysql cluster failed. cluster_name: %s, env_id: %d\n%s")
	message.Messages[ErrMetadataUpdateMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataUpdateMySQLCluster,
		"metadata: update mysql cluster failed. id: %d\n%s")
	message.Messages[ErrMetadataDeleteMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataDeleteMySQLCluster,
		"metadata: delete mysql cluster failed. id: %d\n%s")
}
