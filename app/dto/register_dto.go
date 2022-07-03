package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	NPM      string `json:"npm" form:"npm" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=8"`
}
