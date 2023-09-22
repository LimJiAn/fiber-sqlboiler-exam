package main

import (
	"log"

	"github.com/LimJiAn/go-sqlboiler-exam/api/route"
	"github.com/LimJiAn/go-sqlboiler-exam/database"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()
	// logger
	app.Use(logger.New())
	route.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
