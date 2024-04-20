package main

import (
	"go-clean-arch/infra"
	"go-clean-arch/models"
	"log"
)

func main() {
	db := infra.NewDB()
	defer infra.CloseDB(db)
	if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Println("Migration Failed!")
		log.Fatalln(err)
	}
	log.Println("Successfully Migrated!")
}
