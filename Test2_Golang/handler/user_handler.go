package handler

import (
	"crud-app/controller"
	"crud-app/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler represents the user handler.
type UserHandler struct {
	userController *controller.UserController
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(userController *controller.UserController) *UserHandler {
	return &UserHandler{userController: userController}
}

// CreateUserHandler handles HTTP requests to create a new user.
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	// Extract data from the request, e.g., using c.PostForm or c.BindJSON
	var req dto.RegisterRequest
	c.Bind(&req)

	// Call the CreateUser method from UserController
	err := h.userController.CreateUser(req)
	if err != nil {
		// Handle error, return appropriate Gin response, log, etc.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// User created successfully, return success response
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Other user-related handler methods here...
