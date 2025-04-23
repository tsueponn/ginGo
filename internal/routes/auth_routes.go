package routes

import (
	"github.com/gin-gonic/gin"
	"twitterc/internal/delivery"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler *delivery.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
}
