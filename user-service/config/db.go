package config

import (
	"fmt"
	"log"
	"os"

	"musicplatform/user-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("✅ Database connected")

	// ⬇️ ДОБАВЬ миграцию
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	fmt.Println("✅ Migration completed")
}
