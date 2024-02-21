package auth

import (
	"backend/internal/modules/auth/models"
	"context"
)

type AuthUseCase interface {
	Login(ctx context.Context, request *models.AuthLoginRequest) (*models.AuthLoginResponse, error)
	RefreshToken(ctx context.Context, request *models.AuthRefreshTokenRequest) (*models.AuthRefreshTokenResponse, error)
}
