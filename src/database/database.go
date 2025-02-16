package database

import (
	"agoravote-app-backend/src/models" // Update with your actual project path
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

	// Auto-migrate the Group, Vote, and Post models
	log.Println("Auto-migrating the Group, Vote, and Post models")
	db.AutoMigrate(&models.Group{}, &models.Vote{}, &models.Post{})

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

	// Populate the database with sample data
	populateGroups()
}

func populateGroups() {
	log.Println("Populating groups with sample data")

	group := models.Group{
		ID:          "1",
		Name:        "Group 1",
		Description: "Description for Group 1",
		IsPrivate:   false,
		LastActive:  "2025-02-16",
	}

	log.Printf("Inserting group: %s", group.Name)
	if err := DB.FirstOrCreate(&group, models.Group{ID: group.ID}).Error; err != nil {
		log.Printf("Failed to insert group %s: %v", group.Name, err)
	} else {
		log.Printf("Successfully inserted group %s", group.Name)
	}
}
