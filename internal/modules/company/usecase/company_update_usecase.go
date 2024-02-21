package usecase

import (
	"backend/internal/constants"
	"backend/internal/constants/errorcode"
	cerrors "backend/internal/core/errors"
	"backend/internal/modules/company/models"
	"context"
	"time"

	"backend/internal/core/validation"
	"backend/internal/core/validation/is"

	errs "github.com/pkg/errors"
)

func (u *companyUseCase) UpdateValidate(request *models.CompanyUpdateRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.ID, validation.Required, is.UUID),
		validation.Field(&request.Name, validation.Required),
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *companyUseCase) UpdateCheckPermissionBeforeProcess(request *models.CompanyUpdateRequest) error {
	role := request.UserRequestInfo.Scope
	if role == constants.ScopeAdmin {
		return nil
	}
	return cerrors.ErrPermissionDenied
}

func (u *companyUseCase) Update(ctx context.Context, request *models.CompanyUpdateRequest) (*models.CompanyUpdateResponse, error) {
	if err := u.UpdateValidate(request); err != nil {
		return nil, errs.Wrap(cerrors.ErrValidationFailed, err.Error())
	}
	if err := u.UpdateCheckPermissionBeforeProcess(request); err != nil {
		return nil, err
	}

	data, err := u.companyRepository.FindByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, cerrors.ErrorValidation(errorcode.ErrCompanyNotFound)
	}

	userId := request.UserRequestInfo.ID

	data.Name = request.Name
	data.UpdateBy = userId
	data.UpdateDate = time.Now()

	_, err = u.companyRepository.Update(ctx, data)
	if err != nil {
		return nil, err
	}

	return &models.CompanyUpdateResponse{
		Success: true,
		Data:    data,
	}, nil
}
