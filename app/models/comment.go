package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	User      User      `json:"user,omitempty"`
	UserID    string    `gorm:"not null;foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user_id"`
	Post      Post      `json:"post,omitempty"`
	PostID    string    `gorm:"not null;foreignkey:PostID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"post_id"`
	Comment   string    `gorm:"type:text;not null" json:"comment"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
