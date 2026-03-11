package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	auth "example.com/trial-go/internal/auth"
	"example.com/trial-go/internal/energy"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// =========================
	// DATABASE CONNECTION
	// =========================
	username := os.Getenv("MYSQL_DB_USERNAME")
	password := os.Getenv("MYSQL_DB_PASSWORD")
	host := os.Getenv("MYSQL_DB_HOST")
	port := os.Getenv("MYSQL_DB_PORT")
	dbname := os.Getenv("MYSQL_DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// Auto migrate
	if err := db.AutoMigrate(&auth.User{}, &auth.Account{}); err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	// =========================
	// AUTH MODULE INIT
	// =========================
	authRepo := &auth.Repository{Db: db}
	authService := &auth.Service{Repo: authRepo}
	authHandler := &auth.Handler{Service: authService}

	// =========================
	// ENERGY MOCK (existing)
	// =========================
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

	// =========================
	// FIBER SETUP
	// =========================
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/go-be")

	// Energy routes
	api.Get("/raw-energy", energyHandler.GetRawEnergy)
	api.Get("/sum-energy", energyHandler.GetSumEnergy)

	// Auth routes
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)
	api.Get("/users", authHandler.GetUsers)
	api.Delete("/users/:id", authHandler.DeleteUser)

	log.Fatal(app.Listen(":3000"))
}
