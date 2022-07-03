package services

import (
	"log"
	"service_social_media/app/dto"
	"service_social_media/app/models"
	"service_social_media/app/repositories"

	"github.com/mashingan/smapping"
)

type CommentService interface {
	Insert(comment dto.CommentDTO) models.Comment
	Update(commentID dto.CommentDTO) models.Comment
	Delete(comment models.Comment)
	All() []models.Comment
	FindByID(commentID string) models.Comment
}

type commentService struct {
	commentRepository repositories.CommentRepository
}

func NewCommentService(commentRep repositories.CommentRepository) CommentService {
	return &commentService{
		commentRepository: commentRep,
	}
}

func (service *commentService) Insert(p dto.CommentDTO) models.Comment {
	comment := models.Comment{}
	err := smapping.FillStruct(&comment, smapping.MapFields(&p))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.commentRepository.InsertComment(comment)

	return res
}

func (service *commentService) Update(p dto.CommentDTO) models.Comment {
	comment := models.Comment{}
	err := smapping.FillStruct(&comment, smapping.MapFields(&p))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)

	}
	res := service.commentRepository.UpdateComment(comment)

	return res
}

func (service *commentService) Delete(comment models.Comment) {
	service.commentRepository.DeleteComment(comment)
}

func (service *commentService) All() []models.Comment {
	return service.commentRepository.AllComment()
}

func (service *commentService) FindByID(commentID string) models.Comment {
	return service.commentRepository.FindCommentByID(commentID)
}
