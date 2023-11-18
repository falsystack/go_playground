package controller

import (
	"github.com/labstack/echo/v4"
	"goecho/model"
	"goecho/usecase"
	"net/http"
	"os"
	"time"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
}

type userControllerImpl struct {
	uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) IUserController {
	return &userControllerImpl{uu: uu}
}

func (uc *userControllerImpl) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userControllerImpl) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userControllerImpl) LogOut(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   -1,
	}
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
