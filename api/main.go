package main

import (
	"log"

	"github.com/deepraj02/pingit/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1/shorten", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	//TODO: Refactor hardcoded string to env variable
	app.Listen(":3000")
}
