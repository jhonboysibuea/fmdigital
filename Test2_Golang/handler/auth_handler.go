package handler

import (
	"crud-app/controller"
	"crud-app/dto"
	"crud-app/logger"
	"crud-app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AuthHandler represents the authentication handler.
type AuthHandler struct {
	authController *controller.AuthController
}

func SetupAuthRoutes(router *gin.RouterGroup, db *gorm.DB) {

	userRepository := repository.NewUserRepository(db)

	authController := controller.NewAuthController(userRepository)

	authHandler := NewAuthHandler(authController)

	router.POST("/register", authHandler.RegisterUserHandler)
	router.POST("/login", authHandler.LoginHandler)
	// router.POST("/refresh-token", authHandler.RefreshTokenHandler)
	// router.POST("/device/login", authHandler.LoginDeviceHandler)
}

// NewAuthHandler creates a new instance of AuthHandler.
func NewAuthHandler(authController *controller.AuthController) *AuthHandler {
	return &AuthHandler{authController: authController}
}

// RegisterUserHandler handles HTTP requests to register a new user.
func (h *AuthHandler) RegisterUserHandler(c *gin.Context) {
	var req dto.RegisterRequest
	c.Bind(&req)

	// Call the CreateUser method from UserController

	// Call the RegisterUser method from AuthController
	err := h.authController.RegisterUser(req)
	if err != nil {
		// Handle error, return appropriate Gin response, log, etc.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// User registered successfully, return success response
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginHandler handles HTTP requests for user login.
func (h *AuthHandler) LoginHandler(c *gin.Context) {
	// Extract data from the request, e.g., using c.PostForm or c.BindJSON
	var req dto.AuthRequest
	c.Bind(&req)
	logger.FromContext(c.Request.Context()).Info("Handling login ")
	// Call the Login method from AuthController
	res, err := h.authController.Login(req.PhoneNumber, req.PIN)
	if err != nil {
		// Handle error, return appropriate Gin response, log, etc.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
		return
	}

	// Login successful, return the access token and refresh token
	c.JSON(http.StatusOK, res)
}
