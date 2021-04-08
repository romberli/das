package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugMonitorSystemMessage()
	initInfoMonitorSystemMessage()
	initErrorMonitorSystemMessage()
}

const (
	// debug
	DebugMetadataGetMonitorSystemAll        = 100601
	DebugMetadataGetMonitorSystemByEnv      = 100602
	DebugMetadataGetMonitorSystemByID       = 100603
	DebugMetadataGetMonitorSystemByHostInfo = 100604
	DebugMetadataAddMonitorSystem           = 100605
	DebugMetadataUpdateMonitorSystem        = 100606
	DebugMetadataDeleteMonitorSystem        = 100607
	// info
	InfoMetadataGetMonitorSystemAll        = 200601
	InfoMetadataGetMonitorSystemByEnv      = 200602
	InfoMetadataGetMonitorSystemByID       = 200603
	InfoMetadataGetMonitorSystemByHostInfo = 200604
	InfoMetadataAddMonitorSystem           = 200605
	InfoMetadataUpdateMonitorSystem        = 200606
	InfoMetadataDeleteMonitorSystem        = 200607
	// error
	ErrMetadataGetMonitorSystemAll        = 400601
	ErrMetadataGetMonitorSystemByEnv      = 400602
	ErrMetadataGetMonitorSystemByID       = 400603
	ErrMetadataGetMonitorSystemByHostInfo = 400604
	ErrMetadataAddMonitorSystem           = 400605
	ErrMetadataUpdateMonitorSystem        = 400606
	ErrMetadataDeleteMonitorSystem        = 400607
)

func initDebugMonitorSystemMessage() {
	message.Messages[DebugMetadataGetMonitorSystemAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMonitorSystemAll, "metadata: get all monitor systems completed. message: %s")
	message.Messages[DebugMetadataGetMonitorSystemByEnv] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMonitorSystemByEnv, "metadata: get monitor systems by environment completed. message: %s")
	message.Messages[DebugMetadataGetMonitorSystemByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMonitorSystemByID, "metadata: get monitor system by id completed. message: %s")
	message.Messages[DebugMetadataGetMonitorSystemByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMonitorSystemByHostInfo, "metadata: get monitor system by host info completed. message: %s")
	message.Messages[DebugMetadataAddMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddMonitorSystem, "metadata: add new monitor system completed. message: %s")
	message.Messages[DebugMetadataUpdateMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateMonitorSystem, "metadata: update monitor system completed. message: %s")
	message.Messages[DebugMetadataDeleteMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteMonitorSystem, "metadata: delete monitor system completed. message: %s")
}

func initInfoMonitorSystemMessage() {
	message.Messages[InfoMetadataGetMonitorSystemAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMonitorSystemAll, "metadata: get all monitor systems completed")
	message.Messages[InfoMetadataGetMonitorSystemByEnv] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMonitorSystemByEnv, "metadata: get monitor systems by environment completed. env_id: %d")
	message.Messages[InfoMetadataGetMonitorSystemByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMonitorSystemByID, "metadata: get monitor system by id completed. id: %d")
	message.Messages[InfoMetadataGetMonitorSystemByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMonitorSystemByHostInfo, "metadata: get monitor system by host info completed. host_ip: %s, port_num: %d")
	message.Messages[InfoMetadataAddMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddMonitorSystem, "metadata: add new monitor system completed. system_name: %s, system_type: %d, host_ip: %s, port_num: %d, port_num_slow: %d, base_url: %s, env_id: %d")
	message.Messages[InfoMetadataUpdateMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateMonitorSystem, "metadata: update monitor system completed. id: %d")
	message.Messages[InfoMetadataDeleteMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteMonitorSystem, "metadata: delete monitor system completed. id: %d")
}

func initErrorMonitorSystemMessage() {
	message.Messages[ErrMetadataGetMonitorSystemAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMonitorSystemAll, "metadata: get all monitor systems failed.\n%s")
	message.Messages[ErrMetadataGetMonitorSystemByEnv] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMonitorSystemByEnv, "metadata: get monitor systems by environment failed. env_id: %d\n%s")
	message.Messages[ErrMetadataGetMonitorSystemByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMonitorSystemByID, "metadata: get monitor system by id failed. id: %d\n%s")
	message.Messages[ErrMetadataGetMonitorSystemByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMonitorSystemByHostInfo, "metadata: get monitor system by host info failed. host_ip: %s, port_num: %d\n%s")
	message.Messages[ErrMetadataAddMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddMonitorSystem, "metadata: add new monitor system failed. system_name: %s, system_type: %d, host_ip: %s, port_num: %d, port_num_slow: %d, base_url: %s, env_id: %d\n%s")
	message.Messages[ErrMetadataUpdateMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateMonitorSystem, "metadata: update monitor system failed. id: %d\n%s")
	message.Messages[ErrMetadataDeleteMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteMonitorSystem, "metadata: delete monitor system failed. id: %d\n%s")
}
