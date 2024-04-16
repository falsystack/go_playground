package main

import (
	"go-gin/infra"
	"go-gin/models"
)

func main() {
	infra.Init()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic("Failed to migrate database")
	}
}
