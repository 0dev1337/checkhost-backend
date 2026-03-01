package helpers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func Response(c fiber.Ctx, status int, message string) error {
	if message == "" {
		message = http.StatusText(status)
	}
	return c.Status(status).JSON(fiber.Map{
		"code":    status,
		"message": message,
	})
}
