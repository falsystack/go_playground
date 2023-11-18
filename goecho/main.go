package main

import (
	"goecho/controller"
	"goecho/db"
	"goecho/repository"
	"goecho/router"
	"goecho/usecase"
)

func main() {
	newDb := db.NewDB()
	// user
	ur := repository.NewUserRepository(newDb)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)
	// task
	tr := repository.NewTaskRepository(newDb)
	tu := usecase.NewTaskUsecase(tr)
	tc := controller.NewTaskController(tu)

	// router
	e := router.NewRouter(uc, tc)
	e.Logger.Fatal(e.Start(":8080"))
}
