package repositories

import (
	"errors"
	"go-gin/models"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
}

func NewItemMemoryRepository(items []models.Item) ItemRepository {
	return &itemMemoryRepository{items: items}
}

type itemMemoryRepository struct {
	items []models.Item
}

func (i *itemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, item := range i.items {
		if item.ID == itemId {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found")
}

func (i *itemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &i.items, nil
}

func (i *itemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(i.items) + 1)
	i.items = append(i.items, newItem)
	return &newItem, nil
}
