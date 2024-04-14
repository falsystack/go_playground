package services

import (
	"go-gin/models"
	"go-gin/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
}

type itemServiceImpl struct {
	repository repositories.ItemRepository
}

func NewItemService(repository repositories.ItemRepository) ItemService {
	return &itemServiceImpl{repository: repository}
}

func (i *itemServiceImpl) FindAll() (*[]models.Item, error) {
	return i.repository.FindAll()
}

func (i *itemServiceImpl) FindById(itemId uint) (*models.Item, error) {
	return i.repository.FindById(itemId)
}
