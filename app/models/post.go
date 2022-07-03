package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID         string       `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	User       User         `json:"user,omitempty"`
	UserID     string       `gorm:"not null;foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user_id"`
	Text       string       `gorm:"type:text;not null" json:"text"`
	Slug       string       `gorm:"size:255;not null" json:"slug"`
	Comments   *[]Comment   `json:"comment,omitempty"`
	Likes      *[]Like      `json:"like,omitempty"`
	PostImages *[]PostImage `json:"post_image,omitempty"`
	CreatedAt  time.Time    `json:"create_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  gorm.DeletedAt
}
