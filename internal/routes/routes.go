package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"twitterc/internal/delivery"
	"twitterc/internal/repository"
	"twitterc/internal/service"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	as := service.NewAuthService(ur)
	ah := delivery.NewAuthHandler(as)

	RegisterAuthRoutes(r, ah)
}
