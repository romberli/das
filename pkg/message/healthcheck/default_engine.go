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
	ErrHealthcheckUpdateOperationStatus       = 401001
	ErrDefaultEngineConfigContent             = 401002
	ErrItemWeightItemInvalid                  = 401003
	ErrLowWatermarkItemInvalid                = 401004
	ErrHighWatermarkItemInvalid               = 401005
	ErrUnitItemInvalid                        = 401006
	ErrScoreDeductionPerUnitHighItemInvalid   = 401007
	ErrMaxScoreDeductionHighItemInvalid       = 401008
	ErrScoreDeductionPerUnitMediumItemInvalid = 401009
	ErrMaxScoreDeductionMediumItemInvalid     = 401010
	ErrItemWeightPercentInvalid               = 401011
	ErrDefaultEngineConfigFormatInValid       = 401012
)

func initDefaultEngineDebugMessage() {

}

func initDefaultEngineInfoMessage() {
}

func initDefaultEngineErrorMessage() {
	message.Messages[ErrHealthcheckUpdateOperationStatus] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckUpdateOperationStatus, "got error when updating operation status\n%s")
	message.Messages[ErrDefaultEngineConfigContent] = config.NewErrMessage(message.DefaultMessageHeader, ErrDefaultEngineConfigContent, "default engine config doesn't have content")
	message.Messages[ErrItemWeightItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrItemWeightItemInvalid, "item weight of %s must be in [1, 100], %d is not valid")
	message.Messages[ErrLowWatermarkItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrLowWatermarkItemInvalid, "low watermark of %s must be higher than 0, %f is not valid")
	message.Messages[ErrHighWatermarkItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHighWatermarkItemInvalid, "high watermark of %s  must be higher than low watermark, %f is not valid")
	message.Messages[ErrUnitItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrUnitItemInvalid, "unit of %s must be higher than 0, %f is not valid")
	message.Messages[ErrScoreDeductionPerUnitHighItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrScoreDeductionPerUnitHighItemInvalid, "score deduction per unit high of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrMaxScoreDeductionHighItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrMaxScoreDeductionHighItemInvalid, "max score deduction high of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrScoreDeductionPerUnitMediumItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrScoreDeductionPerUnitMediumItemInvalid, "score deduction per unit medium of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrMaxScoreDeductionMediumItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrMaxScoreDeductionMediumItemInvalid, "max score deduction medium of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrItemWeightPercentInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrItemWeightPercentInvalid, "all items weight count is not 100")
	message.Messages[ErrDefaultEngineConfigFormatInValid] = config.NewErrMessage(message.DefaultMessageHeader, ErrDefaultEngineConfigFormatInValid, "default engine config format is invalid")
}
