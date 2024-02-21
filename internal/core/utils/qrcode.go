package utils

import "github.com/skip2/go-qrcode"

func CreateQrCode(text string, size int) ([]byte, error) {
	return qrcode.Encode(text, qrcode.Medium, size /*256*/)
}

func WriteQrCodeToFile(text string, size int, filename string) error {
	return qrcode.WriteFile(text, qrcode.Medium, size, filename)
}
