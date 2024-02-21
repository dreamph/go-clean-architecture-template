package messages

import (
	coremodels "backend/internal/core/models"
	"log"
)

var messageMap map[string]string

func Messages() map[string]string {
	return messageMap
}

func New(code string, message string) *coremodels.Message {
	if messageMap == nil {
		messageMap = map[string]string{}
	}
	if _, exists := messageMap[code]; exists {
		log.Fatalf("%s : messageCode Duplicate!", code)
	} else {
		messageMap[code] = message
	}
	return &coremodels.Message{Code: code, Message: message}
}
