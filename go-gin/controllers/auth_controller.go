package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin/dto"
	"go-gin/services"
	"net/http"
)

type AuthController interface {
	Signup(ctx *gin.Context)
}

type authControllerImpl struct {
	authService services.AuthService
}

func (a *authControllerImpl) Signup(ctx *gin.Context) {
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.authService.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.Status(http.StatusCreated)
}

func NewAuthController(authService services.AuthService) AuthController {
	return &authControllerImpl{authService: authService}
}
