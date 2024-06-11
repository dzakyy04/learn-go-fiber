package main

import (
	"learn-go-fiber/database"
	"learn-go-fiber/migration"
	"learn-go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database
	database.DatabaseInit()
	// Run migration
	migration.RunMigration()

	app := fiber.New()

	// Initialize route
	route.RouteInit(app)

	app.Listen(":3000")
}
