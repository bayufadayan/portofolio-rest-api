package models

import (
	"gorm.io/gorm"
)

type PersonalInformation struct {
	gorm.Model
	Name          string        `gorm:"type:varchar(100);not null"`
	Email         string        `gorm:"type:varchar(100);not null"`
	Phone         string        `gorm:"type:varchar(20);not null"`
	Address       string        `gorm:"type:text"`
	Description   string        `gorm:"type:text"`
	ProfileImage  string        `gorm:"type:varchar(255);not null"`
	SocialMedias []SocialMedia  `gorm:"foreignKey:PersonalInformationID"`
	JobTitles    []JobTitle     `gorm:"foreignKey:PersonalInformationID"`
}
