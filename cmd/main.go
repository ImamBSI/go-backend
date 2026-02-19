package main

import (
	"encoding/json"
	"log"
	"os"

	"example.com/trial-go/internal/auth"
	"example.com/trial-go/internal/db"
	"example.com/trial-go/internal/energy"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Energy JSON mock
	file, err := os.ReadFile("data_example.json")
	if err != nil {
		log.Fatal(err)
	}

	var response energy.EnergyResponse
	if err := json.Unmarshal(file, &response); err != nil {
		log.Fatal(err)
	}

	energyService := energy.NewService(response.Data)
	energyHandler := energy.NewHandler(energyService)

	// DB
	database := db.NewConnection()

	// Auth setup
	authRepo := auth.NewRepository(database)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/go-be")

	// Energy routes
	api.Get("/raw-energy", energyHandler.GetRawEnergy)
	api.Get("/sum-energy", energyHandler.GetSumEnergy)

	// Auth routes
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	log.Fatal(app.Listen(":3000"))
}
