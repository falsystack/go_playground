package controller

import (
	"github.com/gin-gonic/gin"
	"go-clean-arch/models"
	"go-clean-arch/usecase"
	"net/http"
	"os"
	"time"
)

type UserController interface {
	Signup(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type userController struct {
	uu usecase.UserUseCase
}

func (u *userController) Signup(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := u.uu.Signup(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": resp})
}

func (u *userController) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := u.uu.Login(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie(
		"token",
		token,
		time.Now().Add(time.Hour*24).Second(), // 24時間
		"/",
		os.Getenv("API_DOMAIN"),
		false, // testの為にfalse
		true)
	ctx.Status(http.StatusOK)
}

func (u *userController) Logout(ctx *gin.Context) {
	ctx.SetCookie(
		"token",
		"",
		-1,
		"/",
		os.Getenv("API_DOMAIN"),
		false, // testの為にfalse
		true)
	ctx.Status(http.StatusNoContent)
}

func NewUserController(uu usecase.UserUseCase) UserController {
	return &userController{uu: uu}
}
