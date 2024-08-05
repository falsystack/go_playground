package services

import (
	"github.com/golang-jwt/jwt/v5"
	"go-fleamarket/models"
	"go-fleamarket/repositories"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService interface {
	Signup(email, password string) error
	Login(email, password string) (*string, error)
}

type authService struct {
	repository repositories.AuthRepository
}

func (a *authService) Login(email, password string) (*string, error) {
	findUser, err := a.repository.FindUser(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	token, err := CreateToken(findUser.ID, findUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a *authService) Signup(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return a.repository.CreateUser(u)
}

func NewAuthService(repository repositories.AuthRepository) AuthService {
	return &authService{repository: repository}
}

func CreateToken(userID uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
