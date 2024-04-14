package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/models"
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

	r := gin.Default() // default ルータの初期化
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
