package models

import (
	"gorm.io/gorm"
)

type Certificate struct {
	gorm.Model
	Title           string `gorm:"type:varchar(100);not null"`
	Issuer          string `gorm:"type:varchar(100);not null"`
	ImageURL        string `gorm:"type:varchar(255);not null"`
	CertificateLink string `gorm:"type:varchar(255);not null"`
}
