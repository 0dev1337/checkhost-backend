package main

import (
	"log"

	"checkhost-backend/internal/app"
	"checkhost-backend/internal/config"
	"checkhost-backend/internal/logging"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
)

const listenAddr = ":3001"

func main() {
	cfg := config.Load()
	if len(cfg.WorkerNodes) == 0 {
		logging.Logger.Warn().Msg("WORKER_NODES is empty, ping and tcp requests will be rejected")
	}
	if cfg.WorkerAPIKey == "" {
		logging.Logger.Warn().Msg("WORKER_API_KEY is empty, workers will reject probe requests")
	}

	server := fiber.New()

	server.Use(recoverer.New())
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	app.RegisterRoutes(server, cfg)
	log.Fatal(server.Listen(listenAddr))
}
