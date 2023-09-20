package main

import (
	"log"

	"github.com/LimJiAn/go-sqlboiler-exam/database"
	"github.com/LimJiAn/go-sqlboiler-exam/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
