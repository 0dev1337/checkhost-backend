package handler

import (
	"checkhost-backend/internal/client/ipapi"
	"checkhost-backend/pkg/logging"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func Info(c fiber.Ctx) error {
	ip := c.Query("host")
	if ip == "" {
		ip = c.IP()
	}
	if strings.Contains(ip, "localhost") || strings.Contains(ip, "127.0.0.1") {
		ip = "1.1.1.1"
	}

	info, err := ipapi.Fetch(ip)
	if err != nil {
		logging.Logger.Error().Err(err).Str("ip", ip).Msg("ipapi fetch failed")
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "failed to fetch IP info",
		})
	}

	return c.JSON(fiber.Map{
		"ipapi": info,
	})
}
