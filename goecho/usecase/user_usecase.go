package usecase

import (
	"goecho/model"
	"goecho/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type innerUserUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &innerUserUsecase{ur}
}

func (uu *innerUserUsecase) SignUp(user model.User) (model.UserResponse, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return model.UserResponse{}, err
	}
	nu := model.User{
		Email:    user.Email,
		Password: string(hp),
	}
	if err = uu.ur.CreateUser(&nu); err != nil {
		return model.UserResponse{}, err
	}
	res := model.UserResponse{
		ID:    nu.ID,
		Email: nu.Email,
	}
	return res, nil
}

func (uu *innerUserUsecase) Login(user model.User) (string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
