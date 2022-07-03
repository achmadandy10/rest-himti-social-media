package repositories

import (
	"service_social_media/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LikeRepository interface {
	InsertLike(like models.Like) models.Like
	UpdateLike(like models.Like) models.Like
	DeleteLike(like models.Like)
	AllLike() []models.Like
	FindLikeByID(likeID string) models.Like
}

type likeConnection struct {
	connection *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeConnection{
		connection: db,
	}
}

func (db *likeConnection) InsertLike(like models.Like) models.Like {
	like.ID = uuid.New().String()

	db.connection.Save(&like)
	db.connection.Preload("User").Preload("Post").Where("id = ?", like.ID).Take(&like)

	return like
}

func (db *likeConnection) UpdateLike(like models.Like) models.Like {
	like.ID = uuid.New().String()

	db.connection.Save(&like)
	db.connection.Preload("User").Preload("Post").Where("id = ?", like.ID).Take(&like)

	return like
}

func (db *likeConnection) DeleteLike(like models.Like) {
	db.connection.Delete(&like)
}

func (db *likeConnection) AllLike() []models.Like {
	var professions []models.Like
	db.connection.Preload("User").Preload("Post").Find(&professions)
	return professions
}

func (db *likeConnection) FindLikeByID(likeID string) models.Like {
	var like models.Like
	db.connection.Preload("User").Preload("Post").Where("id = ?", likeID).Take(&like)
	return like
}
