package repositories

import (
	"errors"
	"go-fleamarket/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindByID(itemID uint, userID uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updatedItem models.Item) (*models.Item, error)
	Delete(itemID uint, userID uint) error
}

type itemORMRepository struct {
	db *gorm.DB
}

func (i *itemORMRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := i.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

func (i *itemORMRepository) FindByID(itemID uint, userID uint) (*models.Item, error) {
	var item models.Item
	result := i.db.First(&item, "id = ? AND user_id = ?", itemID, userID)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("item not found")
		}
		return nil, result.Error
	}

	return &item, nil
}

func (i *itemORMRepository) Create(newItem models.Item) (*models.Item, error) {
	result := i.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newItem, nil
}

func (i *itemORMRepository) Update(updatedItem models.Item) (*models.Item, error) {
	result := i.db.Save(updatedItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedItem, nil
}

func (i *itemORMRepository) Delete(itemID uint, userID uint) error {
	deleteItem, err := i.FindByID(itemID, userID)
	if err != nil {
		return err
	}

	// 물리삭제 : i.db.Unscoped().Delete(&deleteItem)
	// 논리삭제 : i.db.Delete(&deleteItem)
	result := i.db.Unscoped().Delete(&deleteItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewItemORMRepository(db *gorm.DB) ItemRepository {
	return &itemORMRepository{db: db}
}
