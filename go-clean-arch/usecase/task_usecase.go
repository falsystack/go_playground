package usecase

import (
	"go-clean-arch/dto"
	"go-clean-arch/models"
	"go-clean-arch/repository"
)

type TaskUseCase interface {
	GetAllTasksByUserId(userId uint) ([]dto.TaskResponse, error)
	GetTaskById(taskId, userId uint) (*dto.TaskResponse, error)
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task, userId, taskId uint) error
	DeleteTask(taskId, userId uint) error
}

type taskUseCase struct {
	repository repository.TaskRepository
}

func (t *taskUseCase) GetAllTasksByUserId(userId uint) ([]dto.TaskResponse, error) {
	tasks, err := t.repository.GetAllTasksByUserId(userId)
	if err != nil {
		return nil, err
	}

	var taskResponses []dto.TaskResponse
	for _, task := range *tasks {
		taskResponses = append(taskResponses, dto.TaskResponse{
			Model: task.Model,
			Title: task.Title,
		})
	}
	return taskResponses, nil
}

func (t *taskUseCase) GetTaskById(userId, taskId uint) (*dto.TaskResponse, error) {
	task, err := t.repository.GetTaskById(userId, taskId)
	if err != nil {
		return nil, err
	}
	taskResponse := &dto.TaskResponse{
		Model: task.Model,
		Title: task.Title,
	}
	return taskResponse, nil
}

func (t *taskUseCase) CreateTask(task *models.Task) error {
	if err := t.repository.CreateTask(task); err != nil {
		return err
	}
	return nil
}

func (t *taskUseCase) UpdateTask(task *models.Task, userId, taskId uint) error {
	if err := t.repository.UpdateTask(task, userId, taskId); err != nil {
		return err
	}
	return nil
}

func (t *taskUseCase) DeleteTask(userId, taskId uint) error {
	if err := t.repository.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}

func NewTaskUseCase(repository repository.TaskRepository) TaskUseCase {
	return &taskUseCase{repository: repository}
}
