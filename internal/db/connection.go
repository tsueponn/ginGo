package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Load environment variables
	err := godotenv.Load("/Users/tsueshima/GolandProjects/twitterc/.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=3000 sslmode=disable", user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %w", err)
	}

	return db, nil
}
