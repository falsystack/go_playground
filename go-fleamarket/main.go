package main

import (
	"github.com/gin-gonic/gin"
	"go-fleamarket/controllers"
	"go-fleamarket/models"
	"go-fleamarket/repositories"
	"go-fleamarket/services"
)

func main() {

	items := []models.Item{
		{
			ID:          1,
			Name:        "商品１",
			Price:       1000,
			Description: "説明１",
			SoldOut:     false,
		},
		{
			ID:          2,
			Name:        "商品２",
			Price:       2000,
			Description: "説明２",
			SoldOut:     true,
		},
		{
			ID:          3,
			Name:        "商品３",
			Price:       3000,
			Description: "説明３",
			SoldOut:     false,
		},
	}

	itemRepository := repositories.NewItemRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindByID)
	r.PUT("/items/:id", itemController.Update)
	r.POST("/items", itemController.Create)
	r.Run(":8080")

}
