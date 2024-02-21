package usecase

import (
	"backend/internal/constants"
	"backend/internal/constants/errorcode"
	cerrors "backend/internal/core/errors"
	"backend/internal/core/utils"
	"backend/internal/domain"
	"backend/internal/modules/company/models"
	"context"
	"time"

	"backend/internal/core/validation"

	"github.com/guregu/null"
	errs "github.com/pkg/errors"
)

func (u *companyUseCase) CreateValidate(request *models.CompanyCreateRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required),
		validation.Field(&request.Name, validation.Required),
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *companyUseCase) CreateCheckPermissionBeforeProcess(request *models.CompanyCreateRequest) error {
	role := request.UserRequestInfo.Scope
	if role == constants.ScopeAdmin {
		return nil
	}
	return cerrors.ErrPermissionDenied
}

func (u *companyUseCase) Create(ctx context.Context, request *models.CompanyCreateRequest) (*models.CompanyCreateResponse, error) {
	if err := u.CreateValidate(request); err != nil {
		return nil, errs.Wrap(cerrors.ErrValidationFailed, err.Error())
	}
	if err := u.CreateCheckPermissionBeforeProcess(request); err != nil {
		return nil, err
	}

	company, err := u.companyRepository.FindByCode(ctx, request.Code)
	if err != nil {
		return nil, err
	}
	if company != nil {
		return nil, cerrors.ErrorValidation(errorcode.ErrCompanyCodeAlreadyExists)
	}

	userId := request.UserRequestInfo.ID
	currentDate := time.Now()
	id := utils.NewID()

	data := &domain.Company{
		ID:   id,
		Code: null.StringFrom(request.Code),
		Name: request.Name,

		Status:     constants.StatusActive,
		CreateDate: currentDate,
		CreateBy:   userId,
		UpdateDate: currentDate,
		UpdateBy:   userId,
	}

	_, err = u.companyRepository.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &models.CompanyCreateResponse{
		Success: true,
		Data:    data,
	}, nil
}
