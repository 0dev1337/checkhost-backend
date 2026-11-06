package ping

import (
	"checkhost-backend/internal/client/worker"
	"checkhost-backend/internal/httpx"
	"checkhost-backend/internal/logging"

	"github.com/gofiber/fiber/v3"
)

const attempts = 4

type response struct {
	IP       string  `json:"ip"`
	Success  int32   `json:"success"`
	Attempts int32   `json:"attempts"`
	RTT      []int32 `json:"rtt_ms"`
	Location string  `json:"location"`
}

func GET(workers *worker.Client) fiber.Handler {
	return func(c fiber.Ctx) error {
		host := c.Query("host")
		node := c.Query("node")

		res := response{IP: host, Attempts: attempts}
		for range attempts {
			result, err := workers.Ping(host, node)
			if err != nil {
				logging.Logger.Error().Str("host", host).Err(err).Msg("ping attempt failed")
				continue
			}

			logging.Logger.Info().Str("host", host).Msg("ping attempt succeeded")
			res.Success++
			res.RTT = append(res.RTT, result.RTT)
			res.Location = result.Location
		}

		if res.Success == 0 {
			return httpx.Error(c, fiber.StatusBadGateway, "all ping attempts failed")
		}

		return c.JSON(res)
	}
}
