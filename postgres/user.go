package postgres

import (
	"errors"
	"net/mail"
	"strings"

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
	if user == nil {
		return errors.New("NilUser")
	}

	if len(strings.TrimSpace(user.Name)) == 0 {
		return errors.New("EmptyUserName")
	}

	if len(strings.TrimSpace(user.Email)) == 0 {
		return errors.New("EmptyEmail")
	}

	if _, err := mail.ParseAddress(strings.TrimSpace(user.Email)); err != nil {
		return err
	}

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
