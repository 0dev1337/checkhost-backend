package routes

import (
	"checkhost-backend/internal/handler"
	"checkhost-backend/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Use(middleware.Logging)
	api.Get("/ipinfo", handler.Info)
	api.Get("/whois", handler.Whois)
}
