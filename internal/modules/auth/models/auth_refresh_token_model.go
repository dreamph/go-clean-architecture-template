package models

import "backend/internal/core/auth/jwt"

type AuthRefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type AuthRefreshTokenResponse struct {
	Success bool           `json:"success"`
	Token   *jwt.TokenInfo `json:"token"`
}
