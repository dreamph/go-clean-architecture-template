package usecase

import (
	"backend/internal/constants/errorcode"
	"backend/internal/core/errors"
	"backend/internal/modules/auth/models"
	"context"

	"backend/internal/core/validation"

	errs "github.com/pkg/errors"
)

func (u *authUseCase) RefreshTokenValidate(request *models.AuthRefreshTokenRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.RefreshToken, validation.Required),
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *authUseCase) RefreshToken(ctx context.Context, request *models.AuthRefreshTokenRequest) (*models.AuthRefreshTokenResponse, error) {
	if err := u.RefreshTokenValidate(request); err != nil {
		return nil, errs.Wrap(errors.ErrValidationFailed, err.Error())
	}

	err := u.jwtToken.Validate(request.RefreshToken)
	if err != nil {
		return nil, errors.ErrorUnauthorized(errorcode.ErrTokenInvalidOrExpired)
	}

	tokenData, err := u.jwtToken.GetTokenData(request.RefreshToken)
	if err != nil {
		return nil, errors.ErrorUnauthorized(errorcode.ErrTokenInvalidOrExpired)
	}

	token, err := u.jwtToken.Create(tokenData)
	if err != nil {
		return nil, err
	}

	res := &models.AuthRefreshTokenResponse{
		Success: true,
		Token:   token,
	}

	return res, nil
}
