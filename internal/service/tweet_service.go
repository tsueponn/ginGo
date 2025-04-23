package service

import (
	"errors"
	"twitterc/internal/models"
	"twitterc/internal/repository"
)

type TweetService struct {
	tweetRepo *repository.TweetRepository
}

func NewTweetService(tweetRepo *repository.TweetRepository) *TweetService {
	return &TweetService{tweetRepo: tweetRepo}
}

func (s *TweetService) CreateTweet(content string, userID uint) (*models.Tweet, error) {
	tweet := &models.Tweet{
		Content: content,
		UserID:  userID,
	}
	err := s.tweetRepo.CreateTweet(tweet)
	return tweet, err
}

func (s *TweetService) GetAllTweets() ([]models.Tweet, error) {
	return s.tweetRepo.GetAllTweets()
}

func (s *TweetService) GetTweetByID(id uint) (*models.Tweet, error) {
	return s.tweetRepo.GetTweetByID(id)
}

func (s *TweetService) UpdateTweet(id uint, newContent string, userID uint) error {
	// Fetch the tweet by ID
	tweet, err := s.tweetRepo.GetTweetByID(id)
	if err != nil {
		return err // Tweet not found
	}

	// Validate that the userID matches the tweet's UserID
	if tweet.UserID != userID {
		return errors.New("unauthorized: cannot update another user's tweet")
	}

	// Update the tweet's content
	tweet.Content = newContent
	return s.tweetRepo.UpdateTweet(id, tweet.Content)
}

func (s *TweetService) DeleteTweet(id uint) error {
	return s.tweetRepo.DeleteTweet(id)
}
