package database

import (
	"agoravote-app-backend/src/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Print environment variables for debugging
	log.Printf("DB_HOST: %s, DB_USER: %s, DB_PASSWORD: %s, DB_NAME: %s, DB_PORT: %s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Print DSN for debugging
	log.Printf("DSN: %s", dsn)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{})

	DB = db
}
