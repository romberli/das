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
	DebugMetadataGetAppAll    = 100101
	DebugMetadataGetAppByID   = 100102
	DebugMetadataGetAppByName = 100103
	DebugMetadataGetDBIDList  = 100104
	DebugMetadataAddApp       = 100105
	DebugMetadataUpdateApp    = 100106
	DebugMetadataDeleteApp    = 100107
	DebugMetadataAppAddDB     = 100108
	DebugMetadataAppDeleteDB  = 100109
	// info
	InfoMetadataGetAppAll    = 200101
	InfoMetadataGetAppByID   = 200102
	InfoMetadataGetAppByName = 200103
	InfoMetadataGetDBIDList  = 200104
	InfoMetadataAddApp       = 200105
	InfoMetadataUpdateApp    = 200106
	InfoMetadataDeleteApp    = 200107
	InfoMetadataAppAddDB     = 200108
	InfoMetadataAppDeleteDB  = 200109
	// error
	ErrMetadataGetAppAll    = 400101
	ErrMetadataGetAppByID   = 400102
	ErrMetadataGetAppByName = 400103
	ErrMetadataGetDBIDList  = 400104
	ErrMetadataAddApp       = 400105
	ErrMetadataUpdateApp    = 400106
	ErrMetadataDeleteApp    = 400107
	ErrMetadataAppAddDB     = 400108
	ErrMetadataAppDeleteDB  = 400109
)

func initDebugAppMessage() {
	message.Messages[DebugMetadataGetAppAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAppAll, "metadata: get all app completed. message: %s")
	message.Messages[DebugMetadataGetAppByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAppByID, "metadata: get app by id completed. message: %s")
	message.Messages[DebugMetadataGetAppByName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAppByName, "metadata: get app by name completed. message: %s")
	message.Messages[DebugMetadataGetDBIDList] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBIDList, "metadata: get db id list completed. message: %s")
	message.Messages[DebugMetadataAddApp] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddApp, "metadata: add new app completed. message: %s")
	message.Messages[DebugMetadataUpdateApp] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateApp, "metadata: update app completed. message: %s")
	message.Messages[DebugMetadataDeleteApp] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteApp, "metadata: delete app completed. message: %s")
	message.Messages[DebugMetadataAppAddDB] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAppAddDB, "metadata: add map of app and database completed. message: %s")
	message.Messages[DebugMetadataAppDeleteDB] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAppDeleteDB, "metadata: delete map of app and database completed. message: %s")
}

func initInfoAppMessage() {
	message.Messages[InfoMetadataGetAppAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAppAll, "metadata: get app all completed.")
	message.Messages[InfoMetadataGetAppByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAppByID, "metadata: get app by id completed. id: %d")
	message.Messages[InfoMetadataGetAppByName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAppByName, "metadata: get app by name completed. app_name: %s")
	message.Messages[InfoMetadataGetDBIDList] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBIDList, "metadata: get db id list completed. id: %d")
	message.Messages[InfoMetadataAddApp] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddApp, "metadata: add new app completed. app_name: %s")
	message.Messages[InfoMetadataUpdateApp] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateApp, "metadata: update app completed. id: %d")
	message.Messages[InfoMetadataDeleteApp] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteApp, "metadata: delete app completed. id: %d")
	message.Messages[InfoMetadataAppAddDB] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAppAddDB, "metadata: add map of app and database completed. app_id: %d, db_id: %d")
	message.Messages[InfoMetadataAppDeleteDB] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAppDeleteDB, "metadata: delete map of app and database completed. app_id: %d, db_id: %d")
}

func initErrorAppMessage() {
	message.Messages[ErrMetadataGetAppAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAppAll, "metadata: get all app failed.\n%s")
	message.Messages[ErrMetadataGetAppByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAppByID, "metadata: get app by id failed. id: %d\n%s")
	message.Messages[ErrMetadataGetAppByName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAppByName, "metadata: get app by name failed. app_name: %s\n%s")
	message.Messages[ErrMetadataGetDBIDList] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDBIDList, "metadata: get db list failed. id: %d\n%s")
	message.Messages[ErrMetadataAddApp] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddApp, "metadata: add new app failed. app_name: %s\n%s")
	message.Messages[ErrMetadataUpdateApp] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateApp, "metadata: update app failed. id: %d\n%s")
	message.Messages[ErrMetadataDeleteApp] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteApp, "metadata: delete app failed. id: %d\n%s")
	message.Messages[ErrMetadataAppAddDB] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAppAddDB, "metadata: add map of app and database failed. id: %d\n%s")
	message.Messages[ErrMetadataAppDeleteDB] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAppDeleteDB, "metadata: delete map of app and database failed. id: %d\n%s")
}
