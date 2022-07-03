package faker

import (
	"service_social_media/app/models"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func UserFaker(db *gorm.DB) *models.User {
	password := "12345678"
	hash, _ := HashPassword(password)

	return &models.User{
		ID:        uuid.NewV1().String(),
		NPM:       "50418069",
		Name:      faker.Name(),
		Username:  faker.Username(),
		Email:     faker.Email(),
		Password:  hash,
		Token:     "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}
