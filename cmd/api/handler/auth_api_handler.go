package handler

import (
	"backend/internal/core/api"
	coremodels "backend/internal/core/models"
	"backend/internal/modules/auth"
	"backend/internal/modules/auth/models"
	"context"

	"github.com/gofiber/fiber/v2"
)

type AuthAPIHandler struct {
	apiHandler  api.ApiHandler
	router      fiber.Router
	authUseCase auth.AuthUseCase
}

func NewAuthAPIHandler(apiHandler api.ApiHandler, router fiber.Router, authUseCase auth.AuthUseCase) *AuthAPIHandler {
	return &AuthAPIHandler{
		apiHandler:  apiHandler,
		router:      router,
		authUseCase: authUseCase,
	}
}

func (h *AuthAPIHandler) Init() {
	router := h.router
	router.Post("/auth/login", h.AuthLogin)
	router.Post("/auth/refresh", h.AuthRefreshToken)
}

// AuthLogin API
// @ID AuthLogin
// @Tags auth
// @Summary Auth Login
// @Produce json
// @Param body body models.AuthLoginRequest true "body"
// @Success 200 {object} models.AuthLoginResponse
// @Success 200 {string} string
// @Router /auth/login [post]
func (h *AuthAPIHandler) AuthLogin(c *fiber.Ctx) error {
	request := &models.AuthLoginRequest{}
	return h.apiHandler.Do(c, request, func(ctx context.Context, requestInfo *coremodels.RequestInfo) (interface{}, error) {
		return h.authUseCase.Login(ctx, request)
	})
}

// AuthRefreshToken API
// @ID AuthRefreshToken
// @Tags auth
// @Summary Auth Refresh Token
// @Produce json
// @Param body body models.AuthRefreshTokenRequest true "body"
// @Success 200 {object} models.AuthRefreshTokenResponse
// @Success 200 {string} string
// @Router /auth/refresh [post]
func (h *AuthAPIHandler) AuthRefreshToken(c *fiber.Ctx) error {
	request := &models.AuthRefreshTokenRequest{}
	return h.apiHandler.Do(c, request, func(ctx context.Context, requestInfo *coremodels.RequestInfo) (interface{}, error) {
		return h.authUseCase.RefreshToken(ctx, request)
	})
}
