package middleware

import (
	"slices"

	"checkhost-backend/internal/httpx"

	"github.com/gofiber/fiber/v3"
)

func Sanitize(allowedNodes []string) fiber.Handler {
	return func(c fiber.Ctx) error {
		host := c.Query("host")
		if host == "" {
			return httpx.Error(c, fiber.StatusBadRequest, "host query parameter is required")
		}

		node := c.Query("node")
		if node == "" {
			return httpx.Error(c, fiber.StatusBadRequest, "node query parameter is required")
		}
		if !slices.Contains(allowedNodes, node) {
			return httpx.Error(c, fiber.StatusForbidden, "node is not allowed")
		}

		if httpx.IsLoopback(host) {
			return httpx.Error(c, fiber.StatusBadRequest, "localhost and loopback addresses are not allowed")
		}
		return c.Next()
	}
}
