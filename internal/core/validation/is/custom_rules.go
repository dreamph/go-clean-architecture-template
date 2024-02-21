package is

import (
	"backend/internal/core/utils"
	"backend/internal/core/validation"
)

var (
	EnglishLang = validation.By(validation.WithRuleBy[string](utils.ValidateEnglishLang))
)
