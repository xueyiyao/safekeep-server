package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	CreatedAt time.Time
}

type UserService interface {
	FindUserByID(id int) (*User, error)

	CreateUser(user *User) error

	UpdateUser(user *User) (*User, error)

	// DeleteUser(id int) error
}
