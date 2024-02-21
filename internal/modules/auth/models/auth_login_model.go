package models

import (
	"backend/internal/core/auth/jwt"
)

type AuthLoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	Success bool           `json:"success"`
	Token   *jwt.TokenInfo `json:"token"`
	User    *UserData      `json:"user"`
}

type UserData struct {
	ID string `json:"id"`
}
