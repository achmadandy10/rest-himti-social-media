package dto

type UserDTO struct {
	ID       string `json:"id" form:"id" binding:"required"`
	NPM      string `json:"npm" form:"npm" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
}
