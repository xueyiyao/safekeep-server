package repository

import (
	"errors"

	"github.com/xueyiyao/safekeep/domain"
	"github.com/xueyiyao/safekeep/initializers"
	"gorm.io/gorm"
)

type ContainerService struct {
	db *gorm.DB
}

func NewContainerService(db *gorm.DB) *ContainerService {
	return &ContainerService{db: db}
}

func (s *ContainerService) FindContainerByID(id int) (*domain.Container, error) {

	var container domain.Container
	s.db.First(&container, id)

	if container.ID == 0 {
		return nil, errors.New("IdDoesNotExist")
	}

	return &container, nil
}

func (s *ContainerService) FindContainers(user_id int) ([]*domain.Container, error) {
	var containers []*domain.Container
	initializers.DB.Where("user_id <> ?", user_id).Find(&containers)

	// TODO: check for errors

	return containers, nil
}

func (s *ContainerService) CreateContainer(container *domain.Container) error {
	result := initializers.DB.Create(container)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ContainerService) UpdateContainer(container *domain.Container) (*domain.Container, error) {
	result := initializers.DB.Model(container).Updates(*container)

	if result.Error != nil {
		return nil, result.Error
	}

	// TODO: Fix this!
	return &domain.Container{}, nil
}
