package sqladvisor

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
	InfoSQLAdvisorGetFingerprint = 202001
	InfoSQLAdvisorGetSQLID       = 202002
	InfoSQLAdvisorAdvice         = 202003

	// error
	ErrSQLAdvisorAdvice = 402001
)

func initServiceDebugMessage() {

}

func initServiceInfoMessage() {
	message.Messages[InfoSQLAdvisorGetFingerprint] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoSQLAdvisorGetFingerprint,
		"sqladvisor: get fingerprint completed. sql text: %s, fingerprint: %s")
	message.Messages[InfoSQLAdvisorGetSQLID] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoSQLAdvisorGetSQLID,
		"sqladvisor: get sql id completed. sql text: %s, sql id: %s")
	message.Messages[InfoSQLAdvisorAdvice] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoSQLAdvisorAdvice,
		"sqladvisor: advice completed. db id: %d, sql text: %s, advice: %s")
}

func initServiceErrorMessage() {
	message.Messages[ErrSQLAdvisorAdvice] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrSQLAdvisorAdvice,
		"sqladvisor: advice failed. db id: %d, sql text: %s, error: %s")
}
