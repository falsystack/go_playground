package controllers

import (
	"github.com/gin-gonic/gin"
	"go-fleamarket/services"
	"net/http"
	"strconv"
)

type ItemController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type ItemControllerImpl struct {
	service services.ItemService
}

func (i *ItemControllerImpl) FindByID(ctx *gin.Context) {
	parsedItemID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	item, err := i.service.FindByID(uint(parsedItemID))
	if err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

func NewItemController(service services.ItemService) ItemController {
	return &ItemControllerImpl{service: service}
}

func (i *ItemControllerImpl) FindAll(ctx *gin.Context) {
	items, err := i.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
