package main

import (
	"app/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	// Static file server
	app.Static("/", "./public")

	// Default middleware config
	app.Use(recover.New())

	// SetupRoutes
	router.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))

}
