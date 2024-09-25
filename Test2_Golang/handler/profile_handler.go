package handler

import (
	"crud-app/controller"
	"crud-app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProfile(router *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	controller := controller.NewProfileController(repo)
	profileHandler := NewProfileHandler(controller)

	router.PUT("/profile", profileHandler.UpdateProfileHandler)
}

// ProfileHandler handles HTTP requests for partners
type ProfileHandler struct {
	controller *controller.ProfileController
}

// NewProfileHandler creates a new ProfileHandler with the provided controller
func NewProfileHandler(controller *controller.ProfileController) *ProfileHandler {
	return &ProfileHandler{controller: controller}
}

// UpdateProfileHandler handles the PUT request to update a partner
func (h *ProfileHandler) UpdateProfileHandler(c *gin.Context) {
	// var id uint
	var id = c.GetString("user_id")
	user, err := h.controller.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.controller.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
