package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"twitterc/internal/delivery"
	"twitterc/internal/middlewares"
)

func RegisterTweetRoutes(r *gin.Engine, tweetHandler *delivery.TweetHandler, db *gorm.DB) {
	r.GET("/tweets/", tweetHandler.ListTweets)
	r.GET("/tweets/:id", tweetHandler.GetTweet)

	tweets := r.Group("/tweets",
		middlewares.AuthMiddleware(),
		middlewares.CheckIfBlockedMiddleware(db),
	)
	{
		tweets.POST("/", tweetHandler.CreateTweet)
		tweets.PUT("/:id", tweetHandler.UpdateTweet)
		tweets.DELETE("/:id", tweetHandler.DeleteTweet)
	}
}
