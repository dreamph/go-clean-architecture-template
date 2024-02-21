package nullvalue

import (
	"time"

	"github.com/guregu/null"
)

func BlankString() null.String {
	return null.StringFrom("")
}

func StringFrom(val string) null.String {
	return null.NewString(val, val != "")
}

func StringFromPtr(val *string) null.String {
	return null.StringFromPtr(val)
}

func IntFrom(val int64) null.Int {
	return null.NewInt(val, val != 0)
}

func IntFromPtr(val *int64) null.Int {
	return null.IntFromPtr(val)
}

func BoolFrom(val bool) null.Bool {
	return null.NewBool(val, true)
}

func BoolFromPtr(val *bool) null.Bool {
	return null.BoolFromPtr(val)
}

func FloatFrom(val float64) null.Float {
	return null.NewFloat(val, val != 0)
}

func FloatFromPtr(val *float64) null.Float {
	return null.FloatFromPtr(val)
}

func TimeFrom(val time.Time) null.Time {
	return null.NewTime(val, !val.IsZero())
}

func TimeFromPtr(val *time.Time) null.Time {
	return null.TimeFromPtr(val)
}

func IsNullString(val null.String) bool {
	return !val.Valid
}

func IsNullInt(val null.Int) bool {
	return !val.Valid
}

func IsNullBool(val null.Bool) bool {
	return !val.Valid
}

func IsNullFloat(val null.Float) bool {
	return !val.Valid
}

func IsNullTime(val null.Time) bool {
	return !val.Valid
}
