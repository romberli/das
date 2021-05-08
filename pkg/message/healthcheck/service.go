package healthcheck

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initServiceDebugMessage()
	initServiceInfoMessage()
	initServiceErrorMessage()
}

const (
	// debug

	// info

	// error
	ErrHealthcheckDefaultEngineRun = 401011
)

func initServiceDebugMessage() {

}

func initServiceInfoMessage() {
}

func initServiceErrorMessage() {
	message.Messages[ErrHealthcheckDefaultEngineRun] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckDefaultEngineRun,
		"default engine run failed.\n%s")
}
