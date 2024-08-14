package repository

import (
	"go-echo/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	tx := ur.db.First(user, "email = ?", email)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	tx := ur.db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
