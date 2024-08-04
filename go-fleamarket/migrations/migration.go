package main

import (
	"go-fleamarket/infra"
	"go-fleamarket/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("failed to auto migrate items")
	}
}
w