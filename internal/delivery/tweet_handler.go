package delivery

import (
	"net/http"
	"strconv"
	"twitterc/internal/service"

	"github.com/gin-gonic/gin"
)

type TweetHandler struct {
	TweetService *service.TweetService
}

func NewTweetHandler(tweetService *service.TweetService) *TweetHandler {
	return &TweetHandler{TweetService: tweetService}
}

func (h *TweetHandler) CreateTweet(c *gin.Context) {
	var input struct {
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract userID from context (set by AuthMiddleware)
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	// Create tweet using userID
	tweet, err := h.TweetService.CreateTweet(input.Content, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tweet"})
		return
	}

	c.JSON(http.StatusCreated, tweet)
}

func (h *TweetHandler) ListTweets(c *gin.Context) {
	tweets, err := h.TweetService.GetAllTweets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweets"})
		return
	}

	c.JSON(http.StatusOK, tweets)
}

func (h *TweetHandler) GetTweet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tweet ID"})
		return
	}

	tweet, err := h.TweetService.GetTweetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
		return
	}

	c.JSON(http.StatusOK, tweet)
}

func (h *TweetHandler) UpdateTweet(c *gin.Context) {
	// Get tweet ID from URL parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tweet ID"})
		return
	}

	// Extract userID from context (set by AuthMiddleware)
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	// Bind input data
	var input struct {
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the tweet, passing the userID to ensure the correct user is updating the tweet
	err = h.TweetService.UpdateTweet(uint(id), input.Content, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tweet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tweet updated successfully"})
}

func (h *TweetHandler) DeleteTweet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tweet ID"})
		return
	}

	err = h.TweetService.DeleteTweet(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tweet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tweet deleted successfully"})
}
