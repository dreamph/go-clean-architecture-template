package usecase

import (
	"backend/internal/constants"
	"backend/internal/core/utils"
	"backend/internal/domain"
	"backend/internal/modules/company/models"
	"context"
	"time"

	"github.com/dreamph/dbre/query"
	"github.com/guregu/null"
)

func (u *companyUseCase) ExampleDbTransaction(ctx context.Context, request *models.CompanyExampleDbTransactionRequest) (*models.CompanyExampleDbTransactionResponse, error) {
	currentDate := time.Now()

	// With Transaction
	err := u.dbTx.WithTx(ctx, func(ctx context.Context, appDB query.AppIDB) error {
		_, err := u.companyRepository.WithTx(appDB).Create(ctx, &domain.Company{
			ID:   utils.NewID(),
			Code: null.StringFrom("3333"),
			Name: request.Name,

			Status:     constants.StatusActive,
			CreateDate: currentDate,
			CreateBy:   "1234",
			UpdateDate: currentDate,
			UpdateBy:   "1234",
		})
		if err != nil {
			return err
		}

		_, err = u.companyRepository.WithTx(appDB).Create(ctx, &domain.Company{
			ID:   utils.NewID(),
			Code: null.StringFrom("44444"),
			Name: request.Name,

			Status:     constants.StatusActive,
			CreateDate: currentDate,
			CreateBy:   "1234",
			UpdateDate: currentDate,
			UpdateBy:   "1234",
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &models.CompanyExampleDbTransactionResponse{
		Success: true,
	}, nil
}
