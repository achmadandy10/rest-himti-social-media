package services

import (
	"log"
	"service_social_media/app/dto"
	"service_social_media/app/models"
	"service_social_media/app/repositories"

	"github.com/mashingan/smapping"
)

type PostService interface {
	Insert(post dto.PostDTO) models.Post
	Update(postID dto.PostDTO) models.Post
	Delete(post models.Post)
	All() []models.Post
	FindByID(postID string) models.Post
}

type postService struct {
	postRepository repositories.PostRepository
}

func NewPostService(postRep repositories.PostRepository) PostService {
	return &postService{
		postRepository: postRep,
	}
}

func (service *postService) Insert(p dto.PostDTO) models.Post {
	post := models.Post{}
	err := smapping.FillStruct(&post, smapping.MapFields(&p))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.postRepository.InsertPost(post)

	return res
}

func (service *postService) Update(p dto.PostDTO) models.Post {
	post := models.Post{}
	err := smapping.FillStruct(&post, smapping.MapFields(&p))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)

	}
	res := service.postRepository.UpdatePost(post)

	return res
}

func (service *postService) Delete(post models.Post) {
	service.postRepository.DeletePost(post)
}

func (service *postService) All() []models.Post {
	return service.postRepository.AllPost()
}

func (service *postService) FindByID(postID string) models.Post {
	return service.postRepository.FindPostByID(postID)
}
