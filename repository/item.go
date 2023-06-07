package repository

import (
	"errors"
	"strings"

	"github.com/xueyiyao/safekeep/domain"
	"gorm.io/gorm"
)

type ItemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{db: db}
}

func (s *ItemService) FindItemByID(id int) (*domain.Item, error) {
	var item domain.Item
	s.db.First(&item, id)

	if item.ID == 0 {
		return nil, errors.New("IdDoesNotExist")
	}

	return &item, nil
}

func (s *ItemService) FindItems(filter domain.ItemFilter) ([]*domain.Item, error) {
	where, args := []string{}, []interface{}{}
	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}
	if v := filter.User_ID; v != nil {
		where, args = append(where, "user_id = ?"), append(args, *v)
	}
	if v := filter.Container_ID; v != nil {
		where, args = append(where, "container_id = ?"), append(args, *v)
	}
	if v := filter.Name; v != nil {
		where, args = append(where, "name = ?"), append(args, *v)
	}

	queryStr := `SELECT * FROM items WHERE ` + strings.Join(where, " AND ") + ``

	var items []*domain.Item
	result := s.db.Raw(queryStr, args...).Scan(&items)

	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (s *ItemService) CreateItem(item *domain.Item) error {
	result := s.db.Create(item)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ItemService) UpdateItem(item *domain.Item) (*domain.Item, error) {
	result := s.db.Model(item).Updates(*item)

	if result.Error != nil {
		return nil, result.Error
	}

	// TODO: Fix this!
	return &domain.Item{}, nil
}
