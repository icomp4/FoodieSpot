package database

import (
	"foodSharer/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func StartDB() error {
	// Loading env vars
	err := godotenv.Load()
	if err != nil {
		return err
	}
	// Secret DB url
	dbURL := os.Getenv("DB_URL")
	dsn := dbURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Updating global DB var, similar to static in java
	DB = db

	// Migrating tables to ensure database is up-to-date
	migrate := DB.AutoMigrate(&models.User{}, &models.Location{})
	if migrate != nil {
		log.Println(migrate)
	}
	return nil
}
