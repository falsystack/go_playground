package main

import (
	"goecho/controller"
	"goecho/db"
	"goecho/repository"
	"goecho/router"
	"goecho/usecase"
	"goecho/validator"
)

func main() {
	newDb := db.NewDB()
	// validator
	tv := validator.NewTaskValidator()
	uv := validator.NewUserValidator()
	// user
	ur := repository.NewUserRepository(newDb)
	uu := usecase.NewUserUsecase(ur, uv)
	uc := controller.NewUserController(uu)
	// task
	tr := repository.NewTaskRepository(newDb)
	tu := usecase.NewTaskUsecase(tr, tv)
	tc := controller.NewTaskController(tu)

	// router
	e := router.NewRouter(uc, tc)
	e.Logger.Fatal(e.Start(":8080"))
}
