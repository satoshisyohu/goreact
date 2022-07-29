package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required"`
}
