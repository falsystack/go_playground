package controller

import (
	"github.com/gin-gonic/gin"
	"go-clean-arch/usecase"
)

type TaskController interface {
	GetAllTasksByUserId(ctx gin.Context)
	GetTaskById(ctx gin.Context)
	CreateTask(ctx gin.Context)
	UpdateTask(ctx gin.Context)
	DeleteTask(ctx gin.Context)
}

type taskController struct {
	useCase usecase.TaskUseCase
}

func (t *taskController) GetAllTasksByUserId(ctx gin.Context) {
	var userId uint
	ctx.ShouldBindJSON()
}

func (t *taskController) GetTaskById(ctx gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *taskController) CreateTask(ctx gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *taskController) UpdateTask(ctx gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *taskController) DeleteTask(ctx gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewTaskController(useCase usecase.TaskUseCase) TaskController {
	return &taskController{useCase: useCase}
}
