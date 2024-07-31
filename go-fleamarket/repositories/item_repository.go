package repositories

import (
	"errors"
	"go-fleamarket/models"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindByID(itemID uint) (*models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

func (i *ItemMemoryRepository) FindByID(itemID uint) (*models.Item, error) {
	for _, item := range i.items {
		if item.ID == itemID {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

func NewItemRepository(items []models.Item) ItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (i *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &i.items, nil
}
