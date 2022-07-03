package services

import (
	"log"
	"service_social_media/app/dto"
	"service_social_media/app/models"
	"service_social_media/app/repositories"

	"github.com/mashingan/smapping"
)

type LikeService interface {
	Insert(like dto.LikeDTO) models.Like
	Update(likeID dto.LikeDTO) models.Like
	Delete(like models.Like)
	All() []models.Like
	FindByID(likeID string) models.Like
}

type likeService struct {
	likeRepository repositories.LikeRepository
}

func NewLikeService(likeRep repositories.LikeRepository) LikeService {
	return &likeService{
		likeRepository: likeRep,
	}
}

func (service *likeService) Insert(p dto.LikeDTO) models.Like {
	like := models.Like{}
	err := smapping.FillStruct(&like, smapping.MapFields(&p))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.likeRepository.InsertLike(like)

	return res
}

func (service *likeService) Update(p dto.LikeDTO) models.Like {
	like := models.Like{}
	err := smapping.FillStruct(&like, smapping.MapFields(&p))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)

	}
	res := service.likeRepository.UpdateLike(like)

	return res
}

func (service *likeService) Delete(like models.Like) {
	service.likeRepository.DeleteLike(like)
}

func (service *likeService) All() []models.Like {
	return service.likeRepository.AllLike()
}

func (service *likeService) FindByID(likeID string) models.Like {
	return service.likeRepository.FindLikeByID(likeID)
}
