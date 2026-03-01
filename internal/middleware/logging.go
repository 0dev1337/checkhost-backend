package middleware

import (
	"time"

	"checkhost-backend/pkg/logging"

	"github.com/gofiber/fiber/v3"
)

func Logging(c fiber.Ctx) error {
	start := time.Now()
	err := c.Next()

	logging.Logger.Info().
		Str("method", string(c.Request().Header.Method())).
		Str("path", string(c.Request().URI().Path())).
		Int("status", c.Response().StatusCode()).
		Int64("latency_ms", time.Since(start).Milliseconds()).
		Err(err).
		Msg("request")

	return err
}
