package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Load environment variables from the specified path
	err := godotenv.Load("/Users/tsueshima/GolandProjects/twitterc/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Build the connection string
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:3000/%s?sslmode=disable", user, password, dbname)

	// Open the database connection
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening DB connection: %v", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL database successfully.")
}
