package handler

import (
	"backend/internal/constants/permissions"
	"backend/internal/core/api"
	"backend/internal/core/api/middleware"
	coremodels "backend/internal/core/models"
	"backend/internal/modules/company"
	"backend/internal/modules/company/models"
	"context"

	"github.com/gofiber/fiber/v2"
)

type CompanyAPIHandler struct {
	apiHandler     api.ApiHandler
	router         fiber.Router
	companyUseCase company.CompanyUseCase
}

func NewCompanyAPIHandler(apiHandler api.ApiHandler, router fiber.Router, companyUseCase company.CompanyUseCase) *CompanyAPIHandler {
	return &CompanyAPIHandler{
		apiHandler:     apiHandler,
		router:         router,
		companyUseCase: companyUseCase,
	}
}

func (h *CompanyAPIHandler) Init(jwtAuth middleware.Auth, authPermissions middleware.AuthPermissions) {
	router := h.router
	router.Post("/companies",
		jwtAuth.Auth,
		authPermissions.RequiresPermissions([]string{permissions.CompanyCreate}),
		h.CompanyCreate,
	)
	router.Put("/companies",
		jwtAuth.Auth,
		authPermissions.RequiresPermissions([]string{permissions.CompanyUpdate}),
		h.CompanyUpdate,
	)

	router.Get("/companies/:id",
		jwtAuth.Auth,
		authPermissions.RequiresPermissions([]string{permissions.CompanyFindByID}),
		h.CompanyFindByID,
	)
	router.Delete("/companies",
		jwtAuth.Auth,
		authPermissions.RequiresPermissions([]string{permissions.CompanyDelete}),
		h.CompanyDelete,
	)
	router.Post("/companies/list",
		jwtAuth.Auth,
		authPermissions.RequiresPermissions([]string{permissions.CompanyList}),
		h.CompanyList,
	)
}

// CompanyList API
// @Security ApiKeyAuth
// @Tags companies
// @Summary Company List
// @Produce json
// @Param body body models.CompanyListRequest true "body"
// @Success 200 {object} models.CompanyListResponse
// @Router /companies/list [post]
func (h *CompanyAPIHandler) CompanyList(c *fiber.Ctx) error {
	request := &models.CompanyListRequest{}
	return h.apiHandler.Do(c, request, func(ctx context.Context, requestInfo *coremodels.RequestInfo) (interface{}, error) {
		request.UserRequestInfo = requestInfo.UserRequestInfo
		return h.companyUseCase.List(ctx, request)
	})
}

// CompanyFindByID API
// @Security ApiKeyAuth
// @Tags companies
// @Summary Company Find By ID
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.CompanyFindByIDResponse
// @Router /companies/{id} [get]
func (h *CompanyAPIHandler) CompanyFindByID(c *fiber.Ctx) error {
	request := &models.CompanyFindByIDRequest{ID: c.Params("id")}
	return h.apiHandler.Do(c, request, func(ctx context.Context, requestInfo *coremodels.RequestInfo) (interface{}, error) {
		request.UserRequestInfo = requestInfo.UserRequestInfo
		return h.companyUseCase.FindByID(ctx, request)
	})
}

// CompanyCreate API
// @Security ApiKeyAuth
// @Tags companies
// @Summary Company Create
// @Produce json
// @Param body body models.CompanyCreateRequest true "body"
// @Success 200 {object} models.CompanyCreateResponse
// @Router /companies [post]
func (h *CompanyAPIHandler) CompanyCreate(c *fiber.Ctx) error {
	request := &models.CompanyCreateRequest{}
	return h.apiHandler.Do(c, request, func(ctx context.Context, requestInfo *coremodels.RequestInfo) (interface{}, error) {
		request.UserRequestInfo = requestInfo.UserRequestInfo
		return h.companyUseCase.Create(ctx, request)
	})
}

// CompanyUpdate API
// @Security ApiKeyAuth
// @Tags companies
// @Summary Company Update
// @Produce json
// @Param body body models.CompanyUpdateRequest true "body"
// @Success 200 {object} models.CompanyUpdateResponse
// @Router /companies [put]
func (h *CompanyAPIHandler) CompanyUpdate(c *fiber.Ctx) error {
	request := &models.CompanyUpdateRequest{}
	return h.apiHandler.Do(c, request, func(ctx context.Context, requestInfo *coremodels.RequestInfo) (interface{}, error) {
		request.UserRequestInfo = requestInfo.UserRequestInfo
		return h.companyUseCase.Update(ctx, request)
	})
}

// CompanyDelete API
// @Security ApiKeyAuth
// @Tags companies
// @Summary Company Delete
// @Produce json
// @Param body body models.CompanyDeleteRequest true "body"
// @Success 200 {object} models.CompanyDeleteResponse
// @Router /companies [delete]
func (h *CompanyAPIHandler) CompanyDelete(c *fiber.Ctx) error {
	request := &models.CompanyDeleteRequest{}
	return h.apiHandler.Do(c, request, func(ctx context.Context, requestInfo *coremodels.RequestInfo) (interface{}, error) {
		request.UserRequestInfo = requestInfo.UserRequestInfo
		return h.companyUseCase.Delete(ctx, request)
	})
}
