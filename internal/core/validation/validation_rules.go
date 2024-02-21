package validation

import (
	"backend/internal/core/models"
	"backend/internal/core/utils"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type DynamicRule struct {
	RuleKey         string
	ValidationRules map[string][]validation.Rule
	FormatData      func(value string) string
}

func (d DynamicRule) Get() []validation.Rule {
	return d.ValidationRules[d.RuleKey]
}

// ValidateRule
/*
	err := validation.ValidateStruct(request,
		validation.Field(&request.Limit, validation.Required, validation.WithRule(func() error {
			//return validation.ValidatePageLimit(request.Limit)
			return validation.ValidateStruct(request,
				validation.Field(&request.Limit.PageNumber, validation.Required),
				validation.Field(&request.Limit.PageSize, validation.Required),
			)
		})),
	)
	if err != nil {
		return err
	}
	return nil
*/
type validateRule struct {
	rule func() error
}

func WithRule(rule func() error) validation.Rule {
	return &validateRule{rule: rule}
}

func (l *validateRule) Validate(value interface{}) error {
	return l.rule()
}

/*
var ValidatePageLimitRule = NewPageLimitRule()

type limitRule struct {
}

func NewPageLimitRule() validation.Rule {
	return &limitRule{}
}

func (l *limitRule) Validate(value interface{}) error {
	limit, ok := value.(*models.PageLimit)
	if !ok {
		return errors.New("limit: is not *models.PageLimit type")
	}
	return ValidatePageLimit(limit)
}

*/

func ValidatePageLimit(limit *models.PageLimit) error {
	err := validation.ValidateStruct(limit,
		validation.Field(&limit.PageNumber, validation.Required, validation.Min(1)),
		validation.Field(&limit.PageSize, validation.Required, validation.Min(1)),
	)
	if err != nil {
		return err
	}
	return nil
}

func validateInRule(equalsIgnoreCase bool, values ...string) RuleFunc {
	return func(value interface{}) error {
		valStr, ok := value.(string)
		if !ok {
			return errors.New("must be a string")
		}
		if utils.IsEmpty(valStr) {
			return nil
		}
		for _, m := range values {
			if equalsIgnoreCase {
				if strings.EqualFold(m, valStr) {
					return nil
				}
			} else {
				if m == valStr {
					return nil
				}
			}
		}
		return errors.New("invalid value")
	}
}
