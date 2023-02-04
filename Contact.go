package main

import (
	"github.com/gofiber/fiber/v2"
	AppDatabase "gontact/App/Database"
	"gontact/UI"
)

func main() {
	app := fiber.New()
	AppDatabase.Connect()
	UI.GetRoutes(app)
	UI.GetPostRoutes()
	app.Listen(":3000")
}
