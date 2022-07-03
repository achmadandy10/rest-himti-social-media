package repositories

import (
	"service_social_media/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostImageRepository interface {
	InsertPostImage(post_image models.PostImage) models.PostImage
	UpdatePostImage(post_image models.PostImage) models.PostImage
	DeletePostImage(post_image models.PostImage)
	AllPostImage() []models.PostImage
	FindPostImageByID(postImageID string) models.PostImage
}

type postImageConnection struct {
	connection *gorm.DB
}

func NewPostImageRepository(db *gorm.DB) PostImageRepository {
	return &postImageConnection{
		connection: db,
	}
}

func (db *postImageConnection) InsertPostImage(post_image models.PostImage) models.PostImage {
	post_image.ID = uuid.New().String()

	db.connection.Save(&post_image)
	db.connection.Preload("User").Preload("Post").Where("id = ?", post_image.ID).Take(&post_image)

	return post_image
}

func (db *postImageConnection) UpdatePostImage(post_image models.PostImage) models.PostImage {
	post_image.ID = uuid.New().String()

	db.connection.Save(&post_image)
	db.connection.Preload("User").Preload("Post").Where("id = ?", post_image.ID).Take(&post_image)

	return post_image
}

func (db *postImageConnection) DeletePostImage(post_image models.PostImage) {
	db.connection.Delete(&post_image)
}

func (db *postImageConnection) AllPostImage() []models.PostImage {
	var professions []models.PostImage
	db.connection.Preload("User").Preload("Post").Find(&professions)
	return professions
}

func (db *postImageConnection) FindPostImageByID(postImageID string) models.PostImage {
	var post_image models.PostImage
	db.connection.Preload("User").Preload("Post").Where("id = ?", postImageID).Take(&post_image)
	return post_image
}
