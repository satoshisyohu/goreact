package model

import "gorm.io/gorm"

type Beans struct {
	gorm.Model
	Country  string `form:"beans" binding:"required" gorm:"unique;not null"`
	Quantity string `form:"quantity" binding:"required"`
}
