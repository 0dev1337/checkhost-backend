package main

import (
	"checkhost-backend/internal/routes"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()

	// cors should be configured before setting up routes and use New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3001"))
}
