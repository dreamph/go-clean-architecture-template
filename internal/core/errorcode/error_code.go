package errorcode

import (
	"backend/internal/constants/messages"
	"backend/internal/core/errors"
	"log"
)

var errorCodeMap map[string]string

func ErrorCodes() map[string]string {
	return errorCodeMap
}

func ErrCode(errorCode string, message string) *errors.ErrMessage {
	if errorCodeMap == nil {
		errorCodeMap = map[string]string{}
	}
	if _, exists := errorCodeMap[errorCode]; exists {
		log.Fatalf("%s : ErrCode Duplicate!", errorCode)
	} else {
		errorCodeMap[errorCode] = message
	}
	return errors.NewErrMessage(errorCode, message)
}

var (
	ErrInternalErrorDefault = ErrCode("999", "Something went wrong.")

	//AuthFlow : EC11XXX
	ErrAuthFlowCodeNotFound       = ErrCode("EC11001", messages.E00001.FormatMessage("authFlowCode"))
	ErrAuthFlowTokenRefIsRequired = ErrCode("EC11002", messages.E00008.FormatMessage("tokenRef"))
	//
)
