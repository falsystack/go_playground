package services

import (
	"github.com/golang-jwt/jwt/v5"
	"go-gin/models"
	"go-gin/repositories"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService interface {
	Signup(email string, password string) error
	Login(email string, password string) (*string, error)
}

type authServiceImpl struct {
	repository repositories.AuthRepository
}

func createToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId, // subject : userの識別子、ここではuserId
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (a *authServiceImpl) Login(email string, password string) (*string, error) {
	foundUser, err := a.repository.FindUser(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	token, err := createToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
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
