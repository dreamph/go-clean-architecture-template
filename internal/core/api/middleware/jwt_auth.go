package middleware

import (
	apicommons "backend/internal/core/api/commons"
	"backend/internal/core/auth/jwt"
	coreconstants "backend/internal/core/constants"
	"backend/internal/core/errors"

	"github.com/gofiber/fiber/v2"
	errs "github.com/pkg/errors"
)

type JwtAuth struct {
	option *JwtAuthOption
}

type JwtAuthOption struct {
	JwtToken         jwt.JwtToken
	ExternalValidate func(c *fiber.Ctx, token string) error
	Enable           bool
}

func NewJwtAuth(jwtOption *JwtAuthOption) Auth {
	return &JwtAuth{
		option: jwtOption,
	}
}

func (j *JwtAuth) Auth(c *fiber.Ctx) error {
	if !j.option.Enable {
		return c.Next()
	}

	token := jwt.ExtractToken(c.Get(coreconstants.AuthorizationHeaderName))
	err := j.option.JwtToken.Validate(token)
	if err != nil {
		return apicommons.ResponseError(c, errs.Wrap(errors.ErrUnauthorized, "Token invalid"))
	}

	if j.option.ExternalValidate != nil {
		err = j.option.ExternalValidate(c, token)
		if err != nil {
			return apicommons.ResponseError(c, errs.Wrap(errors.ErrUnauthorized, "Token invalid"))
		}
	}
	return c.Next()
}
