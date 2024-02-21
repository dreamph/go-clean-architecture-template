package models

import "backend/internal/core/models"

type CompanyDeleteRequest struct {
	UserRequestInfo *models.UserRequestInfo `json:"-"`
	ID              string                  `json:"id"`
}

type CompanyDeleteResponse struct {
	Success bool `json:"success"`
}
