package delivery

import (
	"fmt"
	"net/http"
	"twitterc/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Удаление твита по ID админом
func DeleteTweetByAdmin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tweetID := c.Param("id")
		if err := db.Delete(&models.Tweet{}, tweetID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tweet"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Tweet deleted"})
	}
}

// Блокировка пользователя по ID
func BlockUserByAdmin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		user.IsBlocked = true

		// Печать информации о пользователе перед сохранением
		fmt.Println("Blocking user:", user)

		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to block user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User blocked"})
	}
}
