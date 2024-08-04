package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Price       uint   `gorm:"not null"`
	Description string
	SoldOut     bool `gorm:"default:false;not null"`
}
