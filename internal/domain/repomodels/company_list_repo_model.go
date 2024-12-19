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

	Limit *models.PageLimit `json:"limit"`
	Sort  *dbre.Sort        `json:"sort"`
}

type CompanyInfoData struct {
	domain.Company
	OtherData string `json:"otherData"`
}
