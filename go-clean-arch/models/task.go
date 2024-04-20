package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	User   User   `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId uint   `json:"user_id" gorm:"not null"`
}
