package domain

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	ContainerID uint
	CreatedAt   time.Time
}

type ItemService interface {
	FindItemByID(id int) (*Item, error)

	FindItems(filter ItemFilter) ([]*Item, error)

	CreateItem(user *Item) error

	UpdateItem(user *Item) (*Item, error)

	// DeleteItem(id int) error
}

type ItemFilter struct {
	ID           *int    `json:"id"`
	User_ID      *int    `json:"user_id"`
	Container_ID *int    `json:"container_id"`
	Name         *string `json:"name"`
}
