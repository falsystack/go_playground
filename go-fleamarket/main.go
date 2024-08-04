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

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	itemRouter := r.Group("/items")
	authRouter := r.Group("/auth")

	itemRouter.GET("/", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindByID)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.POST("/", itemController.Create)
	itemRouter.DELETE("/:id", itemController.Delete)

	authRouter.POST("/", authController.Signup)
	r.Run(":8080")

}
