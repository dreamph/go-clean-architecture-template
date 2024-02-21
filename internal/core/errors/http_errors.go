package errors

import (
	errs "github.com/pkg/errors"
)

type ApiClientError struct {
	Err        error
	StatusCode int
}

func (e *ApiClientError) Error() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

func NewApiClientError(err error, statusCode int) error {
	appError := &ApiClientError{
		Err:        errs.WithStack(err),
		StatusCode: statusCode,
	}
	return appError
}
