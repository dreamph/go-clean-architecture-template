package models

import (
	"backend/internal/core/models"
	"backend/internal/domain"
)

type CompanyFindByIDRequest struct {
	UserRequestInfo *models.UserRequestInfo `json:"-"`
	ID              string                  `json:"id"`
}

type CompanyFindByIDResponse struct {
	Success bool            `json:"success"`
	Data    *domain.Company `json:"data"`
}
