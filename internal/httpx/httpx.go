package httpx

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

const loopbackFallback = "1.1.1.1"

func IsLoopback(host string) bool {
	return strings.Contains(host, "localhost") || strings.Contains(host, "127.0.0.1")
}

func TargetHost(c fiber.Ctx) string {
	host := c.Query("host")
	if host == "" {
		host = c.IP()
	}
	if IsLoopback(host) {
		return loopbackFallback
	}
	return host
}

func Error(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{"error": message})
}
