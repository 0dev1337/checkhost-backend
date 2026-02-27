package info

import (
	"checkhost-backend/internal/routes/public/info/ipapi"
	"checkhost-backend/internal/routes/public/info/ipinfo"
	"checkhost-backend/pkg/logging"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func Info(c fiber.Ctx) error {
	var host string
	if strings.Contains(c.Host(), "localhost") {
		host = "1.1.1.1"
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
