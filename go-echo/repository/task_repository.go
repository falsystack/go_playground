package repository

import (
	"fmt"
	"go-echo/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userID uint) error
	GetTaskByID(task *model.Task, userID uint, taskID uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userID uint, taskID uint) error
	DeleteTask(userID uint, taskID uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userID uint) error {
	if err := tr.db.
		Joins("User").
		Where("user_id = ?", userID).
		Order("created_at").
		Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskByID(task *model.Task, userID uint, taskID uint) error {
	if err := tr.db.
		Joins("User").
		Where("user_id = ?", userID).
		Find(task, taskID).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userID uint, taskID uint) error {
	// 更新した後のtaskのオブジェクトをModel()に渡したポインターが指し示した書き込んでくれる
	tx := tr.db.Model(task).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", taskID, userID).Update("title", task.Title)
	if tx.Error != nil {
		return tx.Error
	}
	// 更新に失敗してもエラーにならないのでrowsaAffectを確認する
	if tx.RowsAffected < 1 {
		return fmt.Errorf("task with id %d not found", taskID)
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userID uint, taskID uint) error {
	tx := tr.db.Where("id = ? AND user_id = ?", taskID, userID).Delete(&model.Task{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected < 1 {
		return fmt.Errorf("task with id %d not found", taskID)
	}
	return nil
}
