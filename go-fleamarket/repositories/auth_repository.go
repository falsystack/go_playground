package repositories

import (
	"errors"
	"go-fleamarket/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user models.User) error
	FindUser(email string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func (a *authRepository) FindUser(email string) (*models.User, error) {
	var user models.User
	result := a.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (a *authRepository) CreateUser(user models.User) error {
	result := a.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}
