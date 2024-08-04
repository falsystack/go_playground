package services

import (
	"go-fleamarket/dto"
	"go-fleamarket/models"
	"go-fleamarket/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
	FindByID(itemID uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	Update(itemID uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemID uint) error
}

func NewItemService(repository repositories.ItemRepository) ItemService {
	return &ItemServiceImpl{repository: repository}
}

type ItemServiceImpl struct {
	repository repositories.ItemRepository
}

func (i *ItemServiceImpl) Delete(itemID uint) error {
	return i.repository.Delete(itemID)
}

func (i *ItemServiceImpl) Update(itemID uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	foundItem, err := i.FindByID(itemID)
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

func (i *ItemServiceImpl) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return i.repository.Create(newItem)
}

func (i *ItemServiceImpl) FindByID(itemID uint) (*models.Item, error) {
	return i.repository.FindByID(itemID)
}

func (i *ItemServiceImpl) FindAll() (*[]models.Item, error) {
	return i.repository.FindAll()
}
