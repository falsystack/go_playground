package usecase

import (
	"github.com/golang-jwt/jwt/v5"
	"go-clean-arch/dto"
	"go-clean-arch/models"
	"go-clean-arch/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserUseCase interface {
	Signup(user *models.User) (dto.UserResponse, error)
	Login(user *models.User) (string, error)
}

type userUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &userUseCase{
		repository: repository,
	}
}

func (u *userUseCase) Signup(user *models.User) (dto.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponse{}, err
	}
	newUser := models.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	if err := u.repository.CreateUser(&newUser); err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}, nil
}

func (u *userUseCase) Login(user *models.User) (string, error) {
	foundUser, err := u.repository.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	// TODO: token生成の役割分離
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   foundUser.ID,
		"email": foundUser.Email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
