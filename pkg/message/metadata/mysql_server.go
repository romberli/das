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
	// debug
	DebugMetadataGetMySQLServerAll         = 100801
	DebugMetadataGetMySQLServerByClusterID = 100802
	DebugMetadataGetMySQLServerByID        = 100803
	DebugMetadataGetMySQLServerByHostInfo  = 100804
	DebugMetadataAddMySQLServer            = 100805
	DebugMetadataUpdateMySQLServer         = 100806
	DebugMetadataDeleteMySQLServer         = 100807
)
const (
	// info
	InfoMetadataGetMySQLServerAll         = 200801
	InfoMetadataGetMySQLServerByClusterID = 200802
	InfoMetadataGetMySQLServerByID        = 200803
	InfoMetadataGetMySQLServerByHostInfo  = 200804
	InfoMetadataAddMySQLServer            = 200805
	InfoMetadataUpdateMySQLServer         = 200806
	InfoMetadataDeleteMySQLServer         = 200807
)
const (
	// error
	ErrMetadataGetMySQLServerAll         = 400801
	ErrMetadataGetMySQLServerByClusterID = 400802
	ErrMetadataGetMySQLServerByID        = 400803
	ErrMetadataGetMySQLServerByHostInfo  = 400804
	ErrMetadataAddMySQLServer            = 400805
	ErrMetadataUpdateMySQLServer         = 400806
	ErrMetadataDeleteMySQLServer         = 400807
)

func initDebugMySQLServerMessage() {
	message.Messages[DebugMetadataGetMySQLServerAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLServerAll,
		"metadata: get all mysql servers message: %s")
	message.Messages[DebugMetadataGetMySQLServerByClusterID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLServerByClusterID,
		"metadata: get mysql server by cluster id message: %s")
	message.Messages[DebugMetadataGetMySQLServerByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLServerByID,
		"metadata: get mysql server by id message: %s")
	message.Messages[DebugMetadataGetMySQLServerByHostInfo] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLServerByHostInfo,
		"metadata: get mysql server by host info message: %s")
	message.Messages[DebugMetadataAddMySQLServer] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataAddMySQLServer,
		"metadata: add new mysql server message: %s")
	message.Messages[DebugMetadataUpdateMySQLServer] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataUpdateMySQLServer,
		"metadata: update mysql server message: %s")
	message.Messages[DebugMetadataDeleteMySQLServer] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataDeleteMySQLServer,
		"metadata: delete mysql server message: %s")
}

func initInfoMySQLServerMessage() {
	message.Messages[InfoMetadataGetMySQLServerAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLServerAll,
		"metadata: get mysql server all completed")
	message.Messages[InfoMetadataGetMySQLServerByClusterID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLServerByClusterID,
		"metadata: get mysql server by cluster id completed. cluster_id: %d")
	message.Messages[InfoMetadataGetMySQLServerByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLServerByID,
		"metadata: get mysql server by id completed. id: %d")
	message.Messages[InfoMetadataGetMySQLServerByHostInfo] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLServerByHostInfo,
		"metadata: get mysql server by host info completed. id: %d")
	message.Messages[InfoMetadataAddMySQLServer] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataAddMySQLServer,
		"metadata: add new mysql server completed. server_name: %s, cluster_id: %d, host_ip: %s, port_num: %d, deployment_type: %d")
	message.Messages[InfoMetadataUpdateMySQLServer] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataUpdateMySQLServer,
		"metadata: update mysql server completed. id: %d")
	message.Messages[InfoMetadataDeleteMySQLServer] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataDeleteMySQLServer,
		"metadata: delete mysql server completed. id: %d")
}

func initErrorMySQLServerMessage() {
	message.Messages[ErrMetadataGetMySQLServerAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLServerAll,
		"metadata: get all mysql server failed.\n%s")
	message.Messages[ErrMetadataGetMySQLServerByClusterID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLServerByClusterID,
		"metadata: get mysql server by cluster id failed. cluster_id: %d\n%s")
	message.Messages[ErrMetadataGetMySQLServerByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLServerByID,
		"metadata: get mysql server by id failed. id: %d\n%s")
	message.Messages[ErrMetadataGetMySQLServerByHostInfo] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLServerByHostInfo,
		"metadata: get mysql server by host info failed. id: %d\n%s")
	message.Messages[ErrMetadataAddMySQLServer] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataAddMySQLServer,
		"metadata: add new mysql server failed. server_name: server_name: %s, cluster_id: %d, host_ip: %s, port_num: %d, deployment_type: %d\n%s")
	message.Messages[ErrMetadataUpdateMySQLServer] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataUpdateMySQLServer,
		"metadata: update mysql server failed. server_name: %s\n%s")
	message.Messages[ErrMetadataDeleteMySQLServer] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataDeleteMySQLServer,
		"metadata: delete mysql server failed. id: %s\n%s")
}
