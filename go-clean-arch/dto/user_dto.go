package dto

type UserResponse struct {
	// DTOなのにgormいる？
	ID    uint   `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
}
