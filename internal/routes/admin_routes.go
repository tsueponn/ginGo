package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"twitterc/internal/delivery"
	"twitterc/internal/middlewares"
)

func RegisterAdminRoutes(r *gin.Engine, db *gorm.DB) {
	admin := r.Group("/admin")
	admin.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware(db))
	{
		admin.DELETE("/tweets/:id", delivery.DeleteTweetByAdmin(db))
		admin.POST("/block/:id", delivery.BlockUserByAdmin(db))
	}
}
