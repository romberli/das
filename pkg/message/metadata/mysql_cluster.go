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
	DebugMetadataGetMySQLClusterAll  = 100701
	DebugMetadataGetMySQLClusterByID = 100702
	DebugMetadataAddMySQLCluster     = 100703
	DebugMetadataUpdateMySQLCluster  = 100704
	// debug
	InfoMetadataGetMySQLClusterAll  = 200701
	InfoMetadataGetMySQLClusterByID = 200702
	InfoMetadataAddMySQLCluster     = 200703
	InfoMetadataUpdateMySQLCluster  = 200704
	// error
	ErrMetadataGetMySQLClusterAll  = 400701
	ErrMetadataGetMySQLClusterByID = 400702
	ErrMetadataAddMySQLCluster     = 400703
	ErrMetadataUpdateMySQLCluster  = 400704
)

func initDebugMySQLCLusterMessage() {
	message.Messages[DebugMetadataGetMySQLClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMySQLClusterAll, "metadata: get all mysql clusters message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMySQLClusterByID, "metadata: get mysql cluster by id message: %s")
	message.Messages[DebugMetadataAddMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddMySQLCluster, "metadata: add new mysql cluster message: %s")
	message.Messages[DebugMetadataUpdateMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateMySQLCluster, "metadata: update mysql cluster message: %s")
}

func initInfoMySQLCLusterMessage() {
	message.Messages[InfoMetadataGetMySQLClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMySQLClusterAll, "metadata: get mysql cluster all completed")
	message.Messages[InfoMetadataGetMySQLClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMySQLClusterByID, "metadata: get mysql cluster by id completed. id: %s")
	message.Messages[InfoMetadataAddMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddMySQLCluster, "metadata: add new mysql cluster completed. env_name: %s")
	message.Messages[InfoMetadataUpdateMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateMySQLCluster, "metadata: update mysql cluster completed. id: %s")
}

func initErrorMySQLCLusterMessage() {
	message.Messages[ErrMetadataGetMySQLClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMySQLClusterAll, "metadata: get all mysql cluster failed.\n%s")
	message.Messages[ErrMetadataGetMySQLClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMySQLClusterByID, "metadata: get mysql cluster by id failed. id: %s\n%s")
	message.Messages[ErrMetadataAddMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddMySQLCluster, "metadata: add new mysql cluster failed. cluster_name: %s middleware_cluster_id: %s monitor_system_id: %s owner_id: %s owner_group: %s env_id: %s\n%s")
	message.Messages[ErrMetadataUpdateMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateMySQLCluster, "metadata: update mysql cluster failed. id: %s\n%s")

}
