package models

import "github.com/jinzhu/gorm"

type Vote struct {
	gorm.Model
	CommentId uint `json:"commentID" gorm:"not null"`
	UserID    uint `json:"userID" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
