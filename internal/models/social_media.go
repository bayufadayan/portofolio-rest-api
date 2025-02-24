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
	Platform string          `json:"platform" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	Type     SocialMediaType `json:"type" binding:"required" gorm:"type:social_media_type;not null"`
	Icon     string          `json:"icon" binding:"required" gorm:"type:varchar(255);not null"`
	URL      string          `json:"url" binding:"required,url" gorm:"type:varchar(255);not null"`
}
