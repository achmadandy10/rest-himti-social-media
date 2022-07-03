package repositories

import (
	"service_social_media/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository interface {
	InsertPost(post models.Post) models.Post
	UpdatePost(post models.Post) models.Post
	DeletePost(post models.Post)
	AllPost() []models.Post
	FindPostByID(postID string) models.Post
}

type postConnection struct {
	connection *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postConnection{
		connection: db,
	}
}

func (db *postConnection) InsertPost(post models.Post) models.Post {
	post.ID = uuid.New().String()

	db.connection.Save(&post)
	db.connection.Preload("User").Preload("Like").Preload("Comment").Preload("PostImage").Where("id = ?", post.ID).Take(&post)

	return post
}

func (db *postConnection) UpdatePost(post models.Post) models.Post {
	post.ID = uuid.New().String()

	db.connection.Save(&post)
	db.connection.Preload("User").Preload("Like").Preload("Comment").Preload("PostImage").Where("id = ?", post.ID).Take(&post)

	return post
}

func (db *postConnection) DeletePost(post models.Post) {
	db.connection.Delete(&post)
}

func (db *postConnection) AllPost() []models.Post {
	var professions []models.Post
	db.connection.Preload("User").Preload("Like").Preload("Comment").Preload("PostImage").Find(&professions)
	return professions
}

func (db *postConnection) FindPostByID(postID string) models.Post {
	var post models.Post
	db.connection.Preload("User").Preload("Like").Preload("Comment").Preload("PostImage").Where("id = ?", postID).Take(&post)
	return post
}
