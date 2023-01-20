package models

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
