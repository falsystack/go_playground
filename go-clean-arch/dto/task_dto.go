package dto

import "gorm.io/gorm"

type TaskResponse struct {
	gorm.Model
	Title string `json:"title" gorm:"not null"`
}
