package models

import (
	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	Institution string `json:"institution" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	Degree      string `json:"degree" binding:"max=20" gorm:"type:varchar(20)"`
	Major       string `json:"major" binding:"max=100" gorm:"type:varchar(100)"`
	StartYear   int    `json:"start_year" binding:"required" gorm:"type:int;not null"`
	EndYear     int    `json:"end_year" gorm:"type:int"`
	Description string `json:"description" gorm:"type:text"`
}
