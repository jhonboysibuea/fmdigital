package controller

import (
	"crud-app/dto"
	"crud-app/model"
	"crud-app/repository"
)

// UserController represents the user controller.
type UserController struct {
	userRepository *repository.UserRepository
}

// NewUserController creates a new instance of UserController.
func NewUserController(userRepository *repository.UserRepository) *UserController {
	return &UserController{userRepository: userRepository}
}

// CreateUser creates a new user.
func (c *UserController) CreateUser(req dto.RegisterRequest ) error {
	hashedPassword := hashPassword(req.PIN)

	user := &model.User{
		FirstName: req.FirstName,
		PIN: hashedPassword,
		LastName:     req.LastName,
		Address: req.Address,
		PhoneNumber: req.PhoneNumber,
	}

	return c.userRepository.CreateUser(user)
}

// GetUserByUsername retrieves a user by username.
func (c *UserController) GetUserByUsername(username string) (*model.User, error) {
	return c.userRepository.GetUserByUsername(username)
}
