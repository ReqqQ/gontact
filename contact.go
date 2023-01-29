package main

import (
	"github.com/gofiber/fiber/v2"
	"gontact/database"
	"gontact/routes"
)

func main() {
	app := fiber.New()
	database.Connect()
	routes.GetRoutes(app)

	app.Listen(":3000")
}
