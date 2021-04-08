package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugMySQLServerMessage()
	initInfoMySQLServerMessage()
	initErrorMySQLServerMessage()
}

const (
	// FIXME: 此处建议使用iota防止错误码改动的重新编序
	// debug
	DebugMetadataGetMySQLServerAll = iota + 100801
	DebugMetadataGetMySQLServerByClusterID
	DebugMetadataGetMySQLServerByID
	DebugMetadataAddMySQLServer
	DebugMetadataUpdateMySQLServer
)
const (
	// info
	InfoMetadataGetMySQLServerAll = iota + 200801
	InfoMetadataGetMySQLServerByClusterID
	InfoMetadataGetMySQLServerByID
	InfoMetadataAddMySQLServer
	InfoMetadataUpdateMySQLServer
)
const (
	// error
	ErrMetadataGetMySQLServerAll = iota + 400801
	ErrMetadataGetMySQLServerByClusterID
	ErrMetadataGetMySQLServerByID
	ErrMetadataAddMySQLServer
	ErrMetadataUpdateMySQLServer
)

func initDebugMySQLServerMessage() {
	message.Messages[DebugMetadataGetMySQLServerAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMySQLServerAll, "metadata: get all mysql servers message: %s")
	message.Messages[DebugMetadataGetMySQLServerByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMySQLServerByID, "metadata: get mysql server by id message: %s")
	message.Messages[DebugMetadataAddMySQLServer] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddMySQLServer, "metadata: add new mysql server message: %s")
	message.Messages[DebugMetadataUpdateMySQLServer] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateMySQLServer, "metadata: update mysql server message: %s")
}

func initInfoMySQLServerMessage() {
	message.Messages[InfoMetadataGetMySQLServerAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMySQLServerAll, "metadata: get mysql server all completed")
	message.Messages[InfoMetadataGetMySQLServerByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMySQLServerByID, "metadata: get mysql server by id completed. id: %s")
	message.Messages[InfoMetadataAddMySQLServer] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddMySQLServer, "metadata: add new mysql server completed. env_name: %s")
	message.Messages[InfoMetadataUpdateMySQLServer] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateMySQLServer, "metadata: update mysql server completed. id: %s")
}

func initErrorMySQLServerMessage() {
	message.Messages[ErrMetadataGetMySQLServerAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMySQLServerAll, "metadata: get all mysql server failed.\n%s")
	message.Messages[ErrMetadataGetMySQLServerByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMySQLServerByID, "metadata: get mysql server by id failed. id: %s\n%s")
	message.Messages[ErrMetadataAddMySQLServer] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddMySQLServer, "metadata: add new mysql server failed. cluster_id: %s server_name: %s host_ip: %s port_num: %s deployment_type: %s version: %s\n%s")
	message.Messages[ErrMetadataUpdateMySQLServer] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateMySQLServer, "metadata: update mysql server failed. id: %s\n%s")
}
