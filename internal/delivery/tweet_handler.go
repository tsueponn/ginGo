package delivery

import (
	"net/http"
	"twitterc/internal/models"
	"twitterc/internal/repository"

	"github.com/gin-gonic/gin"
)

type TweetHandler struct {
	TweetRepo *repository.TweetRepository
}

func NewTweetHandler(tweetRepo *repository.TweetRepository) *TweetHandler {
	return &TweetHandler{
		TweetRepo: tweetRepo,
	}
}

// CreateTweet handles creating a new tweet
func (h *TweetHandler) CreateTweet(c *gin.Context) {
	var tweet models.Tweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.TweetRepo.CreateTweet(&tweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tweet"})
		return
	}

	c.JSON(http.StatusCreated, tweet)
}

// ListTweets handles fetching all tweets
func (h *TweetHandler) ListTweets(c *gin.Context) {
	tweets, err := h.TweetRepo.GetAllTweets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweets"})
		return
	}

	c.JSON(http.StatusOK, tweets)
}
