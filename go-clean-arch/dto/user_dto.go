package dto

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
}
