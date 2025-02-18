package models

import (
	"gorm.io/gorm"
)

type JobTitle struct {
	gorm.Model
	Title                 string `gorm:"type:varchar(100);not null"`
	PersonalInformationID uint   `gorm:"not null"`
}
