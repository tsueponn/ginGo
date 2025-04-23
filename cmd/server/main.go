package main

import (
	"fmt"
	"log"
	"os"
	"twitterc/internal/models"
	"twitterc/internal/routes"

	//kpl
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "twitterc/internal/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load("/Users/tsueshima/GolandProjects/twitterc/.env")
	if err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}

	// Build connection string
	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=3000 sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	// Connect to DB with GORM + enable SQL logs
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	// Auto-migrate your models
	err = db.AutoMigrate(&models.User{}, &models.Tweet{})
	if err != nil {
		log.Fatalf("❌ Auto migration failed: %v", err)
	}

	fmt.Println("✅ Connected to DB and models migrated")

	r := gin.Default()

	routes.SetupRoutes(r, db)

	// Start server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
