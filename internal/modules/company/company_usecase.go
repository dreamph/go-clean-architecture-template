package company

import (
	"backend/internal/modules/company/models"
	"context"
)

type CompanyUseCase interface {
	Create(ctx context.Context, request *models.CompanyCreateRequest) (*models.CompanyCreateResponse, error)
	Update(ctx context.Context, request *models.CompanyUpdateRequest) (*models.CompanyUpdateResponse, error)

	Delete(ctx context.Context, request *models.CompanyDeleteRequest) (*models.CompanyDeleteResponse, error)
	FindByID(ctx context.Context, request *models.CompanyFindByIDRequest) (*models.CompanyFindByIDResponse, error)
	List(ctx context.Context, request *models.CompanyListRequest) (*models.CompanyListResponse, error)

	ExampleDbTransaction(ctx context.Context, request *models.CompanyExampleDbTransactionRequest) (*models.CompanyExampleDbTransactionResponse, error)
}
