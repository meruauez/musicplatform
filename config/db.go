package config

import (
	"fmt"
	"log"
	"musicplatform/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgre17 dbname=musicplatform port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Автоматическая миграция для всех моделей
	DB.AutoMigrate(&models.Artist{}, &models.Genre{}, &models.Song{})

	fmt.Println("✅ Database connected and migrated")

	DB.AutoMigrate(&models.User{})

}
