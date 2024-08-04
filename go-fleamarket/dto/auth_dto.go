package dto

type SignupInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" biding:"required,min=8"`
}
