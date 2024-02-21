package usecase

import (
	"backend/internal/constants"
	cerrors "backend/internal/core/errors"
	"backend/internal/core/utils"
	"backend/internal/modules/company/models"
	"context"

	"backend/internal/core/validation"

	errs "github.com/pkg/errors"
)

func (u *companyUseCase) ListValidate(request *models.CompanyListRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Limit, validation.Required, validation.PageLimit),

		// optional
		validation.Field(&request.Status, validation.In(constants.StatusInActive, constants.StatusActive)),
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *companyUseCase) List(ctx context.Context, request *models.CompanyListRequest) (*models.CompanyListResponse, error) {
	if err := u.ListValidate(request); err != nil {
		return nil, errs.Wrap(cerrors.ErrValidationFailed, err.Error())
	}

	data, total, err := u.companyRepository.ListData(ctx, request.CompanyListCriteria)
	if err != nil {
		return nil, err
	}

	res := &models.CompanyListResponse{
		Success: true,
		Data:    data,
		Page:    utils.GetPageResult(request.Limit, total),
	}

	return res, nil
}
