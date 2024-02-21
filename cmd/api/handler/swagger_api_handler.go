package handler

import (
	"backend/docs"
	"backend/internal/core/api"
	apiswagger "backend/internal/core/api/swagger"

	"github.com/gofiber/fiber/v2"
)

type SwaggerAPIHandler struct {
	apiHandler api.ApiHandler
	router     fiber.Router
}

func NewSwaggerAPIHandler(apiHandler api.ApiHandler, router fiber.Router) *SwaggerAPIHandler {
	return &SwaggerAPIHandler{
		apiHandler: apiHandler,
		router:     router,
	}
}

func (h *SwaggerAPIHandler) Init(contextPath string, prodMode bool) {
	router := h.router

	apiswagger.RegisterSwaggerRouter(
		router,
		contextPath,
		prodMode,
		docs.SwaggerInfo,
	)
}
