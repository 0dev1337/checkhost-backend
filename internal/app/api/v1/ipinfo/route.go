package ipinfo

import (
	"checkhost-backend/internal/client/ipapi"
	"checkhost-backend/internal/httpx"
	"checkhost-backend/internal/logging"

	"github.com/gofiber/fiber/v3"
)

func GET(c fiber.Ctx) error {
	host := httpx.TargetHost(c)

	info, err := ipapi.Fetch(host)
	if err != nil {
		logging.Logger.Error().Err(err).Str("ip", host).Msg("ipapi fetch failed")
		return httpx.Error(c, fiber.StatusBadGateway, "failed to fetch IP info")
	}

	return c.JSON(fiber.Map{"ipapi": info})
}
