package models

import (
	"time"

	"gorm.io/gorm"
)

type PostImage struct {
	ID        string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	Post      Post      `json:"post,omitempty"`
	PostID    string    `gorm:"not null;foreignkey:PostID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"post_id"`
	Path      string    `gorm:"size:255;not null" json:"path"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
