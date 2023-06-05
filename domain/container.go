package domain

import (
	"time"

	"gorm.io/gorm"
)

type Container struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	CreatedAt time.Time
}

type ContainerService interface {
	FindContainerByID(id int) (*Container, error)

	FindContainers(user_id int) ([]*Container, error)

	CreateContainer(user *Container) error

	UpdateContainer(user *Container) (*Container, error)

	// DeleteContainer(id int) error
}

type ContainerFilter struct {
	ID      *int `json:"id"`
	User_ID *int `json:"user_id"`
}
