package tcp

import (
	"net"
	"strconv"

	"checkhost-backend/internal/client/worker"
	"checkhost-backend/internal/httpx"
	"checkhost-backend/internal/logging"

	"github.com/gofiber/fiber/v3"
)

func GET(workers *worker.Client) fiber.Handler {
	return func(c fiber.Ctx) error {
		host := c.Query("host")
		port := c.Query("port")

		if !validPort(port) {
			return httpx.Error(c, fiber.StatusBadRequest, "port must be a number between 1 and 65535")
		}

		result, err := workers.TCP(net.JoinHostPort(host, port), c.Query("node"))
		if err != nil {
			logging.Logger.Error().Str("host", host).Err(err).Msg("tcp probe attempt failed")
			return httpx.Error(c, fiber.StatusBadGateway, "tcp probe failed")
		}

		return c.JSON(result)
	}
}

func validPort(port string) bool {
	n, err := strconv.ParseUint(port, 10, 16)
	return err == nil && n != 0
}
