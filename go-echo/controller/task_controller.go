package controller

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go-echo/model"
	"go-echo/usecase"
	"net/http"
	"strconv"
)

type TaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskByID(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.TaskUsecase
}

func NewTaskController(tu usecase.TaskUsecase) TaskController {
	return &taskController{tu: tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]

	tasks, err := tc.tu.GetAllTasks(uint(userID.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

func (tc *taskController) GetTaskByID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]
	id := c.Param("taskId")
	taskID, _ := strconv.Atoi(id)
	task, err := tc.tu.GetTaskByID(uint(userID.(float64)), uint(taskID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]

	var task model.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserID = uint(userID.(float64))
	taskResponse, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskResponse)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]
	id := c.Param("taskId")
	taskID, _ := strconv.Atoi(id)

	var task model.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	taskResponse, err := tc.tu.UpdateTask(task, uint(userID.(float64)), uint(taskID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskResponse)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]
	id := c.Param("taskId")
	taskID, _ := strconv.Atoi(id)

	if err := tc.tu.DeleteTask(uint(userID.(float64)), uint(taskID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
