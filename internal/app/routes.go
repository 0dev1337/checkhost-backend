package app

import (
	"checkhost-backend/internal/app/api/v1/ipinfo"
	"checkhost-backend/internal/app/api/v1/ping"
	"checkhost-backend/internal/app/api/v1/tcp"
	"checkhost-backend/internal/app/api/v1/whois"
	"checkhost-backend/internal/client/worker"
	"checkhost-backend/internal/config"
	"checkhost-backend/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(server *fiber.App, cfg config.Config) {
	workers := worker.New(cfg.WorkerAPIKey)
	sanitize := middleware.Sanitize(cfg.WorkerNodes)

	v1 := server.Group("/api/v1")
	v1.Use(middleware.Logging)

	v1.Get("/ipinfo", ipinfo.GET)
	v1.Get("/whois", whois.GET)
	v1.Get("/ping", sanitize, ping.GET(workers))
	v1.Get("/tcp", sanitize, tcp.GET(workers))
}
