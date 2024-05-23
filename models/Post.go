package models

import "gorm.io/gorm"

type Post struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   uint   `json:"user_id"`
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `gorm:"type:TEXT" json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
	gorm.Model
}
