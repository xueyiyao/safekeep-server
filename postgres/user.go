package postgres

import (
	"errors"

	"github.com/xueyiyao/safekeep/domain"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) FindUserByID(id int) (*domain.User, error) {
	var user domain.User
	s.db.First(&user, id)

	if user.ID == 0 {
		return nil, errors.New("IdDoesNotExist")
	}

	return &user, nil
}

func (s *UserService) CreateUser(user *domain.User) error {
	result := s.db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *UserService) UpdateUser(user *domain.User) (*domain.User, error) {
	result := s.db.Model(user).Updates(*user)

	if result.Error != nil {
		return nil, result.Error
	}

	// TODO: Fix this!
	return &domain.User{}, nil
}
