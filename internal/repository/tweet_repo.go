package repository

import (
	"twitterc/internal/models"

	"gorm.io/gorm"
)

type TweetRepository struct {
	DB *gorm.DB
}

func NewTweetRepository(db *gorm.DB) *TweetRepository {
	return &TweetRepository{DB: db}
}

func (r *TweetRepository) CreateTweet(tweet *models.Tweet) error {
	return r.DB.Create(tweet).Error
}

func (r *TweetRepository) GetAllTweets() ([]models.Tweet, error) {
	var tweets []models.Tweet
	err := r.DB.Find(&tweets).Error
	return tweets, err
}
