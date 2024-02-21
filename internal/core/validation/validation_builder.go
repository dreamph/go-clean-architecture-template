package validation

import (
	"backend/internal/core/validation/rules"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

/*
type Builder interface {
	Validate() error
	Add(fields ...*FieldRules) Builder
	AddByCondition(condition bool, fields ...*FieldRules) Builder
	Required(fields ...*FieldRules) Builder
}
*/

func NewStructValidationBuilder(structPtr interface{}, fields ...*FieldRules) *Builder {
	return &Builder{structPtr: structPtr, fields: fields}
}

type Builder struct {
	structPtr      any
	fields         []*FieldRules
	requiredFields []*FieldRules
}

/*
func (v *Builder) AddFieldRulesByCondition(condition bool, fields ...*FieldRules) *Builder {
	if condition {
		v.AddFieldRules(fields...)
	}
	return v
}*/

func (v *Builder) AddFieldRules(fields ...*FieldRules) *Builder {
	if len(fields) > 0 {
		v.fields = append(v.fields, fields...)
	}
	return v
}

func (v *Builder) AddRequiredFieldRules(fields ...*FieldRules) *Builder {
	if len(fields) > 0 {
		v.requiredFields = append(v.requiredFields, fields...)
	}
	return v
}

func (v *Builder) Validate() error {
	if len(v.requiredFields) > 0 {
		err := ValidateStruct(v.structPtr,
			v.requiredFields...,
		)
		if err != nil {
			return err
		}
	}

	return ValidateStruct(v.structPtr,
		v.fields...,
	)
}

func StructField[T any](fieldPtr any, rule func(value T) error) *FieldRules {
	return validation.Field(fieldPtr, rules.ByRule(rule))
}

func ArrayField[T any](fieldPtr any, rule func(value T, i int) error) *FieldRules {
	return validation.Field(fieldPtr, rules.ByRuleArray(rule))
}
