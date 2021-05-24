package healthcheck

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initDefaultEngineDebugMessage()
	initDefaultEngineInfoMessage()
	initDefaultEngineErrorMessage()
}

const (
	// debug

	// info

	// error
	ErrHealthcheckUpdateOperationStatus = 401001
)

func initDefaultEngineDebugMessage() {

}

func initDefaultEngineInfoMessage() {
}

func initDefaultEngineErrorMessage() {
	message.Messages[ErrHealthcheckUpdateOperationStatus] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckUpdateOperationStatus, "got error when updating operation status.\n%s")
}
