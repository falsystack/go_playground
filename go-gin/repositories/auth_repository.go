package repositories

import (
	"go-gin/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user models.User) error
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func (a *authRepositoryImpl) CreateUser(user models.User) error {
	result := a.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{db: db}
}
