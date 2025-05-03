package main

import (
	"fmt"
	"log"
	"twitterc/internal/db"
	"twitterc/internal/routes"

	//kpl
	"github.com/gin-gonic/gin"
	_ "twitterc/internal/routes"
)

func main() {
	// Load environment variables
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	db.RunMigrationsUp()

	fmt.Println("✅ Connected to DB and models migrated")

	r := gin.Default()

	routes.SetupRoutes(r, database)

	// Start server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
