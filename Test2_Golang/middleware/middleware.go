package middleware

import (
	"crud-app/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// AuthMiddleware is a middleware that checks for authentication.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := extractUserIDFromToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
			c.Abort()
			return
		}

		// Store the user ID in the Gin context
		c.Set("user_id", userID)

		c.Next()
	}
}

// Mocked IsValidToken function, replace with actual token validation logic
func IsValidToken(c *gin.Context) bool {
	// Example: Check for a Bearer token in the Authorization header
	header := c.GetHeader("Authorization")
	if header == "" {
		return false
	}

	// Check if the Authorization header starts with "Bearer "
	const prefix = "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return false
	}

	// Extract the token from the header
	tokenString := strings.TrimPrefix(header, prefix)

	// Validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Replace with your own secret key or public key for verification
		return []byte(config.Jwt().Key), nil
	})

	return err == nil && token.Valid
}
func extractUserIDFromToken(c *gin.Context) (string, error) {
	// Example: Check for a Bearer token in the Authorization header
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", fmt.Errorf("no token provided")
	}

	// Check if the Authorization header starts with "Bearer "
	const prefix = "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return "", fmt.Errorf("invalid token format")
	}

	// Extract the token from the header
	tokenString := strings.TrimPrefix(header, prefix)

	// Validate the JWT token and extract user ID
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt().Key), nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Extract user ID from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["id"].(string)
	if !ok {
		return "", fmt.Errorf("id not found in token claims")
	}

	return userID, nil
}
