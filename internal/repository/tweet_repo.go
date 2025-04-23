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

// CreateTweet creates a new tweet
func (r *TweetRepository) CreateTweet(tweet *models.Tweet) error {
	return r.DB.Create(tweet).Error
}

// GetAllTweets fetches all tweets
func (r *TweetRepository) GetAllTweets() ([]models.Tweet, error) {
	var tweets []models.Tweet
	err := r.DB.Find(&tweets).Error
	return tweets, err
}

// GetTweetByID fetches a single tweet by ID
func (r *TweetRepository) GetTweetByID(id uint) (*models.Tweet, error) {
	var tweet models.Tweet
	err := r.DB.First(&tweet, id).Error
	if err != nil {
		return nil, err
	}
	return &tweet, nil
}

// UpdateTweet updates the content of a tweet
func (r *TweetRepository) UpdateTweet(id uint, newContent string) error {
	return r.DB.Model(&models.Tweet{}).Where("id = ?", id).Update("content", newContent).Error
}

// DeleteTweet deletes a tweet by ID
func (r *TweetRepository) DeleteTweet(id uint) error {
	return r.DB.Delete(&models.Tweet{}, id).Error
}
