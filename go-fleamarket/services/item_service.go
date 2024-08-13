package services

import (
	"go-fleamarket/dto"
	"go-fleamarket/models"
	"go-fleamarket/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
	FindByID(itemID uint, userID uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput, userID uint) (*models.Item, error)
	Update(itemID uint, updateItemInput dto.UpdateItemInput, userID uint) (*models.Item, error)
	Delete(itemID uint, userID uint) error
}

func NewItemService(repository repositories.ItemRepository) ItemService {
	return &ItemServiceImpl{repository: repository}
}

type ItemServiceImpl struct {
	repository repositories.ItemRepository
}

func (i *ItemServiceImpl) Delete(itemID uint, userID uint) error {
	return i.repository.Delete(itemID, userID)
}

func (i *ItemServiceImpl) Update(itemID uint, updateItemInput dto.UpdateItemInput, userID uint) (*models.Item, error) {
	foundItem, err := i.FindByID(itemID, userID)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != nil {
		foundItem.Name = *updateItemInput.Name
	}

	if updateItemInput.Price != nil {
		foundItem.Price = *updateItemInput.Price
	}

	if updateItemInput.Description != nil {
		foundItem.Description = *updateItemInput.Description
	}

	if updateItemInput.SoldOut != nil {
		foundItem.SoldOut = *updateItemInput.SoldOut
	}

	return i.repository.Update(*foundItem)
}

func (i *ItemServiceImpl) Create(createItemInput dto.CreateItemInput, userID uint) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
		UserID:      userID,
	}
	return i.repository.Create(newItem)
}

func (i *ItemServiceImpl) FindByID(itemID uint, userID uint) (*models.Item, error) {
	return i.repository.FindByID(itemID, userID)
}

func (i *ItemServiceImpl) FindAll() (*[]models.Item, error) {
	return i.repository.FindAll()
}
