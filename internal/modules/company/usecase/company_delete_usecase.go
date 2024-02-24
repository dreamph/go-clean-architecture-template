package usecase

import (
	"backend/internal/constants/errorcode"
	cerrors "backend/internal/core/errors"
	"backend/internal/modules/company/models"
	"context"

	"github.com/dreamph/validation"
	"github.com/dreamph/validation/is"

	errs "github.com/pkg/errors"
)

func (u *companyUseCase) DeleteValidate(request *models.CompanyDeleteRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.ID, validation.Required, is.UUID),
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *companyUseCase) Delete(ctx context.Context, request *models.CompanyDeleteRequest) (*models.CompanyDeleteResponse, error) {
	if err := u.DeleteValidate(request); err != nil {
		return nil, errs.Wrap(cerrors.ErrValidationFailed, err.Error())
	}

	data, err := u.companyRepository.FindByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, cerrors.ErrorValidation(errorcode.ErrCompanyNotFound)
	}

	err = u.companyRepository.Delete(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	res := &models.CompanyDeleteResponse{
		Success: true,
	}
	return res, nil
}
