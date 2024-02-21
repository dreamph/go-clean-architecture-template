package errors

import (
	"github.com/pkg/errors"
)

var (
	ErrValidationFailed          = errors.New("ValidationFailed")
	ErrBusinessProcessFailed     = errors.New("BusinessProcessFailed")
	ErrNotFound                  = errors.New("RecordNotFound")
	ErrUnauthorized              = errors.New("Unauthorized")
	ErrPermissionDenied          = errors.New("PermissionDenied")
	ErrInternalServerError       = errors.New("InternalServerError")
	ErrCallExternalServiceFailed = errors.New("CallExternalServiceFailed")
	ErrAuthInvalidRequest        = errors.New("AuthInvalidRequest")
)
