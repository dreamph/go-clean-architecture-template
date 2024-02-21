package api

import (
	apicommons "backend/internal/core/api/commons"
	apimodels "backend/internal/core/api/models"
	cerrors "backend/internal/core/errors"
	"backend/internal/core/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func NewServerConfig() fiber.Config {
	fe := FiberErrorHandler{}
	cfg := fiber.Config{
		ErrorHandler: fe.ErrorHandler,

		//customize
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
	return cfg
}

type FiberErrorHandler struct {
}

func (f *FiberErrorHandler) ErrorHandler(c *fiber.Ctx, err error) error {
	fiberError := fiber.ErrInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		fiberError = e
	}
	_ = c.SendStatus(fiberError.Code)

	errorStatus := &apimodels.ErrorStatus{
		FiberError: fiberError,
		Code:       cerrors.ErrInternalServerError.Error(),
	}

	return apicommons.ResponseErrorWithCode(c, errorStatus, err)
}
