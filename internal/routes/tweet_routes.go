package routes

import (
	"github.com/gin-gonic/gin"
	"twitterc/internal/delivery"
	"twitterc/internal/middlewares"
)

func RegisterTweetRoutes(r *gin.Engine, tweetHandler *delivery.TweetHandler) {

	r.GET("/tweets/", tweetHandler.ListTweets)
	r.GET("/tweets/:id", tweetHandler.GetTweet)
	tweets := r.Group("/tweets", middlewares.AuthMiddleware())
	{
		tweets.POST("/", tweetHandler.CreateTweet)
		tweets.PUT("/:id", tweetHandler.UpdateTweet)
		tweets.DELETE("/:id", tweetHandler.DeleteTweet)
	}
}
