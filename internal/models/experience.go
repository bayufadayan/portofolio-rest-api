package models

import (
	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	Title        string `json:"title" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	Company      string `json:"company" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	Location     string `json:"location" gorm:"type:varchar(255)"`
	StartDate    string `json:"start_date" binding:"required" gorm:"type:date;not null"`
	EndDate      string `json:"end_date" gorm:"type:date"`
	Description  string `json:"description" gorm:"type:text"`
	Technologies string `json:"technologies" gorm:"type:varchar(255)"`
}
