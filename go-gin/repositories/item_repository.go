package repositories

import (
	"errors"
	"go-gin/models"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
}

func NewItemMemoryRepository(items []models.Item) ItemRepository {
	return &itemMemoryRepository{items: items}
}

type itemMemoryRepository struct {
	items []models.Item
}

func (r *itemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == itemId {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found")
}

func (r *itemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}
