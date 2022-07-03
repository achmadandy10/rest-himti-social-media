package services

import (
	"log"
	"service_social_media/app/dto"
	"service_social_media/app/models"
	"service_social_media/app/repositories"

	"github.com/mashingan/smapping"
)

type PostImageService interface {
	Insert(post_image dto.PostImageDTO) models.PostImage
	Update(postImageID dto.PostImageDTO) models.PostImage
	Delete(post_image models.PostImage)
	All() []models.PostImage
	FindByID(postImageID string) models.PostImage
}

type postImageService struct {
	postImageRepository repositories.PostImageRepository
}

func NewPostImageService(postImageRep repositories.PostImageRepository) PostImageService {
	return &postImageService{
		postImageRepository: postImageRep,
	}
}

func (service *postImageService) Insert(p dto.PostImageDTO) models.PostImage {
	post_image := models.PostImage{}
	err := smapping.FillStruct(&post_image, smapping.MapFields(&p))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.postImageRepository.InsertPostImage(post_image)

	return res
}

func (service *postImageService) Update(p dto.PostImageDTO) models.PostImage {
	post_image := models.PostImage{}
	err := smapping.FillStruct(&post_image, smapping.MapFields(&p))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)

	}
	res := service.postImageRepository.UpdatePostImage(post_image)

	return res
}

func (service *postImageService) Delete(post_image models.PostImage) {
	service.postImageRepository.DeletePostImage(post_image)
}

func (service *postImageService) All() []models.PostImage {
	return service.postImageRepository.AllPostImage()
}

func (service *postImageService) FindByID(postImageID string) models.PostImage {
	return service.postImageRepository.FindPostImageByID(postImageID)
}
