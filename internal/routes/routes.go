package routes

import (
	"checkhost-backend/internal/routes/public/info"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	publicGroup := app.Group("/api/v1/public")
	publicGroup.Get("/info", info.Info)
}
