package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	//TODO: SetUp Routes
}

func main() {
	app := fiber.New()

	app.Use(logger.New())
}
