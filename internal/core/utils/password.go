package utils

import (
	crand "crypto/rand"
	"math/big"
	"strings"
)

func GenerateRefCode(length int) string {
	return GeneratePassword(length)
}

func GenerateOTPCode(length int) string {
	const otpChars = "1234567890"
	buffer := make([]byte, length)
	_, err := crand.Read(buffer)
	if err != nil {
		return ""
	}
	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}
	return string(buffer)
}

func GeneratePassword(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	lengthNum := 8
	if length > 0 {
		lengthNum = length
	}

	var generateResult strings.Builder
	for i := 0; i < lengthNum; i++ {
		rdNum, _ := crand.Int(crand.Reader, big.NewInt(int64(len(chars))))
		generateResult.WriteRune(chars[rdNum.Int64()])
	}

	return generateResult.String()
}
