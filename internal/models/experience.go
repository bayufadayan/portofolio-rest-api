package models

import (
	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	Title          string `gorm:"type:varchar(100);not null"`
	Company        string `gorm:"type:varchar(100);not null"`
	Location       string `gorm:"type:varchar(255)"`
	StartDate      string `gorm:"type:date;not null"`
	EndDate        string `gorm:"type:date"`
	Description    string `gorm:"type:text"`
	Technologies   string `gorm:"type:varchar(255)"`
}
