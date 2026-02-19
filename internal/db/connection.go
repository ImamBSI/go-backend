package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	username := os.Getenv("MYSQL_DB_USERNAME")
	password := os.Getenv("MYSQL_DB_PASSWORD")
	host := os.Getenv("MYSQL_DB_HOST")
	port := os.Getenv("MYSQL_DB_PORT")
	dbname := os.Getenv("MYSQL_DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	return database
}
