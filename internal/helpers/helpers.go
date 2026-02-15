package helpers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func Response(c fiber.Ctx, status int, data ...any) error {
	message := ""

	if len(data) > 0 {
		message = fmt.Sprintf("%v", data[0])
	} else {
		message = fmt.Sprintf("%v: %v", status, http.StatusText(status))
	}

	return c.Status(status).JSON(fiber.Map{
		"code":    status,
		"message": message,
	})
}
