package models

import (
	"backend/internal/core/errors"
	"time"
)

type RequestContext struct {
	RequestId string
}

type RequestInfo struct {
	UserRequestInfo *UserRequestInfo
	Token           string
}

type UserRequestInfo struct {
	UserRequestIP  string `json:"-"`
	ID             string `json:"-"`
	Scope          string `json:"-"`
	UserType       string `json:"-"`
	Org            string `json:"-"`
	OrgIssuer      string `json:"-"`
	Info           string `json:"-"`
	HasUserRequest bool   `json:"-"`
	IsBatch        bool   `json:"-"`
}

type InfoData struct {
}

type APIInfoResponse struct {
	Name   string    `json:"name"`
	Status string    `json:"status"`
	Time   time.Time `json:"time" swaggertype:"string" format:"date-time"`
}

type CheckAPIInfoResponse struct {
	Name    string    `json:"name"`
	Status  string    `json:"status"`
	Time    time.Time `json:"time" swaggertype:"string" format:"date-time"`
	Message string    `json:"message"`
}

type APIErrorResponse struct {
	Status        bool                   `json:"status"`
	StatusCode    int                    `json:"statusCode"`
	StatusMessage string                 `json:"statusMessage"`
	Type          string                 `json:"type"`
	Code          string                 `json:"code"`
	Message       string                 `json:"message"`
	ErrorMessage  string                 `json:"errorMessage"`
	Time          time.Time              `json:"time" swaggertype:"string" format:"date-time"`
	Detail        string                 `json:"detail"`
	ErrorData     *[]errors.AppErrorData `json:"errorData"`
	Cause         error                  `json:"-"`
}

type EchoRequest struct {
	Text string `json:"text"`
}

type TotalResponse struct {
	TotalProcessed int64 `json:"totalProcessed"`
	TotalError     int64 `json:"totalError"`
}

type AttemptLimitData struct {
	IsReachAttemptLimit bool       `json:"isReachAttemptLimit"`
	AvailableDateTime   *time.Time `json:"availableDateTime,omitempty" swaggertype:"string" format:"date-time"`
	AttemptLimit        int32      `json:"attemptLimit"`
	Attempts            int32      `json:"attempts"`
	AvailableTimeLeft   float64    `json:"availableTimeLeft"`
}
