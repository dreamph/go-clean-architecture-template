package models

import (
	"backend/internal/core/models"
	"backend/internal/domain"
)

type CompanyUpdateRequest struct {
	UserRequestInfo *models.UserRequestInfo `json:"-"`
	ID              string                  `json:"id" form:"id" query:"id"`
	Name            string                  `json:"name" form:"name" query:"name"`
}

type CompanyUpdateResponse struct {
	Success bool            `json:"success"`
	Data    *domain.Company `json:"data"`
}
