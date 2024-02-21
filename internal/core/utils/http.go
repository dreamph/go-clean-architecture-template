package utils

import (
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{Timeout: time.Duration(30) * time.Second}

func HttpClient() *http.Client {
	return client
}

func IsMultipartContentType(contentType string) bool {
	return strings.Contains(contentType, "multipart")
}

func IsJsonContentType(contentType string) bool {
	return strings.Contains(contentType, "application/json")
}

func IsProblemJsonContentType(contentType string) bool {
	return strings.Contains(contentType, "application/problem+json")
}

func IsApplicationOctetStreamContentType(contentType string) bool {
	return strings.Contains(contentType, "octet-stream")
}
