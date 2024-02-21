package commons

import (
	apimodels "backend/internal/core/api/models"
	"backend/internal/core/errorcode"
	cerrors "backend/internal/core/errors"
	coremodels "backend/internal/core/models"
	"backend/internal/core/utils"
	"context"
	"errors"
	"os"
	"strings"
	"time"

	realip "github.com/ferluci/fast-realip"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	fiberutils "github.com/gofiber/fiber/v2/utils"
)

func GetClientIP(c *fiber.Ctx) string {
	return realip.FromRequest(c.Context())
	//utils.CopyString(c.IPs()[0]))
}

func IsContentTypeApplicationOcspRequest(contentType string) bool {
	return contentType == "application/ocsp-request"
}

func IsContentTypeApplicationJSON(contentType string) bool {
	contentType = fiberutils.ToLower(contentType)
	contentType = fiberutils.ParseVendorSpecificContentType(contentType)
	return strings.HasPrefix(contentType, fiber.MIMEApplicationJSON)
}

func IsMultipartForm(contentType string) bool {
	contentType = fiberutils.ToLower(contentType)
	contentType = fiberutils.ParseVendorSpecificContentType(contentType)
	return strings.HasPrefix(contentType, fiber.MIMEMultipartForm)
}

func ResponseDownloadSuccessByFile(c *fiber.Ctx, fileResponse *coremodels.FileDownloadByFileResponse) error {
	c.Response().Header.Add(fiber.HeaderContentDisposition, utils.GetDownloadFileHeaderValue(fileResponse.DownloadFileName))
	defer os.Remove(fileResponse.FilePath)
	return c.SendFile(fileResponse.FilePath, false)
}

func ResponseDownloadSuccessByBytes(c *fiber.Ctx, fileResponse *coremodels.FileDownloadByBytesResponse) error {
	filePath := utils.GenerateFileName(fileResponse.DownloadFileName)
	_ = utils.WriteFile(filePath, fileResponse.FileData)
	defer os.Remove(filePath)
	c.Response().Header.Add(fiber.HeaderContentDisposition, utils.GetDownloadFileHeaderValue(fileResponse.DownloadFileName))
	return c.SendFile(filePath, false)
}

// ResponseSuccess ...
func ResponseSuccess(c *fiber.Ctx, payload interface{}) error {
	return ResponseSuccessWithStatusCode(c, 200, payload)
}

// ResponseSuccessWithStatusCode ...
func ResponseSuccessWithStatusCode(c *fiber.Ctx, code int, payload interface{}) error {
	return c.Status(code).JSON(payload)
}

func ResponseError(c *fiber.Ctx, err error) error {
	var appError *cerrors.AppError
	ok := errors.As(err, &appError)
	if ok {
		errCause := appError.ErrType
		errorStatus := MapErrorStatus(errCause)
		return ResponseErrorWithCode(c, errorStatus, err)
	} else {
		errCause := cerrors.GetErrorCause(err)
		var appError *cerrors.AppError
		ok := errors.As(errCause, &appError)
		if ok {
			errorStatus := MapErrorStatus(appError.ErrType)
			return ResponseErrorWithCode(c, errorStatus, appError)
		} else {
			errorStatus := MapErrorStatus(errCause)
			return ResponseErrorWithCode(c, errorStatus, err)
		}
	}
}

// ResponseErrorWithCode ...
func ResponseErrorWithCode(c *fiber.Ctx, errorStatus *apimodels.ErrorStatus, err error) error {
	requestID := utils.InterfaceToString(c.Locals(requestid.ConfigDefault.ContextKey))
	httpStatus := errorStatus.FiberError
	errInternalErrorDefault := errorcode.ErrInternalErrorDefault
	apiErrorResponse := &coremodels.APIErrorResponse{
		Status:        false,
		StatusCode:    httpStatus.Code,
		StatusMessage: httpStatus.Message,
		Time:          time.Now(),
		ErrorMessage:  err.Error(),
		Detail:        "uri:" + c.OriginalURL() + "|x-request-id:" + requestID,
		Cause:         err,
	}

	if httpStatus.Code == 500 {
		apiErrorResponse.Message = errInternalErrorDefault.Message
	} else {
		apiErrorResponse.Message = err.Error()
	}

	apiErrorResponse.Type = errorStatus.Code
	apiErrorResponse.Code = errInternalErrorDefault.Code

	var appError *cerrors.AppError
	ok := errors.As(err, &appError)
	if ok {
		if utils.IsNotEmpty(appError.ErrCode) {
			apiErrorResponse.Code = appError.ErrCode
		}
		if utils.IsNotEmpty(appError.ErrMessage) {
			apiErrorResponse.Message = appError.ErrMessage
		}
		if appError.ErrorData != nil {
			apiErrorResponse.ErrorData = appError.ErrorData
		}
	}
	return c.Status(httpStatus.Code).JSON(apiErrorResponse)
}

type NotifyRequestInfo struct {
	Env         string
	Name        string
	Time        time.Time
	Method      string
	Path        string
	RequestID   string
	TraceID     string
	TraceURL    string
	StatusCode  int
	Body        string
	Error       string
	ClientIP    string
	RequestInfo string
}

func MapErrorStatus(errCause error) *apimodels.ErrorStatus {
	if errors.Is(errCause, cerrors.ErrUnauthorized) {
		return &apimodels.ErrorStatus{
			FiberError: fiber.ErrUnauthorized,
			Code:       cerrors.ErrUnauthorized.Error(),
		}
	}
	if errors.Is(errCause, cerrors.ErrPermissionDenied) {
		return &apimodels.ErrorStatus{
			FiberError: fiber.ErrForbidden,
			Code:       cerrors.ErrPermissionDenied.Error(),
		}
	}
	if errors.Is(errCause, cerrors.ErrAuthInvalidRequest) {
		return &apimodels.ErrorStatus{
			FiberError: fiber.ErrBadRequest,
			Code:       cerrors.ErrAuthInvalidRequest.Error(),
		}
	}
	if errors.Is(errCause, cerrors.ErrValidationFailed) {
		return &apimodels.ErrorStatus{
			FiberError: fiber.ErrBadRequest,
			Code:       cerrors.ErrValidationFailed.Error(),
		}
	}
	if errors.Is(errCause, cerrors.ErrCallExternalServiceFailed) {
		return &apimodels.ErrorStatus{
			FiberError: fiber.ErrBadRequest,
			Code:       cerrors.ErrCallExternalServiceFailed.Error(),
		}
	}
	if errors.Is(errCause, cerrors.ErrNotFound) {
		return &apimodels.ErrorStatus{
			FiberError: fiber.ErrNotFound,
			Code:       cerrors.ErrNotFound.Error(),
		}
	}
	return &apimodels.ErrorStatus{
		FiberError: fiber.ErrInternalServerError,
		Code:       cerrors.ErrInternalServerError.Error(),
	}
}

func GetContext(c *fiber.Ctx) context.Context {
	return c.UserContext()
}
