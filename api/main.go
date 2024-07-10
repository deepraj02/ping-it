package main

import (
	"github.com/deepraj02/pingit/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)

	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	//TODO: Refactor hardcoded string to env variable
	app.Listen(":3000")
}
