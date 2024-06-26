package main

import (
	"go-admin/database"
	"go-admin/routes"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// 	AllowOrigins:     "https://gradbook-cet.netlify.com",
	// }))

	routes.Setup(app)

	app.Listen(":8000")
}
