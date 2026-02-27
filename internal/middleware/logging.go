package middleware

import (
	"checkhost-backend/pkg/logging"
	"time"

	"github.com/gofiber/fiber/v3"
)

func Logging(c fiber.Ctx) error {
	start := time.Now()
	err := c.Next()

	latency := time.Since(start)
	logging.Logger.Info().
		Str("latency", latency.String()).
		Str("path", string(c.Request().URI().Path())).Int("status", c.Response().StatusCode()).
		Str("method", string(c.Request().Header.Method())).Msg("Intercepted request")

	return err
}
