package main

import (
	"github.com/gin-gonic/gin"
	"go-clean-arch/controller"
	"go-clean-arch/infra"
	"go-clean-arch/repository"
	"go-clean-arch/router"
	"go-clean-arch/usecase"
	"gorm.io/gorm"
	"log"
)

func main() {
	r := gin.Default()
	db := infra.NewDB()

	userDomain(db, r)

	log.Fatalln(r.Run(":8080"))
}

func userDomain(db *gorm.DB, r *gin.Engine) {
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	router.UserRouter(userController, r)
}
