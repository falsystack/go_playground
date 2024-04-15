package repositories

import (
	"errors"
	"go-gin/models"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint) error
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

func (i *itemMemoryRepository) Update(updateItem models.Item) (*models.Item, error) {
	for idx, v := range i.items {
		if v.ID == updateItem.ID {
			i.items[idx] = updateItem
			return &i.items[idx], nil
		}
	}
	return nil, errors.New("Unexpected error")
}

func (i *itemMemoryRepository) Delete(itemId uint) error {
	for idx, item := range i.items {
		if item.ID == itemId {
			i.items = append(i.items[:idx], i.items[idx+1:]...)
			return nil
		}
	}
	return errors.New("Item not found")
}
