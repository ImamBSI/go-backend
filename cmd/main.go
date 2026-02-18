package main

import (
	"encoding/json"
	"log"
	"os"

	"example.com/trial-go/internal/energy"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load JSON file (simulasi database)
	file, err := os.ReadFile("data_example.json")
	if err != nil {
		log.Fatal(err)
	}

	var response energy.EnergyResponse
	if err := json.Unmarshal(file, &response); err != nil {
		log.Fatal(err)
	}

	// Dependency injection
	service := energy.NewService(response.Data)
	handler := energy.NewHandler(service)

	app := fiber.New()

	api := app.Group("/go-be")

	api.Get("/raw-energy", handler.GetRawEnergy)
	api.Get("/sum-energy", handler.GetSumEnergy)

	log.Fatal(app.Listen(":3000"))
}
