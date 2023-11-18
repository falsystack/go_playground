package db

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		check(err)
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("DB"),
	)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	check(err)
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	err := sqlDB.Close()
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
