package dao

import "gorm.io/gorm"

type Dao interface {
	Create(*gorm.Model) (*gorm.DB, error)
	Read(int) (*gorm.Model, error)
	ReadAll() ([]*gorm.Model, error)
	Update(*gorm.Model) (*gorm.DB, error)
	// Delete() error
}
