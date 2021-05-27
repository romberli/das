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
	DebugHealthcheckGetResultByOperationID = 101012
	DebugHealthcheckCheck                  = 101013
	DebugHealthcheckCheckByHostInfo        = 101014
	DebugHealthcheckReviewAccurate         = 101015
	// info
	InfoHealthcheckGetResultByOperationID = 201012
	InfoHealthcheckCheck                  = 201013
	InfoHealthcheckCheckByHostInfo        = 201014
	InfoHealthcheckReviewAccurate         = 201015
	// error
	ErrHealthcheckDefaultEngineRun       = 401011
	ErrHealthcheckGetResultByOperationID = 401012
	ErrHealthcheckCheck                  = 401013
	ErrHealthcheckCheckByHostInfo        = 401014
	ErrHealthcheckReviewAccurate         = 401015
	ErrHealthcheckCloseConnection        = 401016
)

func initServiceDebugMessage() {
	message.Messages[DebugHealthcheckGetResultByOperationID] = config.NewErrMessage(
		message.DefaultMessageHeader, DebugHealthcheckGetResultByOperationID,
		"healthcheck: get result by operation id message: %s")
	message.Messages[DebugHealthcheckCheck] = config.NewErrMessage(
		message.DefaultMessageHeader, DebugHealthcheckCheck,
		"healthcheck: check message: %s")
	message.Messages[DebugHealthcheckCheckByHostInfo] = config.NewErrMessage(
		message.DefaultMessageHeader, DebugHealthcheckCheckByHostInfo,
		"healthcheck: check by host info message: %s")
	message.Messages[DebugHealthcheckReviewAccurate] = config.NewErrMessage(
		message.DefaultMessageHeader, DebugHealthcheckReviewAccurate,
		"healthcheck: review accurate message: %s")
}

func initServiceInfoMessage() {
	message.Messages[InfoHealthcheckGetResultByOperationID] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoHealthcheckGetResultByOperationID,
		"healthcheck: get result by operation id compeleted. operation_id: %d")
	message.Messages[InfoHealthcheckCheck] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoHealthcheckCheck,
		"healthcheck: check compeleted. %s")
	message.Messages[InfoHealthcheckCheckByHostInfo] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoHealthcheckCheckByHostInfo,
		"healthcheck: check by host info compeleted. %s")
	message.Messages[InfoHealthcheckReviewAccurate] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoHealthcheckReviewAccurate,
		"healthcheck: review accurate compeleted. %s")
}

func initServiceErrorMessage() {
	message.Messages[ErrHealthcheckDefaultEngineRun] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckDefaultEngineRun,
		"default engine run failed.\n%s")
	message.Messages[ErrHealthcheckGetResultByOperationID] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckGetResultByOperationID,
		"healthcheck: get result by operation id failed. operation_id: %d\n%s")
	message.Messages[ErrHealthcheckCheck] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckCheck,
		"healthcheck: check failed.  %s")
	message.Messages[ErrHealthcheckCheckByHostInfo] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckCheckByHostInfo,
		"healthcheck: check by host info failed. host info %s")
	message.Messages[ErrHealthcheckCloseConnection] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckCloseConnection,
		"healthcheck: close middleware connection failed.\n%s")

}
