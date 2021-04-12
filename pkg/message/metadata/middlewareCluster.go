package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugMiddlewareClusterMessage()
	initInfoMiddlewareClusteMessage()
	initErrorMiddlewareClusteMessage()
}

const (
	// debug
	DebugMetadataGetMiddlewareClusterAll    = 100401
	DebugMetadataGetMiddlewareClusterByEnv  = 100402
	DebugMetadataGetMiddlewareClusterByID   = 100403
	DebugMetadataGetMiddlewareClusterByName = 100404
	DebugMetadataGetMiddlewareServerIDList  = 100405
	DebugMetadataAddMiddlewareCluster       = 100406
	DebugMetadataUpdateMiddlewareCluster    = 100407
	DebugMetadataDeleteMiddlewareCluster    = 100408
	// info
	InfoMetadataGetMiddlewareClusterAll    = 200401
	InfoMetadataGetMiddlewareClusterByEnv  = 200402
	InfoMetadataGetMiddlewareClusterByID   = 200403
	InfoMetadataGetMiddlewareClusterByName = 200404
	InfoMetadataGetMiddlewareServerIDList  = 200405
	InfoMetadataAddMiddlewareCluster       = 200406
	InfoMetadataUpdateMiddlewareCluster    = 200407
	InfoMetadataDeleteMiddlewareCluster    = 200408
	// error
	ErrMetadataGetMiddlewareClusterAll    = 400401
	ErrMetadataGetMiddlewareClusterByEnv  = 400402
	ErrMetadataGetMiddlewareClusterByID   = 400403
	ErrMetadataGetMiddlewareClusterByName = 400404
	ErrMetadataGetMiddlewareServerIDList  = 400405
	ErrMetadataAddMiddlewareCluster       = 400406
	ErrMetadataUpdateMiddlewareCluster    = 400407
	ErrMetadataDeleteMiddlewareCluster    = 400408
)

func initDebugMiddlewareClusterMessage() {
	message.Messages[DebugMetadataGetMiddlewareClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClusterAll, "metadata: get all middleware clusters message: %s")
	message.Messages[DebugMetadataGetMiddlewareClusterByEnv] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClusterByEnv, "metadata: get middleware cluster by environment completed. message: %s")
	message.Messages[DebugMetadataGetMiddlewareClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id message: %s")
	message.Messages[DebugMetadataGetMiddlewareClusterByName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClusterByName, "metadata: get middleware cluster by name message: %s")
	message.Messages[DebugMetadataGetMiddlewareServerIDList] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareServerIDList, "metadata: get server id list completed. message: %s")
	message.Messages[DebugMetadataAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddMiddlewareCluster, "metadata: add new middleware cluster message: %s")
	message.Messages[DebugMetadataUpdateMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster message: %s")
	message.Messages[DebugMetadataDeleteMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteMiddlewareCluster, "metadata: delete middleware cluster completed. message: %s")
}

func initInfoMiddlewareClusteMessage() {
	message.Messages[InfoMetadataGetMiddlewareClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClusterAll, "metadata: get middleware clusters all completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareClusterByEnv] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClusterByEnv, "metadata: get middleware clusters by environment completed. env_id: %d")
	message.Messages[InfoMetadataGetMiddlewareClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareClusterByName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClusterByName, "metadata: get middleware cluster by name completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareServerIDList] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareServerIDList, "metadata: get middleware server id list completed. id: %d")
	message.Messages[InfoMetadataAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddMiddlewareCluster, "metadata: add new middleware cluster completed. cluster_name: %s")
	message.Messages[InfoMetadataUpdateMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster completed. id: %d")
	message.Messages[InfoMetadataDeleteMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteMiddlewareCluster, "metadata: delete middleware cluster completed. cluster_name: %s")
}

func initErrorMiddlewareClusteMessage() {
	message.Messages[ErrMetadataGetMiddlewareClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClusterAll, "metadata: get all middleware clusters failed.\n%s")
	message.Messages[ErrMetadataGetMiddlewareClusterByEnv] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClusterByEnv, "metadata: get middleware cluster by environment failed.\n%s")
	message.Messages[ErrMetadataGetMiddlewareClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id failed. id: %d\n%s")
	message.Messages[ErrMetadataGetMiddlewareClusterByName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClusterByName, "metadata: get middleware cluster by name failed. id: %d\n%s")
	message.Messages[ErrMetadataGetMiddlewareServerIDList] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareServerIDList, "metadata: get middleware server id list failed.\n%s")
	message.Messages[ErrMetadataAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddMiddlewareCluster, "metadata: add new middleware cluster failed. env_name: %s\n%s")
	message.Messages[ErrMetadataUpdateMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster failed. id: %d\n%s")
	message.Messages[ErrMetadataDeleteMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteMiddlewareCluster, "metadata: delete middleware cluster failed. cluster_name: %s\n%s")
}
