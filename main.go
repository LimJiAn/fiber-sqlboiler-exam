package main

import (
	"log"

	"github.com/LimJiAn/go-sqlboiler-exam/api/route"
	"github.com/LimJiAn/go-sqlboiler-exam/database"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()
	route.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
