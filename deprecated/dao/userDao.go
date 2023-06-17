package dao

import (
	"errors"

	"github.com/xueyiyao/safekeep/deprecated/initializers"
	"github.com/xueyiyao/safekeep/deprecated/models"
	"gorm.io/gorm"
)

type UserDao struct {
}

func (dao *UserDao) Create(user *models.User) (*gorm.DB, error) {
	result := initializers.DB.Create(user)

	if result.Error != nil {
		return nil, errors.New("CouldNotSaveToDB")
	}

	return result, nil
}

func (dao *UserDao) Read(id int) (*models.User, error) {
	var user models.User
	initializers.DB.First(&user, id)

	if user.ID == 0 {
		return nil, errors.New("IdDoesNotExist")
	}

	return &user, nil
}

func (dao *UserDao) Update(user *models.User) error {
	result := initializers.DB.Model(user).Updates(*user)

	if result.Error != nil {
		return errors.New("CouldNotUpdateToDB")
	}

	return nil
}
