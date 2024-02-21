package rules

import (
	"reflect"
)

type ValidateByRule[T any] struct {
	rule func(value T) error
}

func ByRule[T any](rule func(value T) error) *ValidateByRule[T] {
	return &ValidateByRule[T]{rule: rule}
}

func (l *ValidateByRule[T]) Validate(value interface{}) error {
	fv := reflect.ValueOf(value)
	if fv.Kind() == reflect.Ptr {
		value = fv.Elem().Interface()
	}
	return l.rule(value.(T))
}
