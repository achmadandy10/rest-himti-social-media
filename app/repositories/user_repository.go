package repositories

import (
	"log"
	"service_social_media/app/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user models.User) models.User
	UpdateUser(user models.User) models.User
	VerifyCredential(username string, password string) interface{}
	IsDuplicateUsername(username string) (tx *gorm.DB)
	IsDuplicateNPM(npm string) (tx *gorm.DB)
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) models.User
	ProfileUser(userID string) models.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user models.User) models.User {
	user.ID = uuid.New().String()
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser models.User
		db.connection.Find(&tempUser, user.ID)
		user.ID = tempUser.ID
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user models.User) models.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser models.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(useraname string, password string) interface{} {
	var user models.User
	res := db.connection.Where("username = ?", useraname).Or("email = ?", useraname).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateUsername(username string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("username = ?", username).Take(&user)
}

func (db *userConnection) IsDuplicateNPM(npm string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("npm = ?", npm).Take(&user)
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) models.User {
	var user models.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) models.User {
	var user models.User
	db.connection.Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
		panic("Failed to hash a Password")
	}

	return string(hash)
}
