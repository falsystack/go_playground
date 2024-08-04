package repositories

import (
	"errors"
	"go-fleamarket/models"
)

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemInMemoryRepository(items []models.Item) ItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (i *ItemMemoryRepository) Delete(itemID uint) error {
	for idx, item := range i.items {
		if item.ID == itemID {
			i.items = append(i.items[:idx], i.items[idx+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

func (i *ItemMemoryRepository) Update(updatedItem models.Item) (*models.Item, error) {
	for idx, item := range i.items {
		if item.ID == updatedItem.ID {
			i.items[idx] = updatedItem
			return &i.items[idx], nil
		}
	}
	return nil, errors.New("item not found")
}

func (i *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(i.items) + 1)
	i.items = append(i.items, newItem)
	return &newItem, nil
}

func (i *ItemMemoryRepository) FindByID(itemID uint) (*models.Item, error) {
	for _, item := range i.items {
		if item.ID == itemID {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

func (i *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &i.items, nil
}
