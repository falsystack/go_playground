package main

import (
	"github.com/joho/godotenv"
	"go-web-tucker/todo/app"
	"log"
	"net/http"
)

func init() {
	err := godotenv.Load("todo/.env")
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Error loading .env file")
	}
}

func main() {
	m := app.MakeHandler()
	defer m.Close()

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		panic(err)
	}
}
