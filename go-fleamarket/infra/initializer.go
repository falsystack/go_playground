package infra

import (
	"github.com/joho/godotenv"
	"log"
)

func Initialize() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}
