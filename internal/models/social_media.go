package models

import (
	"gorm.io/gorm"
)

type SocialMediaType string

const (
	Primary   SocialMediaType = "primary"
	Secondary SocialMediaType = "secondary"
	Tertiary  SocialMediaType = "tertiary"
)

type SocialMedia struct {
	gorm.Model
	Platform              string          `gorm:"type:varchar(100);not null"`
	Type                  SocialMediaType `gorm:"type:social_media_type;not null"`
	URL                   string          `gorm:"type:varchar(255);not null"`
	PersonalInformationID uint            `gorm:"not null"`
}
