package main

import (
	"fmt"
	"go-echo/db"
	"go-echo/model"
	"log"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	if err := dbConn.AutoMigrate(&model.User{}, &model.Task{}); err != nil {
		log.Fatalln("Auto migration failed : ", err)
	}
}
