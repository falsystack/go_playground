package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-fleamarket/controllers"
	"go-fleamarket/infra"
	"go-fleamarket/middlewares"
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
	r.Use(cors.Default())
	itemRouter := r.Group("/items")
	itemRouterWithAuth := itemRouter.Group("", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")

	itemRouter.GET("/", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindByID)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter.POST("/", authController.Signup)
	authRouter.POST("/login", authController.Login)
	r.Run(":8080")

}
