package usecase

import (
	"github.com/golang-jwt/jwt/v5"
	"go-echo/model"
	"go-echo/repository"
	"go-echo/validator"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserUsecase interface {
	Signup(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.UserRepository
	uv validator.UserValidator
}

func NewUserUsecase(ur repository.UserRepository, uv validator.UserValidator) UserUsecase {
	return &userUsecase{ur: ur, uv: uv}
}

func (uu *userUsecase) Signup(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	response := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return response, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
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
