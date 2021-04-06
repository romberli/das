package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugAppMessage()
	initInfoAppMessage()
	initErrorAppMessage()
}

const (
	// debug
	DebugMetadataGetMiddlewareClusterAll  = 100401
	DebugMetadataGetMiddlewareClusterByID = 100402
	DebugMetadataAddMiddlewareCluster     = 100403
	DebugMetadataUpdateMiddlewareCluster  = 100404
	// info
	InfoMetadataGetMiddlewareClusterAll  = 200401
	InfoMetadataGetMiddlewareClusterByID = 200402
	InfoMetadataAddMiddlewareCluster     = 200403
	InfoMetadataUpdateMiddlewareCluster  = 200404
	// error
	ErrMetadataGetMiddlewareClusterAll  = 400401
	ErrMetadataGetMiddlewareClusterByID = 400402
	ErrMetadataAddMiddlewareCluster     = 400403
	ErrMetadataUpdateMiddlewareCluster  = 400404
)

func initDebugMiddlewareClusterMessage() {
	message.Messages[DebugMetadataGetMiddlewareClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClusterAll, "metadata: get all middleware cluster message: %s")
	message.Messages[DebugMetadataGetMiddlewareClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id message: %s")
	message.Messages[DebugMetadataAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddMiddlewareCluster, "metadata: add new middleware cluster message: %s")
	message.Messages[DebugMetadataUpdateMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster message: %s")

}

func initInfoMiddlewareClusteMessage() {
	message.Messages[InfoMetadataGetMiddlewareClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClusterAll, "metadata: get middleware cluster all completed. id: %s")
	message.Messages[InfoMetadataGetMiddlewareClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id completed. id: %s")
	message.Messages[InfoMetadataAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddMiddlewareCluster, "metadata: add new middleware cluster completed. cluster_name: %s")
	message.Messages[InfoMetadataUpdateMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster completed. id: %s")
}

func initErrorMiddlewareClusteMessage() {
	message.Messages[ErrMetadataGetMiddlewareClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClusterAll, "metadata: get all middleware cluster failed.\n%s")
	message.Messages[ErrMetadataGetMiddlewareClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id failed. id: %s\n%s")
	message.Messages[ErrMetadataAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddMiddlewareCluster, "metadata: add new middleware cluster failed. env_name: %s\n%s")
	message.Messages[ErrMetadataUpdateMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster failed. id: %s\n%s")
}
