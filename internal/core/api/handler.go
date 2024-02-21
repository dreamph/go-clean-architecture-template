package api

import (
	apicommons "backend/internal/core/api/commons"
	"backend/internal/core/appcontext"
	"backend/internal/core/auth/jwt"
	applogger "backend/internal/core/logger"
	coremodels "backend/internal/core/models"
	"backend/internal/core/utils"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type DoFunc func(ctx context.Context, requestInfo *coremodels.RequestInfo) (interface{}, error)

type ApiHandler interface {
	GetUserRequestInfo(c *fiber.Ctx) *coremodels.UserRequestInfo
	Do(c *fiber.Ctx, request interface{}, doFunc DoFunc) error
}

type HandlerOption struct {
	JwtToken                   jwt.JwtToken
	TransformRequestBodyEnable bool
	TransformRequestBody       func(requestInfo *coremodels.RequestInfo, requestPtr interface{}) (interface{}, error)
}

type apiHandler struct {
	option *HandlerOption
}

func NewApiHandler(option *HandlerOption) ApiHandler {
	return &apiHandler{
		option: option,
	}
}

func (h *apiHandler) GetUserRequestInfo(c *fiber.Ctx) *coremodels.UserRequestInfo {
	return getUserRequestInfo(c, h.option.JwtToken)
}

func (h *apiHandler) Do(c *fiber.Ctx, requestPtr interface{}, doFunc DoFunc) error {
	hasBody, err := h.bodyParserIfRequired(c, requestPtr)
	if err != nil {
		return err
	}

	ctx := apicommons.GetContext(c)
	ctx = h.initializeValueToContext(c, ctx)

	requestInfo := &coremodels.RequestInfo{
		UserRequestInfo: h.GetUserRequestInfo(c),
		Token:           getUserRequestToken(c),
	}
	requestPtr, err = h.transformRequestBodyIfRequired(hasBody, requestInfo, requestPtr)
	if err != nil {
		return err
	}

	data, err := doFunc(ctx, requestInfo)
	if err != nil {
		return apicommons.ResponseError(c, err)
	}

	fileDownloadByBytesResponse, ok := data.(*coremodels.FileDownloadByBytesResponse)
	if ok {
		return apicommons.ResponseDownloadSuccessByBytes(c, fileDownloadByBytesResponse)
	}

	fileDownloadByFileResponse, ok := data.(*coremodels.FileDownloadByFileResponse)
	if ok {
		return apicommons.ResponseDownloadSuccessByFile(c, fileDownloadByFileResponse)
	}

	rawData, ok := data.(*coremodels.RawResponse)
	if ok {
		return c.Status(rawData.HttpStatus).Send(rawData.Data)
	}

	return apicommons.ResponseSuccess(c, data)
}

func (h *apiHandler) transformRequestBodyIfRequired(hasBody bool, requestInfo *coremodels.RequestInfo, requestPtr interface{}) (interface{}, error) {
	if hasBody && h.option.TransformRequestBodyEnable && h.option.TransformRequestBody != nil {
		return h.option.TransformRequestBody(requestInfo, requestPtr)
	}
	return requestPtr, nil
}

func (h *apiHandler) bodyParserIfRequired(c *fiber.Ctx, requestPtr interface{}) (bool, error) {
	if c.Method() == http.MethodGet {
		return false, nil
	}

	if requestPtr == nil {
		return false, nil
	}

	err := c.BodyParser(requestPtr)
	if err != nil {
		return false, apicommons.ResponseError(c, err)
	}

	return true, nil
}

func (h *apiHandler) initializeValueToContext(c *fiber.Ctx, ctx context.Context) context.Context {
	requestId := utils.InterfaceToString(c.Locals(requestid.ConfigDefault.ContextKey))
	info := map[string]string{
		"requestId": requestId,
	}
	return appcontext.WithValue(applogger.WithValue(ctx, info), &coremodels.RequestContext{RequestId: requestId})
}
