package message

import (
	"github.com/romberli/go-util/config"
)

const (
	// server
	InfoServerStart      = 200001
	InfoServerStop       = 200002
	InfoServerIsRunning  = 200003
	InfoServerNotRunning = 200004

	InfoMetadataGetEnvAll  = 200201
	InfoMetadataGetEnvByID = 200202
	InfoMetadataAddEnv     = 200203
	InfoMetadataUpdateEnv  = 200204

	InfoMetadataGetDbAll  = 200401
	InfoMetadataGetDbByID = 200402
	InfoMetadataAddDb    = 200403
	InfoMetadataUpdateDb  = 200404

	InfoMetadataGetMSAll  = 200601
	InfoMetadataGetMSByID = 200602
	InfoMetadataAddMS    = 200603
	InfoMetadataUpdateMS  = 200604
)

func initInfoMessage() {
	// server
	Messages[InfoServerStart] = config.NewErrMessage(DefaultMessageHeader, InfoServerStart, "das started successfully. pid: %d, pid file: %s")
	Messages[InfoServerStop] = config.NewErrMessage(DefaultMessageHeader, InfoServerStop, "das stopped successfully. pid: %d, pid file: %s")
	Messages[InfoServerIsRunning] = config.NewErrMessage(DefaultMessageHeader, InfoServerIsRunning, "das is running. pid: %d")
	Messages[InfoServerNotRunning] = config.NewErrMessage(DefaultMessageHeader, InfoServerNotRunning, "das is not running. pid: %d")

	Messages[InfoMetadataGetEnvAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetEnvAll, "metadata: get environment all completed")
	Messages[InfoMetadataGetEnvByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetEnvByID, "metadata: get environment by id completed. id: %s")
	Messages[InfoMetadataAddEnv] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddEnv, "metadata: add new environment completed. env_name: %s")
	Messages[InfoMetadataUpdateEnv] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateEnv, "metadata: update environment completed. id: %s")

	Messages[InfoMetadataGetDbAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetDbAll, "metadata: get database all completed")
	Messages[InfoMetadataGetDbByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetDbByID, "metadata: get database by id completed. id: %s")
	Messages[InfoMetadataAddDb] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddDb, "metadata: add new database completed. db_name and owner_id and env_id: %s")
	Messages[InfoMetadataUpdateDb] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateDb, "metadata: update database completed. id: %s")

	Messages[InfoMetadataGetMSAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMSAll, "metadata: get monitor systems all completed")
	Messages[InfoMetadataGetMSByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMSByID, "metadata: get monitor system by id completed. id: %s")
	Messages[InfoMetadataAddMS] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMS, "metadata: add new monitor system completed. system_name and system_type and host_ip and port_num and port_num_slow and base_url: %s")
	Messages[InfoMetadataUpdateMS] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMS, "metadata: update monitor system completed. id: %s")
}
