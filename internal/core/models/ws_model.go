package models

type WSNotificationNotify struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
