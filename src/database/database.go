package database

import (
	"agoravote-app-backend/src/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	log.Printf("Connecting to database with DSN: %s", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the Group, Vote, Post, and GroupMember models
	log.Println("Auto-migrating the Group, Vote, Post, and GroupMember models")
	db.AutoMigrate(&models.Group{}, &models.Vote{}, &models.Post{}, &models.GroupMember{})

	// Verify the connection by running a simple query
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Successfully connected to the database")

	DB = db
}
