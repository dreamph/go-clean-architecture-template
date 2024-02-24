package usecase

import (
	"backend/internal/modules/company"
	"backend/internal/repository"

	"github.com/dreamph/dbre/query"
)

type companyUseCase struct {
	dbTx              query.DBTx
	companyRepository repository.CompanyRepository
}

func NewCompanyUseCase(dbTx query.DBTx, companyRepository repository.CompanyRepository) company.CompanyUseCase {
	return &companyUseCase{
		dbTx:              dbTx,
		companyRepository: companyRepository,
	}
}
