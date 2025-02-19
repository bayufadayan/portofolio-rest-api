package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ProjectCategoryID uint            `json:"category" binding:"required" gorm:"not null"`
	ProjectCategory   ProjectCategory `gorm:"foreignKey:ProjectCategoryID"`
	Title             string          `json:"title" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	Description       string          `json:"description" gorm:"type:text"`
	RepoURL           string          `json:"repo_url" binding:"required,url" gorm:"type:varchar(255);not null"`
	DemoURL           string          `json:"demo_url" binding:"url" gorm:"type:varchar(255)"`
	DemoVideo         string          `json:"demo_video" binding:"url" gorm:"type:varchar(255)"`
}
