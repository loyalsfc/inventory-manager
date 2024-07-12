package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/loyalsfc/investrite/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file")
	}
	dsn := os.Getenv("DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("fail to connect to the database")
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Order{})

	return db, nil
}
