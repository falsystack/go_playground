package services

import (
	"go-gin/models"
	"go-gin/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Signup(email string, password string) error
}

type authServiceImpl struct {
	repository repositories.AuthRepository
}

func (a *authServiceImpl) Signup(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return a.repository.CreateUser(user)
}

func NewAuthService(repository repositories.AuthRepository) AuthService {
	return &authServiceImpl{repository: repository}
}
