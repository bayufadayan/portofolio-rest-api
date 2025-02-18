package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ProjectCategoryID uint            `gorm:"not null"`
	ProjectCategory   ProjectCategory `gorm:"foreignKey:ProjectCategoryID"`
	Title             string          `gorm:"type:varchar(100);not null"`
	Description       string          `gorm:"type:text"`
	RepoURL           string          `gorm:"type:varchar(255);not null"`
	DemoURL           string          `gorm:"type:varchar(255)"`
	DemoVideo         string          `gorm:"type:varchar(255)"`
}
