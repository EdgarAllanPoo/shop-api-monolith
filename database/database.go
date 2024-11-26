package database

import (
	"log"
	"os"

	"github.com/EdgarAllanPoo/shop-api-monolith/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL is not set in the environment")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	log.Println("Database connected successfully")
}
