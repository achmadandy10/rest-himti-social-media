package services

import (
	"log"
	"service_social_media/app/dto"
	"service_social_media/app/models"
	"service_social_media/app/repositories"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(username string, password string) interface{}
	CreateUser(user dto.RegisterDTO) models.User
	FindByEmail(email string) models.User
	IsDuplicateUsername(username string) bool
	IsDuplicateNPM(npm string) bool
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRep repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(username string, password string) interface{} {
	res := service.userRepository.VerifyCredential(username, password)
	if v, oke := res.(models.User); oke {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Username == username && comparedPassword {
			return res
		} else if v.Email == username && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) models.User {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) models.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateUsername(useraname string) bool {
	res := service.userRepository.IsDuplicateUsername(useraname)
	return !(res.Error == nil)
}

func (service *authService) IsDuplicateNPM(npm string) bool {
	res := service.userRepository.IsDuplicateNPM(npm)
	return !(res.Error == nil)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
