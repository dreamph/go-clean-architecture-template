package usecase

import (
	"backend/internal/core/database/query"
	"backend/internal/modules/company"
	"backend/internal/repository"
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
