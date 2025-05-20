package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"twitterc/internal/delivery"
	"twitterc/internal/repository"
	"twitterc/internal/service"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Auth
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := delivery.NewAuthHandler(authService)

	// Tweets
	tweetRepo := repository.NewTweetRepository(db)
	tweetService := service.NewTweetService(tweetRepo)
	tweetHandler := delivery.NewTweetHandler(tweetService)

	// Register routes
	RegisterAuthRoutes(r, authHandler)
	RegisterTweetRoutes(r, tweetHandler, db) // <--- передаём db
	RegisterAdminRoutes(r, db)
}
