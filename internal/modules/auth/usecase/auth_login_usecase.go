package usecase

import (
	"backend/internal/constants"
	"backend/internal/core/auth/jwt"
	"backend/internal/core/errors"
	"backend/internal/modules/auth/models"
	"context"
	"strings"

	"github.com/dreamph/validation"
	errs "github.com/pkg/errors"
)

func (u *authUseCase) LoginValidate(request *models.AuthLoginRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.UserName, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *authUseCase) Login(ctx context.Context, request *models.AuthLoginRequest) (*models.AuthLoginResponse, error) {
	if err := u.LoginValidate(request); err != nil {
		return nil, errs.Wrap(errors.ErrValidationFailed, err.Error())
	}

	res := &models.AuthLoginResponse{
		Success: false,
	}

	if !strings.EqualFold(request.UserName, "admin") || !strings.EqualFold(request.Password, "admin") {
		return res, nil
	}

	token, err := u.jwtToken.Create(&jwt.TokenData{
		ID:    "adcc77b9-7565-4e55-b968-63280bc831cd",
		Scope: constants.ScopeUser,
	})
	if err != nil {
		return nil, err
	}

	res.Token = token
	res.User = &models.UserData{
		ID: "5b2fe55d-1235-4e28-afc4-1059a2793de8",
	}

	return res, nil
}
