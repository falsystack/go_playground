package repository

import (
	"fmt"
	"go-clean-arch/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository interface {
	GetAllTasksByUserId(userId uint) (*[]models.Task, error)
	GetTaskById(userId uint, taskId uint) (*models.Task, error)
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type taskRepositoryImpl struct {
	db *gorm.DB
}

func (t *taskRepositoryImpl) GetAllTasksByUserId(userId uint) (*[]models.Task, error) {
	var tasks *[]models.Task
	tx := t.db.Joins("User").Where("user_id = ?", userId).Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tasks, nil
}

func (t *taskRepositoryImpl) GetTaskById(userId uint, taskId uint) (*models.Task, error) {
	var task *models.Task
	tx := t.db.Joins("User").Where("user_id = ?", userId).Where("task_id = ?", taskId).Find(&task)
	t.db.Joins("User").First(&task, "user_id = ? AND task_id = ?", userId, taskId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return task, nil
}

func (t *taskRepositoryImpl) CreateTask(task *models.Task) error {
	tx := t.db.Create(task)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *taskRepositoryImpl) UpdateTask(task *models.Task, userId uint, taskId uint) error {
	tx := t.db.
		Model(task).
		Clauses(clause.Returning{}).
		Where("id = ? AND user_id = ?", taskId, userId).
		Updates(task)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected < 1 {
		return fmt.Errorf("cannot update task")
	}
	return nil
}

func (t *taskRepositoryImpl) DeleteTask(userId uint, taskId uint) error {
	// bad
	// t.db.Where("task_id = ? AND user_id = ?", taskId, userId).Delete(&models.Task{})
	task, err := t.GetTaskById(userId, taskId)
	if err != nil {
		return err
	}
	tx := t.db.Delete(task)
	if tx.RowsAffected < 0 {
		return fmt.Errorf("cannot delete task")
	}
	return nil

}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepositoryImpl{db: db}
}
