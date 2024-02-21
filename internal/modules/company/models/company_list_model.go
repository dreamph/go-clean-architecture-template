package models

import (
	"backend/internal/core/models"
	"backend/internal/domain/repomodels"
)

type CompanyListRequest struct {
	UserRequestInfo *models.UserRequestInfo `json:"-"`
	*repomodels.CompanyListCriteria
}

type CompanyListResponse struct {
	Success bool                          `json:"success"`
	Data    *[]repomodels.CompanyInfoData `json:"data"`
	Page    *models.PageResult            `json:"page"`
}
