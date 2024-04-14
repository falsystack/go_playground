package repositories

import "go-gin/models"

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
}

func NewItemMemoryRepository(items []models.Item) ItemRepository {
	return &itemMemoryRepository{items: items}
}

type itemMemoryRepository struct {
	items []models.Item
}

func (r *itemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}
