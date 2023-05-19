package logic

import (
	"errors"

	"github.com/xueyiyao/safekeep/dao"
	"github.com/xueyiyao/safekeep/models"
	"gorm.io/gorm"
)

func CreateUser(user *models.User) (*gorm.DB, error) {
	userDao := dao.UserDao{}
	result, err := userDao.Create(user)

	if err != nil {
		return nil, errors.New("CouldNotSaveToDB")
	}

	return result, nil
}

func ReadUser(id int) (*models.User, error) {
	userDao := dao.UserDao{}
	user, err := userDao.Read(id)

	if err != nil {
		return nil, errors.New("IdDoesNotExist")
	}

	return user, nil
}

func UpdateUser(user *models.User) error {
	userDao := dao.UserDao{}
	err := userDao.Update(user)

	if err != nil {
		return errors.New("CouldNotUpdateToDB")
	}

	return nil
}
