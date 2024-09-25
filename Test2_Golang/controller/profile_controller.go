package controller

import (
	"crud-app/model"
	"crud-app/repository"
)

// ProfileController handles business logic for partners
type ProfileController struct {
	repo *repository.UserRepository
}

// NewUserController creates a new instance of PartnerController.
func NewProfileController(merchantRepository *repository.UserRepository) *ProfileController {
	return &ProfileController{repo: merchantRepository}
}

// UpdateUser
func (c *ProfileController) UpdateUser(merchant *model.User) error {
	return c.repo.UpdateUser(merchant)
}

func (c *ProfileController) GetUserByID(id string) (*model.User, error) {
	return c.repo.GetUserByID(id)
}
