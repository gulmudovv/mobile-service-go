package main

import (
	"gobackend/database"
	"gobackend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{

		AllowOrigins:     "http://localhost:8080/",
		AllowCredentials: true}))

	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))
}
