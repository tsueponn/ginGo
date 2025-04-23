package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Validate token and extract user ID
		userID, err := ValidateToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Save user ID to context
		c.Set("userID", userID)

		c.Next()
	}
}

var JWTSecret = []byte("your_secret_key")

// ValidateToken validates the JWT and returns the user ID if the token is valid
func ValidateToken(tokenString string) (uint, error) {
	// Split token to get the Bearer part
	parts := strings.SplitN(tokenString, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, errors.New("invalid token format")
	}

	// Parse token
	token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JWTSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid or expired token")
	}

	// Extract user ID from token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid claims in token")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id not found in token claims")
	}

	return uint(userID), nil
}
