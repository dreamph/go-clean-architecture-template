package middleware

import (
	apicommons "backend/internal/core/api/commons"
	"backend/internal/core/errors"

	"github.com/gofiber/fiber/v2"
	errs "github.com/pkg/errors"
)

const (
	ApiKeyModeStatic  = 1
	ApiKeyModeDynamic = 2
)

type apiKeyAuth struct {
	option *ApiKeyOption
}

type ApiKeyOption struct {
	Enable          bool
	HeaderKey       string
	Key             string
	ApiKeyMode      int
	Skip            func(c *fiber.Ctx) bool
	DynamicValidate func(c *fiber.Ctx) error
}

func NewApiKeyAuth(option *ApiKeyOption) Auth {
	return &apiKeyAuth{
		option: option,
	}
}

func (j *apiKeyAuth) Auth(c *fiber.Ctx) error {
	if !j.option.Enable {
		return c.Next()
	}
	if j.option.Skip != nil && j.option.Skip(c) {
		return c.Next()
	}
	token := c.Get(j.option.HeaderKey)
	if j.option.ApiKeyMode == ApiKeyModeDynamic {
		err := j.option.DynamicValidate(c)
		if err != nil {
			return apicommons.ResponseError(c, errs.Wrap(errors.ErrUnauthorized, "Invalid api-key"))
		}
	} else {
		if token != j.option.Key {
			return apicommons.ResponseError(c, errs.Wrap(errors.ErrUnauthorized, "Invalid api-key"))
		}
	}

	return c.Next()
}
