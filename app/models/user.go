package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	NPM       string    `gorm:"size:8;not null;uniqueIndex;" json:"npm"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Username  string    `gorm:"size:100;not null" json:"username"`
	Email     string    `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Password  string    `gorm:"->;<-;not null" json:"-"`
	Token     string    `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
