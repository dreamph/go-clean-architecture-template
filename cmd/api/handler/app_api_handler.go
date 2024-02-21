package handler

import (
	"backend/internal/core/api"
	"context"
	"time"

	"backend/internal/core/models"

	"github.com/gofiber/fiber/v2"
)

type AppAPIHandler struct {
	apiHandler api.ApiHandler
	router     fiber.Router
}

func NewAppAPIHandler(apiHandler api.ApiHandler, router fiber.Router) *AppAPIHandler {
	return &AppAPIHandler{
		apiHandler: apiHandler,
		router:     router,
	}
}

func (h *AppAPIHandler) Init() {
	router := h.router
	router.Get("/", h.AppIndex)
	router.Get("/health", h.AppHealth)
}

// AppIndex API
// @ID AppIndex
// @Tags app
// @Summary Home
// @Produce json
// @Success 200 {object} models.APIInfoResponse
// @Router / [get]
func (h *AppAPIHandler) AppIndex(c *fiber.Ctx) error {
	return h.apiHandler.Do(c, nil, func(ctx context.Context, requestInfo *models.RequestInfo) (interface{}, error) {
		return &models.APIInfoResponse{
			Name:   "API",
			Status: "Online",
			Time:   time.Now(),
		}, nil
	})
}

// AppHealth API
// @ID AppHealth
// @Tags app
// @Summary Health
// @Produce json
// @Success 200 {object} models.APIInfoResponse
// @Router /health [get]
func (h *AppAPIHandler) AppHealth(c *fiber.Ctx) error {
	return h.apiHandler.Do(c, nil, func(ctx context.Context, requestInfo *models.RequestInfo) (interface{}, error) {
		return &models.APIInfoResponse{
			Name:   "API",
			Status: "Healthy",
			Time:   time.Now(),
		}, nil
	})
}
