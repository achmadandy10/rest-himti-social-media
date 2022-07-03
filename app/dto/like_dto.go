package dto

type LikeDTO struct {
	UserID string `json:"user_id" form:"user_id" binding:"required"`
	PostID string `json:"post_id" form:"post_id" binding:"required"`
}
