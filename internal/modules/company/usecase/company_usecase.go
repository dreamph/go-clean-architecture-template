package usecase

import (
	"backend/internal/modules/company"
	"backend/internal/repository"

	"github.com/dreamph/dbre"
)

type companyUseCase struct {
	dbTx              dbre.DBTx
	companyRepository repository.CompanyRepository
}

func NewCompanyUseCase(dbTx dbre.DBTx, companyRepository repository.CompanyRepository) company.CompanyUseCase {
	return &companyUseCase{
		dbTx:              dbTx,
		companyRepository: companyRepository,
	}
}
