package models

import (
	"gorm.io/gorm"
)

type Certificate struct {
	gorm.Model
	Title           string `json:"title" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	Issuer          string `json:"issuer" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	ImageURL        string `json:"image_url" binding:"required,url" gorm:"type:varchar(255);not null"`
	CertificateLink string `json:"certificate_link" binding:"required,url" gorm:"type:varchar(255);not null"`
}
