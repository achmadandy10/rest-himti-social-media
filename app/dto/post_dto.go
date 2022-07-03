package dto

type PostDTO struct {
	UserID string `json:"user_id" form:"user_id" binding:"required"`
	Text   string `json:"text" form:"text" binding:"required"`
	Slug   string `json:"slug" form:"slug" binding:"required"`
}
