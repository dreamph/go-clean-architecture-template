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
	//Common : E101XXX
	ErrInternalServerError              = ErrCode("E101001", messages.E10001.Message)
	ErrVerificationCodeFailed           = ErrCode("E101002", messages.E00012.Message)
	ErrUserWalletAddressAlreadyAssigned = ErrCode("E101003", messages.E00013.FormatMessage("user's wallet address"))
	ErrInvalidWalletAddress             = ErrCode("E101004", messages.E00014.FormatMessage("wallet address"))
	ErrInvalidTemplate                  = ErrCode("E101005", messages.E00014.FormatMessage("template"))
	ErrDuplicatedName                   = ErrCode("E101006", messages.E00015.FormatMessage("name"))
	ErrInvalidData                      = ErrCode("E101007", messages.E00014.FormatMessage("data"))
	ErrCallExternalService              = ErrCode("E101008", messages.E00004.Message)
	ErrTokenInvalidOrExpired            = ErrCode("E101009", messages.E00016.FormatMessage("token"))
	ErrInvalidRequestData               = ErrCode("E101010", messages.E00011.Message)
	ErrDuplicatedEmail                  = ErrCode("E101011", messages.E00015.FormatMessage("email"))
	ErrCannotBeDeleted                  = ErrCode("E101012", messages.E00017.Message)
	ErrCannotBeProcessInThisState       = ErrCode("E101013", messages.E00018.Message)
	ErrDuplicatedIDCardNo               = ErrCode("E101014", messages.E00015.FormatMessage("id card no"))

	//Company : E105XXX
	ErrCompanyNotFound          = ErrCode("E105001", messages.E00001.FormatMessage("company"))
	ErrCompanyCodeAlreadyExists = ErrCode("E105005", messages.E00009.FormatMessage("This company code"))
)
