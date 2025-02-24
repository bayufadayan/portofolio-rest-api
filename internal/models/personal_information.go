package models

import (
	"gorm.io/gorm"
)

type PersonalInformation struct {
	gorm.Model
	Name         string `json:"name" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	Email        string `json:"email" binding:"required,max=100,email" gorm:"type:varchar(100);not null"`
	Phone        string `json:"phone" binding:"required,max=20" gorm:"type:varchar(20);not null"`
	Address      string `json:"address" gorm:"type:text"`
	Description  string `json:"description" gorm:"type:text"`
	ProfileImage string `json:"profile_image" binding:"required,url" gorm:"type:varchar(255);not null"`
	ResumeLink string `json:"resume_link" binding:"required,url" gorm:"type:varchar(255);not null"`
}
