package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID          string `gorm:"type:varchar(36);primarykey" json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	PIN         string `json:"pin"` // You might want to hash this for security
}

func (o *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	o.ID = uuid.NewString()
	return
}
