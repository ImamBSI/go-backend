package main

import (
	"encoding/json"
	"log"
	"os"

	"example.com/trial-go/handlers"
	"example.com/trial-go/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// SIMULASI: Membaca file data_example.json
	file, _ := os.ReadFile("data_example.json")
	var response models.EnergyResponse
	json.Unmarshal(file, &response)

	// Inject data ke handler (Simulasi Database)
	handlers.EnergyData = response.Data

	app := fiber.New()

	// Membuat Group dengan prefix /go-be
	api := app.Group("/go-be")

	// Sekarang rutenya menjadi: /go-be/raw-energy
	api.Get("/raw-energy", handlers.GetRawEnergy)

	log.Fatal(app.Listen(":3000"))
}
