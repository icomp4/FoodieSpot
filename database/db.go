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
	err := godotenv.Load()
	if err != nil {
		return err
	}
	dbURL := os.Getenv("DB_URL")
	dsn := dbURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	migrate := DB.AutoMigrate(&models.User{}, &models.Location{})
	if migrate != nil {
		log.Println(migrate)
	}
	return nil
}
