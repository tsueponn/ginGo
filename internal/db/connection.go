package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load environment variables
	err := godotenv.Load("/Users/tsueshima/GolandProjects/twitterc/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=3000 sslmode=disable", user, password, dbname)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL database successfully.")
}
