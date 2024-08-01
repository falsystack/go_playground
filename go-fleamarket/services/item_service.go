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
}

func NewItemService(repository repositories.ItemRepository) ItemService {
	return &ItemServiceImpl{repository: repository}
}

type ItemServiceImpl struct {
	repository repositories.ItemRepository
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
