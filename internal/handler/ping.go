package handler

import (
	"checkhost-backend/internal/helpers"
	"checkhost-backend/pkg/logging"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type PingResponse struct {
	IP       string  `json:"ip"`
	Success  int32   `json:"success"`
	Attempts int32   `json:"attempts"`
	RTT      []int32 `json:"rtt_ms"`
	Location string  `json:"location"`
}

func Ping(c fiber.Ctx) error {
	host := c.Query("host")
	if host == "" {
		host = c.IP()
	}
	// Handle localhost and loopback addresses
	if strings.Contains(host, "localhost") || strings.Contains(host, "127.0.0.1") {
		host = "1.1.1.1"
	}

	var response PingResponse
	for i := 0; i < 4; i++ {
		response.Attempts++
		result, err := helpers.PingICMP(host)
		if err != nil {
			logging.Logger.Error().Str("host", host).Err(err).Msg("ping attempt failed")
			continue
		}

		logging.Logger.Info().Str("host", host).Msg("ping attempt succeeded")
		response.Success++
		response.RTT = append(response.RTT, result.RTT)
		response.Location = result.Location
	}

	response.IP = host

	if response.Success == 0 {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "all ping attempts failed",
		})
	}

	return c.JSON(response)
}
