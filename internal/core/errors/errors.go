package errors

import (
	"errors"
	"strings"

	errs "github.com/pkg/errors"
)

type AppErrorData struct {
	Reference    string           `json:"reference"`
	ErrorDetails []AppErrorDetail `json:"errorDetails"`
}

type AppErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AppError struct {
	ErrType    error           `json:"-"`
	Err        error           `json:"-"`
	ErrCode    string          `json:"errCode"`
	ErrMessage string          `json:"errMessage"`
	ErrorData  *[]AppErrorData `json:"errorData"`
}

func (e *AppError) Error() string {
	return e.ErrType.Error() + ":" + e.Err.Error()
}

func (e *AppError) WithErrorData(appErrorList *[]AppErrorData) *AppError {
	e.ErrorData = appErrorList
	return e
}

func newError(errType error, err error, errCode string, errMessage string) *AppError {
	appError := &AppError{
		ErrType: errType,
	}
	if err != nil {
		appError.Err = errs.WithStack(err)
	}
	appError.ErrCode = errCode
	if len(errMessage) != 0 {
		appError.ErrMessage = errMessage
	}
	return appError
}

func ErrorValidation(message *ErrMessage) *AppError {
	return newError(ErrValidationFailed, errors.New(message.Message), message.Code, message.Message)
}

func ErrorUnauthorized(message *ErrMessage) *AppError {
	return newError(ErrUnauthorized, errors.New(message.Message), message.Code, message.Message)
}

func ErrorBusinessProcess(message *ErrMessage) *AppError {
	return newError(ErrBusinessProcessFailed, errors.New(message.Message), message.Code, message.Message)
}

func ErrorInternalServer(err error, message *ErrMessage) *AppError {
	return newError(ErrInternalServerError, err, message.Code, message.Message)
}

func ErrorNotFound(message *ErrMessage) *AppError {
	return newError(ErrNotFound, errors.New(message.Message), message.Code, message.Message)
}

func ErrorAuthInvalidRequest(message *ErrMessage) *AppError {
	return newError(ErrAuthInvalidRequest, errors.New(message.Message), message.Code, message.Message)
}

func ErrorCallExternalService(message *ErrMessage) *AppError {
	return newError(ErrCallExternalServiceFailed, errors.New(message.Message), message.Code, message.Message)
}

func WithStack(err error) error {
	return errs.WithStack(err)
}

func WithMessage(err error, message string) error {
	return errs.WithMessage(err, message)
}

func Wrap(err error, message string) error {
	return errs.Wrap(err, message)
}

func GetErrorCause(err error) error {
	var errCause error
	var appError *AppError
	ok := errors.As(err, &appError)
	if ok {
		errCause = appError.Err
	} else {
		errCause = errs.Cause(err)
	}
	return errCause
}

func GetErrorMessage(err error) string {
	var buf strings.Builder
	var appError *AppError
	ok := errors.As(err, &appError)
	if ok {
		buf.WriteString(appError.ErrCode)
		buf.WriteString(":")
		buf.WriteString(appError.ErrMessage)
	} else {
		buf.WriteString(err.Error())
	}
	return buf.String()
}

func GetApiClientError(err error) *ApiClientError {
	var appError *ApiClientError
	ok := errors.As(err, &appError)
	if ok {
		return appError
	}
	return nil
}
