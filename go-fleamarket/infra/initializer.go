package infra

import (
	"github.com/joho/godotenv"
	"log"
)

func Initialize() {
	// Initialize 함수를 실행하는 함수의 위치에 따라 패스가 변하는 문제가 있음.
	// root 의 main.go 에서 실행할 경우 .env 지만 migration 디렉토리 안의 migration.go 에서 실행할 경우 ../.env 여야 함
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}
