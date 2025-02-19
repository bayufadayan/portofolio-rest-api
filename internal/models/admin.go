package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string `json:"name" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" binding:"required,email" gorm:"type:varchar(100);not null"`
	Password string `json:"password" binding:"required,min=4" gorm:"type:varchar(100);not null"`
}
