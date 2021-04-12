package metadata

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initDebugMiddlewareServerMessage()
	initInfoMiddlewareServerMessage()
	initErrorMiddlewareServerMessage()
}

const (
	// debug
	DebugMetadataGetMiddlewareServerAll        = 100501
	DebugMetadataGetMiddlewareSeverByClusterID = 100502
	DebugMetadataGetMiddlewareServerByID       = 100503
	DebugMetadataGetMiddlewareServerByHostInfo = 100504
	DebugMetadataAddMiddlewareServer           = 100505
	DebugMetadataUpdateMiddlewareServer        = 100506
	DebugMetadataDeleteMiddlewareServer        = 100507

	// info
	InfoMetadataGetMiddlewareServerAll        = 200501
	InfoMetadataGetMiddlewareSeverByClusterID = 200502
	InfoMetadataGetMiddlewareServerByID       = 200503
	InfoMetadataGetMiddlewareServerByHostInfo = 200504
	InfoMetadataAddMiddlewareServer           = 200505
	InfoMetadataUpdateMiddlewareServer        = 200506
	InfoMetadataDeleteMiddlewareServer        = 200507
	// error
	ErrMetadataGetMiddlewareServerAll        = 400501
	ErrMetadataGetMiddlewareSeverByClusterID = 400502
	ErrMetadataGetMiddlewareServerByID       = 400503
	ErrMetadataGetMiddlewareServerByHostInfo = 400504
	ErrMetadataAddMiddlewareServer           = 400505
	ErrMetadataUpdateMiddlewareServer        = 400506
	ErrMetadataDeleteMiddlewareServer        = 400507
)

func initDebugMiddlewareServerMessage() {
	message.Messages[DebugMetadataGetMiddlewareServerAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareServerAll, "metadata: get all middleware server message: %s")
	message.Messages[DebugMetadataGetMiddlewareSeverByClusterID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareSeverByClusterID, "metadata: get middleware cluster by cluster completed. message: %s")
	message.Messages[DebugMetadataGetMiddlewareServerByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareServerByID, "metadata: get middleware server by id message: %s")
	message.Messages[DebugMetadataGetMiddlewareServerByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareServerByHostInfo, "metadata: get middleware cluster by host info message: %s")
	message.Messages[DebugMetadataAddMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddMiddlewareServer, "metadata: add new middleware server message: %s")
	message.Messages[DebugMetadataUpdateMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateMiddlewareServer, "metadata: update middleware server message: %s")
	message.Messages[DebugMetadataDeleteMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteMiddlewareServer, "metadata: delete middleware server completed. message: %s")
}

func initInfoMiddlewareServerMessage() {
	message.Messages[InfoMetadataGetMiddlewareServerAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareServerAll, "metadata: get middleware server all completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareSeverByClusterID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareSeverByClusterID, "metadata: get middleware clusters by cluster completed. cluster_id: %d")
	message.Messages[InfoMetadataGetMiddlewareServerByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareServerByID, "metadata: get middleware server by id completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareServerByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareServerByHostInfo, "metadata: get middleware cluster by host info completed. host-ip: %s")
	message.Messages[InfoMetadataAddMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddMiddlewareServer, "metadata: add new middleware server completed. server_name: %s")
	message.Messages[InfoMetadataUpdateMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateMiddlewareServer, "metadata: update middleware server completed. id: %d")
	message.Messages[InfoMetadataDeleteMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteMiddlewareServer, "metadata: delete middleware server completed. server_name: %s")
}

func initErrorMiddlewareServerMessage() {
	message.Messages[ErrMetadataGetMiddlewareServerAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareServerAll, "metadata: get all middleware server failed.\n%s")
	message.Messages[ErrMetadataGetMiddlewareSeverByClusterID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareSeverByClusterID, "metadata: get middleware cluster by cluster failed.\n%s")
	message.Messages[ErrMetadataGetMiddlewareServerByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareServerByID, "metadata: get middleware server by id failed. id: %d\n%s")
	message.Messages[ErrMetadataGetMiddlewareServerByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareServerByHostInfo, "metadata: get middleware cluster by host info failed. host-ip: %s\n%s")
	message.Messages[ErrMetadataAddMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddMiddlewareServer, "metadata: add new middleware server failed. server_name: %s\n%s")
	message.Messages[ErrMetadataUpdateMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateMiddlewareServer, "metadata: update middleware server failed. id: %d\n%s")
	message.Messages[ErrMetadataDeleteMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteMiddlewareServer, "metadata: delete middleware server failed. server_name: %s\n%s")
}
