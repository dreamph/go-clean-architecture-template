package models

import "github.com/gofiber/fiber/v2"

type ErrorStatus struct {
	FiberError *fiber.Error
	Code       string
}
