package repository

import (
	"errors"
	"go-clean-arch/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func (u *userRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	// u.db.Where("email = ?", email).First(&user)と下は同じ表現
	tx := u.db.First(&user, "email = ?", email)
	if tx.Error != nil {
		if tx.Error.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
	}
	return &user, nil
}

func (u *userRepositoryImpl) CreateUser(user *models.User) error {
	tx := u.db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}
