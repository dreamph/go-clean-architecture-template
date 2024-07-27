package models

import (
	"backend/internal/core/models"
)

type CompanyExampleDbTransactionRequest struct {
	UserRequestInfo *models.UserRequestInfo `json:"-"`
	Code            string                  `json:"code" form:"code" query:"code"`
	Name            string                  `json:"name" form:"name" query:"name"`
}

type CompanyExampleDbTransactionResponse struct {
	Success bool `json:"success"`
}
