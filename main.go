package main

import (
	"learn-go-fiber/database"
	"learn-go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initial database
	database.DatabaseInit()
	
	app := fiber.New()

	// Initialize route
	route.RouteInit(app)

	app.Listen(":3000")
}
