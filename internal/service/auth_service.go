package service

import (
	"errors"
	"time"
	"twitterc/internal/models"
	"twitterc/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte("your_secret_key") // TODO: move to config or env later

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
	}
}

// Register a new user with hashed password
func (s *AuthService) Register(name, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: name,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login user and return JWT token
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Create JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // token expires in 72h
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
