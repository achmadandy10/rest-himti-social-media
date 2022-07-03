package repositories

import (
	"service_social_media/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	InsertComment(comment models.Comment) models.Comment
	UpdateComment(comment models.Comment) models.Comment
	DeleteComment(comment models.Comment)
	AllComment() []models.Comment
	FindCommentByID(commentID string) models.Comment
}

type commentConnection struct {
	connection *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentConnection{
		connection: db,
	}
}

func (db *commentConnection) InsertComment(comment models.Comment) models.Comment {
	comment.ID = uuid.New().String()

	db.connection.Save(&comment)
	db.connection.Preload("User").Preload("Post").Where("id = ?", comment.ID).Take(&comment)

	return comment
}

func (db *commentConnection) UpdateComment(comment models.Comment) models.Comment {
	comment.ID = uuid.New().String()

	db.connection.Save(&comment)
	db.connection.Preload("User").Preload("Post").Where("id = ?", comment.ID).Take(&comment)

	return comment
}

func (db *commentConnection) DeleteComment(comment models.Comment) {
	db.connection.Delete(&comment)
}

func (db *commentConnection) AllComment() []models.Comment {
	var professions []models.Comment
	db.connection.Preload("User").Preload("Post").Find(&professions)
	return professions
}

func (db *commentConnection) FindCommentByID(commentID string) models.Comment {
	var comment models.Comment
	db.connection.Preload("User").Preload("Post").Where("id = ?", commentID).Take(&comment)
	return comment
}
