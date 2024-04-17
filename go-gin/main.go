package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/controllers"
	"go-gin/infra"
	"go-gin/middlewares"
	"go-gin/repositories"
	"go-gin/services"
)

func main() {
	infra.Init()
	db := infra.SetupDB()

	//items := []models.Item{
	//	{
	//		ID:          1,
	//		Name:        "商品１",
	//		Price:       1000,
	//		Description: "説明１",
	//		SoldOut:     false,
	//	},
	//	{
	//		ID:          2,
	//		Name:        "商品２",
	//		Price:       2000,
	//		Description: "説明２",
	//		SoldOut:     true,
	//	},
	//	{
	//		ID:          3,
	//		Name:        "商品３",
	//		Price:       3000,
	//		Description: "説明３",
	//		SoldOut:     false,
	//	},
	//}

	//itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default() // default ルータの初期化
	authRouter := r.Group("/auth")
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))

	itemRouter.GET("/", itemController.FindAll)
	itemRouterWithAuth.POST("/", itemController.Create)
	itemRouterWithAuth.GET("/:id", itemController.FindById)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
