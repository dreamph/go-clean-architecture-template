package middleware

import "github.com/gofiber/fiber/v2"

type Auth interface {
	Auth(c *fiber.Ctx) error
}
