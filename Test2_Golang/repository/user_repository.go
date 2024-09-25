// repository/user_repository.go
package repository

import (
	"crud-app/model"

	"gorm.io/gorm"
)

// UserRepository represents the user repository.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser creates a new user.
func (r *UserRepository) CreateUser(user *model.User) error {
	result := r.db.Create(user)
	return result.Error
}

// GetUserByUsername retrieves a user by username.
func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.db.Where("phone_number = ?", username).First(&user)
	return &user, result.Error
}

// CreateUser creates a new user.
func (r *UserRepository) UpdateUser(user *model.User) error {
	result := r.db.Save(user)
	return result.Error
}

func (r *UserRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	result := r.db.Where("id = ?", id).First(&user)
	return &user, result.Error
}
