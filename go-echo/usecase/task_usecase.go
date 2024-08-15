package usecase

import (
	"go-echo/model"
	"go-echo/repository"
	"go-echo/validator"
)

type TaskUsecase interface {
	GetAllTasks(userID uint) ([]model.TaskResponse, error)
	GetTaskByID(userID uint, taskID uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userID uint, taskID uint) (model.TaskResponse, error)
	DeleteTask(userID uint, taskID uint) error
}

type taskUsecase struct {
	tr repository.TaskRepository
	tv validator.TaskValidator
}

func NewTaskUsecase(tr repository.TaskRepository, tv validator.TaskValidator) TaskUsecase {
	return &taskUsecase{tr: tr, tv: tv}
}

func (tu *taskUsecase) GetAllTasks(userID uint) ([]model.TaskResponse, error) {
	var tasks []model.Task
	if err := tu.tr.GetAllTasks(&tasks, userID); err != nil {
		return nil, err
	}
	var resTasks []model.TaskResponse
	for _, task := range tasks {
		resTasks = append(resTasks, model.TaskResponse{
			ID:        task.ID,
			Title:     task.Title,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		})
	}
	return resTasks, nil
}

func (tu *taskUsecase) GetTaskByID(userID uint, taskID uint) (model.TaskResponse, error) {
	var task model.Task
	if err := tu.tr.GetTaskByID(&task, userID, taskID); err != nil {
		return model.TaskResponse{}, err
	}

	return model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}

	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	return model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (tu *taskUsecase) UpdateTask(task model.Task, userID uint, taskID uint) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.UpdateTask(&task, userID, taskID); err != nil {
		return model.TaskResponse{}, err
	}
	return model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (tu *taskUsecase) DeleteTask(userID uint, taskID uint) error {
	if err := tu.tr.DeleteTask(userID, taskID); err != nil {
		return err
	}
	return nil
}
