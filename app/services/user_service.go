package services

import (
	"log"
	"service_social_media/app/dto"
	"service_social_media/app/models"
	"service_social_media/app/repositories"

	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserDTO) models.User
	Profile(userID string) models.User
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserDTO) models.User {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) models.User {
	return service.userRepository.ProfileUser(userID)
}
