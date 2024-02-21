package rules

import (
	"reflect"

	"github.com/pkg/errors"

	"fmt"
)

type ValidateByRuleArray[T any] struct {
	rule func(value T, i int) error
}

func ByRuleArray[T any](rule func(value T, i int) error) *ValidateByRuleArray[T] {
	return &ValidateByRuleArray[T]{rule: rule}
}

func (l *ValidateByRuleArray[T]) Validate(value interface{}) error {
	fv := reflect.ValueOf(value)
	if fv.Kind() == reflect.Ptr {
		value = fv.Elem().Interface()
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(value)
		for i := 0; i < s.Len(); i++ {
			val := fv.Index(i).Interface()
			err := l.rule(val.(T), i)
			if err != nil {
				return errors.New(err.Error() + fmt.Sprintf(" (index:%v)", i))
			}
		}
	default:
	}
	return nil
}
