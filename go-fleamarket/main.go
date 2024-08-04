package main

import (
	"github.com/gin-gonic/gin"
	"go-fleamarket/controllers"
	"go-fleamarket/infra"
	"go-fleamarket/repositories"
	"go-fleamarket/services"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	//itemRepository := repositories.NewItemInMemoryRepository(items)
	itemRepository := repositories.NewItemORMRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindByID)
	r.PUT("/items/:id", itemController.Update)
	r.POST("/items", itemController.Create)
	r.DELETE("/items/:id", itemController.Delete)
	r.Run(":8080")

}
