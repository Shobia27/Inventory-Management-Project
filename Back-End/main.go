package main

import (
	"myproject/database"

	"myproject/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()
	app := fiber.New()
	app.Use(cors.New())

	routes.Setup(app)

	app.Listen(":3000")

}
