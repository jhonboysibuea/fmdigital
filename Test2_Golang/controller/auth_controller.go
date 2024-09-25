// controller/auth_controller.go
package controller

import (
	"crud-app/config"
	"crud-app/dto"
	"crud-app/model"
	"crud-app/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// AuthController represents the authentication controller.
type AuthController struct {
	userRepository *repository.UserRepository
}

// NewAuthController creates a new instance of AuthController.
func NewAuthController(userRepository *repository.UserRepository) *AuthController {
	return &AuthController{userRepository: userRepository}
}

// RegisterUser registers a new user.
func (c *AuthController) RegisterUser(req dto.RegisterRequest) error {

	user := &model.User{
		FirstName:   req.FirstName,
		PIN:         string(hashPassword(req.PIN)),
		LastName:    req.LastName,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
	}

	return c.userRepository.CreateUser(user)
}

// Login generates a JWT token and refresh token upon successful login.
func (c *AuthController) Login(username, password string) (*dto.LoginResponse, error) {

	user, err := c.userRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// Placeholder for password validation logic
	err = bcrypt.CompareHashAndPassword([]byte(user.PIN), []byte(password))
	if err != nil {
		return nil, err
	}

	// Create JWT token
	accessToken := generateAccessToken(user)

	// Create Refresh token
	refreshToken, err := generateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	res := &dto.LoginResponse{
		Status: "SUCCESS",
		Login: dto.Login{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	return res, nil
}

// RefreshAccessToken generates a new access token using a refresh token.
func (c *AuthController) RefreshAccessToken(refreshToken string) (string, error) {
	// Verify refresh token
	claims := jwt.MapClaims{}
	refreshTokenClaims, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt().RefreshKey), nil
	})

	if err != nil || !refreshTokenClaims.Valid {
		return "", err
	}

	// Get user details from refresh token
	username := claims["username"].(string)
	user, err := c.userRepository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	// Generate a new access token
	accessToken := generateAccessToken(user)

	return accessToken, nil
}

// Other authentication-related methods here...

// generateAccessToken generates a new access token for the user.
func generateAccessToken(user *model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Second * time.Duration(config.Jwt().Ttl)).Unix(), // Token expiration time (e.g., 1 hour)
	})

	// Sign the token with the secret key
	tokenString, _ := token.SignedString([]byte(config.Jwt().Key))

	return tokenString
}

// generateRefreshToken generates a new refresh token for the user.
func generateRefreshToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // Refresh token expiration time (e.g., 30 days)
	})

	// Sign the token with the refresh token secret key
	refreshTokenString, err := token.SignedString([]byte(config.Jwt().RefreshKey))
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

// // Other authentication-related methods here...

// // // hashPassword is a placeholder for password hashing logic.
func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

// func GenerateDeviceAccessToken(user *model.User, device *model.Device) string {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id":        user.ID,
// 		"tid":       user.Tid,
// 		"device_id": device.ID,
// 		"exp":       time.Now().Add(time.Second * time.Duration(config.Jwt().Ttl)).Unix(), // Token expiration time (e.g., 1 hour)
// 	})

// 	// Sign the token with the secret key
// 	tokenString, _ := token.SignedString([]byte(config.Jwt().Key))

// 	return tokenString
// }

// // generateRefreshToken generates a new refresh token for the user.
// func GenerateDeviceRefreshToken(user *model.User, device *model.Device) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id":        user.ID,
// 		"tid":       user.Tid,
// 		"device_id": device.ID,
// 		"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(), // Refresh token expiration time (e.g., 30 days)
// 	})

// 	// Sign the token with the refresh token secret key
// 	refreshTokenString, err := token.SignedString([]byte(config.Jwt().RefreshKey))
// 	if err != nil {
// 		return "", err
// 	}

// 	return refreshTokenString, nil
// }
