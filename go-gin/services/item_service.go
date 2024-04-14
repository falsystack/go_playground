package services

import (
	"go-gin/models"
	"go-gin/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
}

type itemServiceImpl struct {
	repository repositories.ItemRepository
}

func NewItemService(repository repositories.ItemRepository) ItemService {
	return &itemServiceImpl{repository: repository}
}

func (s *itemServiceImpl) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}
