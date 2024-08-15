package controller

import (
	"github.com/labstack/echo/v4"
	"go-echo/model"
	"go-echo/usecase"
	"net/http"
	"os"
	"time"
)

type UserController interface {
	Signup(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{uu: uu}
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{"csrf_token": token})
}

func (uc *userController) Signup(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := uc.uu.Signup(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, res)
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(&cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) Logout(c echo.Context) error {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		Expires:  time.Now(),
		Secure:   false,
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(&cookie)
	return c.NoContent(http.StatusOK)
}
