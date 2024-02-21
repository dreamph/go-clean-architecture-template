package validation

import (
	"backend/internal/core/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	LengthEq = func(len int) validation.Rule {
		return validation.RuneLength(len, len)
	}

	LengthMax = func(len int) validation.Rule {
		return validation.RuneLength(0, len)
	}

	LengthMin = func(len int) validation.Rule {
		return validation.RuneLength(len, 0)
	}

	InList = func(equalsIgnoreCase bool, values ...string) validation.Rule {
		return validation.By(validateInRule(equalsIgnoreCase, values...))
	}

	PageLimit = validation.By(WithRuleBy[*models.PageLimit](ValidatePageLimit))
)
