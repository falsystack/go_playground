package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin/services"
	"net/http"
)

type ItemController interface {
	FindAll(ctx *gin.Context)
}

type itemController struct {
	service services.ItemService
}

func NewItemController(service services.ItemService) ItemController {
	return &itemController{service: service}
}

func (c *itemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
