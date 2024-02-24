package repomodels

import (
	"backend/internal/core/models"
	"backend/internal/domain"

	"github.com/dreamph/dbre"
)

type CompanyListCriteria struct {
	Status      int32  `json:"status" example:"20"`
	Code        string `json:"code"`
	CompanyName string `json:"companyName"`
	HasP12File  int32  `json:"hasP12File" example:"20"`

	Limit *models.PageLimit `json:"limit"`
	Sort  *dbre.Sort        `json:"sort"`
}

type CompanyInfoData struct {
	domain.Company
	HasP12File int32 `json:"hasP12File" example:"20"`
}
