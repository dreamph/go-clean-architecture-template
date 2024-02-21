package models

import (
	"backend/internal/core/models"
	"backend/internal/domain"
)

type CompanyCreateRequest struct {
	UserRequestInfo *models.UserRequestInfo `json:"-"`
	Code            string                  `json:"code" form:"code" query:"code"`
	Name            string                  `json:"name" form:"name" query:"name"`
}

type CompanyCreateResponse struct {
	Success bool            `json:"success"`
	Data    *domain.Company `json:"data"`
}
