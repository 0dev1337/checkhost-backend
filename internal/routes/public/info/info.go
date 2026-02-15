package info

import (
	"checkhost-backend/internal/routes/public/info/ipapi"
	"checkhost-backend/internal/routes/public/info/ipinfo"
	"checkhost-backend/pkg/logging"

	"github.com/gofiber/fiber/v3"
)

func Info(c fiber.Ctx) error {

	host := c.Query("host")
	if host == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": "Missing 'host' query parameter",
		})
	}
	ipinfo, err := ipinfo.GetIPInfo(host)
	if err != nil {
		logging.Logger.Error().Msg(err.Error())

	}

	ipapi, err := ipapi.GetIPAPIInfo(host)
	if err != nil {
		logging.Logger.Error().Msg(err.Error())
	}
	return c.JSON(map[string]interface{}{
		"ipinfo": ipinfo,
		"ipapi":  ipapi,
	})
}
