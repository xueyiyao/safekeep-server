package models

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
