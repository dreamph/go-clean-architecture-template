package multierr

import "go.uber.org/multierr"

func Append(currentErr error, newError error) error {
	return multierr.Append(currentErr, newError)
}

func Combine(errors ...error) error {
	return multierr.Combine(errors...)
}
