package usecase

import (
	cerrors "backend/internal/core/errors"
	"backend/internal/modules/company/models"
	"context"

	"backend/internal/core/validation"
	"backend/internal/core/validation/is"

	errs "github.com/pkg/errors"
)

func (u *companyUseCase) FindByIDValidate(request *models.CompanyFindByIDRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.ID, validation.Required, is.UUID),
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *companyUseCase) FindByID(ctx context.Context, request *models.CompanyFindByIDRequest) (*models.CompanyFindByIDResponse, error) {
	if err := u.FindByIDValidate(request); err != nil {
		return nil, errs.Wrap(cerrors.ErrValidationFailed, err.Error())
	}

	res := &models.CompanyFindByIDResponse{
		Success: false,
	}

	companyData, err := u.companyRepository.FindByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}
	if companyData == nil {
		return res, nil
	}

	res.Data = companyData

	return res, nil
}
