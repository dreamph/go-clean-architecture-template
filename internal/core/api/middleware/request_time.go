package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func NewRequestTime() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		start := time.Now()
		err = c.Next()
		stop := time.Now()
		c.Append("server-timing", fmt.Sprintf("app;dur=%v", stop.Sub(start).String()))
		return err
	}
}
