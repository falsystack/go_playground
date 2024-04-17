package repositories

import (
	"errors"
	"go-gin/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint, userId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint, userId uint) error
}

func NewItemMemoryRepository(items []models.Item) ItemRepository {
	return &itemMemoryRepository{items: items}
}

type itemMemoryRepository struct {
	items []models.Item
}

func (i *itemMemoryRepository) FindById(itemId uint, userId uint) (*models.Item, error) {
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

func (i *itemMemoryRepository) Delete(itemId uint, userId uint) error {
	for idx, item := range i.items {
		if item.ID == itemId {
			i.items = append(i.items[:idx], i.items[idx+1:]...)
			return nil
		}
	}
	return errors.New("Item not found")
}

type itemRepositoryImpl struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepositoryImpl{db: db}
}

func (i *itemRepositoryImpl) Create(newItem models.Item) (*models.Item, error) {
	result := i.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

func (i *itemRepositoryImpl) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := i.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

func (i *itemRepositoryImpl) FindById(itemId uint, userId uint) (*models.Item, error) {
	var item models.Item
	result := i.db.First(&item, "id = ? AND user_id = ?", itemId, userId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("Item not found")
		}
		return nil, result.Error
	}
	return &item, nil
}

func (i *itemRepositoryImpl) Update(updateItem models.Item) (*models.Item, error) {
	result := i.db.Save(&updateItem) // update処理に使う
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}

func (i *itemRepositoryImpl) Delete(itemId uint, userId uint) error {
	targetItem, err := i.FindById(itemId, userId)
	if err != nil {
		return err
	}

	// gormのdeleteは論理削除（実際テーブルで削除されるのではなくdeletedAtにマークされる）
	// 物理削除したい場合Unscoped()を付ける必要がある。
	// i.db.Unscoped().Delete(&targetItem)
	result := i.db.Delete(&targetItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
