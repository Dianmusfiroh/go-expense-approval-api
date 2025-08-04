package main

import (
	"log"
	"github.com/Dianmusfiroh/go-expense-approval-api/config"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/Dianmusfiroh/go-expense-approval-api/internal/seeder"
	"os"
)

func main() {
	app := fiber.New()
	db := config.ConnectDB()
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}
	if os.Getenv("SEED") == "true" {
		seeder.Seed(db)
	}
	routes.SetupRoutes(app, db)
	// app.Listen(":3000")
	log.Fatal("Failed to start the server: ", app.Listen(":3000"))
}
