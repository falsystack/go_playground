package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/controllers"
	"go-gin/infra"
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

	r := gin.Default() // default ルータの初期化

	itemRouter := r.Group("/items")
	itemRouter.GET("/", itemController.FindAll)
	itemRouter.POST("/", itemController.Create)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
