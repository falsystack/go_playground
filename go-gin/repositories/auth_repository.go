package repositories

import (
	"errors"
	"go-gin/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user models.User) error
	FindUser(email string) (*models.User, error)
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func (a *authRepositoryImpl) FindUser(email string) (*models.User, error) {
	var user models.User
	result := a.db.First(&user, "email = ?", email)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
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
