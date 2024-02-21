package utils

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

func IsEmpty(value string) bool {
	return value == ""
}

func IsNotEmpty(value string) bool {
	return !IsEmpty(value)
}

func IsEmptyList[T any](value *[]T) bool {
	return value == nil || len(*value) == 0
}

func IsNotEmptyList[T any](value *[]T) bool {
	return !IsEmptyList(value)
}

func ValidateUUIDList(value []string) error {
	for _, val := range value {
		if !IsUUID(val) {
			return errors.New("Invalid UUID:" + val)
		}
	}
	return nil
}

func ValidateUUID(value string) error {
	if !IsUUID(value) {
		return errors.New("Invalid UUID:" + value)
	}
	return nil
}

func IsUUID(value string) bool {
	return govalidator.IsUUID(value)
}

func IsContainsKey[K comparable](dataMap map[K]string, key K) bool {
	_, ok := dataMap[key]
	return ok
}

func IsEmptyOrWhiteSpace(value string) bool {
	return IsEmpty(value) || IsEmpty(strings.TrimSpace(value))
}

func IsNotEmptyOrWhiteSpace(value string) bool {
	return !IsEmptyOrWhiteSpace(value)
}

func ValidateCompanyNameEN(text string) error {
	if !IsCompanyNameEN(text) {
		return errors.New("company name must be English")
	}
	return nil
}

func IsCompanyNameEN(text string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9\s.,()-]*$`).MatchString(text)
}

func ValidateCompanyNameTH(text string) error {
	if !IsCompanyNameTH(text) {
		return errors.New("company name must be Thai")
	}
	return nil
}

func IsCompanyNameTH(text string) bool {
	return regexp.MustCompile(`^[ก-๙0-9\s.,()-]*$`).MatchString(text)
}

func ValidateCompanyNameBasic(text string) error {
	if !IsCompanyNameBasic(text) {
		return errors.New("company name must be valid")
	}
	return nil
}

func IsCompanyNameBasic(text string) bool {
	return regexp.MustCompile(`^[a-zA-Zก-๙0-9\s.,()-]*$`).MatchString(text)
}

func ValidateName(text string) error {
	if !IsName(text) {
		return errors.New("name must be valid")
	}
	return nil
}

func IsName(text string) bool {
	return regexp.MustCompile(`^[a-zA-Zก-๙\s]*$`).MatchString(text)
}

func IsURL(checkUrl string) bool {
	return govalidator.IsURL(checkUrl)
}

func ValidateEnglishLang(value string) error {
	if !IsEnglishLang(value) {
		return errors.New("Must be eng")
	}
	return nil
}

func IsEnglishLang(value string) bool {
	for _, v := range value {
		if unicode.In(v, unicode.Thai) {
			return false
		}
	}
	return true
}

func ValidateThaiLang(value string) error {
	if !IsThaiLang(value) {
		return errors.New("Must be Thai")
	}
	return nil
}

func IsThaiLang(value string) bool {
	for _, v := range value {
		if !unicode.In(v, unicode.Thai, unicode.Space) {
			return false
		}
	}
	return true
}

func ValidateIDCardNo(idCardNo string) error {
	if IsValidIdCardNo(idCardNo) {
		return nil
	}
	return errors.New("IdCardNo wrong format")
}

func IsValidIdCardNo(idCardNo string) bool {
	match, err := regexp.MatchString("^\\d{13}$", idCardNo) //length = 13 && all digit
	if err != nil {
		return false
	}
	if !match {
		return false
	}

	index1, err := strconv.Atoi(idCardNo[0:1])
	if err != nil {
		return false
	}
	index2, err := strconv.Atoi(idCardNo[1:2])
	if err != nil {
		return false
	}
	index3, err := strconv.Atoi(idCardNo[2:3])
	if err != nil {
		return false
	}
	index4, err := strconv.Atoi(idCardNo[3:4])
	if err != nil {
		return false
	}
	index5, err := strconv.Atoi(idCardNo[4:5])
	if err != nil {
		return false
	}
	index6, err := strconv.Atoi(idCardNo[5:6])
	if err != nil {
		return false
	}
	index7, err := strconv.Atoi(idCardNo[6:7])
	if err != nil {
		return false
	}
	index8, err := strconv.Atoi(idCardNo[7:8])
	if err != nil {
		return false
	}
	index9, err := strconv.Atoi(idCardNo[8:9])
	if err != nil {
		return false
	}
	index10, err := strconv.Atoi(idCardNo[9:10])
	if err != nil {
		return false
	}
	index11, err := strconv.Atoi(idCardNo[10:11])
	if err != nil {
		return false
	}
	index12, err := strconv.Atoi(idCardNo[11:12])
	if err != nil {
		return false
	}

	sum := index1*13 + index2*12 + index3*11 + index4*10 + index5*9 + index6*8 + index7*7 + index8*6 + index9*5 + index10*4 + index11*3 + index12*2
	sum = 11 - (sum % 11)
	result := strconv.Itoa(sum)

	return idCardNo[12:13] == result[len(result)-1:]
}

func ValidatePassportNo(passportNo string) error {
	if IsValidPassport(passportNo) {
		return nil
	}
	return errors.New("passport must be valid")
}

func IsValidPassport(text string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString(text)
}

func ValidateNumber(value string) error {
	if !IsNumber(value) {
		return errors.New("must be a numeric string")
	}
	return nil
}

func IsNumber(value string) bool {
	_, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return true
}

func ValidateEmail(value string) error {
	if !IsEmail(value) {
		return errors.New("email must be valid")
	}
	return nil
}

func IsEmail(value string) bool {
	return govalidator.IsEmail(value)
}
