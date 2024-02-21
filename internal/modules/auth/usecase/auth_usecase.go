package usecase

import (
	"backend/internal/core/auth/jwt"
	"backend/internal/modules/auth"
)

type authUseCase struct {
	jwtToken jwt.JwtToken
}

func NewAuthUseCase(jwtToken jwt.JwtToken) auth.AuthUseCase {
	return &authUseCase{
		jwtToken: jwtToken,
	}
}
