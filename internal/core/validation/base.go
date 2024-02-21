package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

var (
	ValidateStruct = validation.ValidateStruct
	Validate       = validation.Validate
	Field          = validation.Field
	By             = validation.By
	Key            = validation.Key
	Map            = validation.Map
	NewStringRule  = validation.NewStringRule

	Required   = validation.Required
	Each       = validation.Each
	Min        = validation.Min
	Max        = validation.Max
	Length     = validation.Length
	RuneLength = validation.RuneLength
	In         = validation.In
	When       = validation.When
	NotNil     = validation.NotNil
	Date       = validation.Date
	Match      = validation.Match
	MultipleOf = validation.MultipleOf
	NotIn      = validation.NotIn
)

type (
	FieldRules = validation.FieldRules
	RuleFunc   = validation.RuleFunc
	Rule       = validation.Rule
	Errors     = validation.Errors
)

func WithRuleBy[T any](fn func(data T) error) RuleFunc {
	return func(value interface{}) error {
		v, ok := value.(T)
		if !ok {
			return errors.New("invalid type")
		}
		return fn(v)
	}
}
