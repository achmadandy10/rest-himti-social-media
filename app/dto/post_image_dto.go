package dto

type PostImageDTO struct {
	PostID string `json:"post_id" form:"post_id" binding:"required"`
	Path   string `json:"path" form:"path" binding:"required"`
}
