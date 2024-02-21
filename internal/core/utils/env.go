package utils

import (
	"os"
	"regexp"
	"strings"
)

var envRegex = regexp.MustCompile(`\${.+?}`)

const defaultValueSep = ":"

// ParseEnvVar
// ParseEnvVar("${APP_NAME}")
// ParseEnvVar("${APP_ENV:dev}")
func ParseEnvVar(val string, getEnvFn func(string) string) string {
	return ParseEnvVarBy(val, defaultValueSep, getEnvFn)
}

func ParseEnvVarBy(val string, defaultValueSep string, getEnvFn func(string) string) string {
	if !strings.Contains(val, "${") {
		return val
	}

	// default use os.Getenv
	if getEnvFn == nil {
		getEnvFn = os.Getenv
	}

	var name, defaultValue string
	return envRegex.ReplaceAllStringFunc(val, func(eVar string) string {
		// eVar like "${NotExist|defaultValue}", first remove "${" and "}", then split it
		ss := strings.SplitN(eVar[2:len(eVar)-1], defaultValueSep, 2)

		// with default value. ${NotExist|defaultValue}
		if len(ss) == 2 {
			name, defaultValue = strings.TrimSpace(ss[0]), strings.TrimSpace(ss[1])
		} else {
			defaultValue = eVar // use raw value
			name = strings.TrimSpace(ss[0])
		}

		// get ENV value by name
		envValue := getEnvFn(name)
		if IsEmpty(envValue) {
			envValue = defaultValue
		}
		return envValue
	})
}
