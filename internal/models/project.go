package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ProjectCategoryID uint             `json:"project_category_id" binding:"required" gorm:"not null"`
	ProjectCategory   *ProjectCategory `gorm:"foreignKey:ProjectCategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title             string           `json:"title" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	Description       string           `json:"description" gorm:"type:text"`
	DescriptionDetail string           `json:"description_detail" gorm:"type:text"`
	RepoURL           string           `json:"repo_url" binding:"required,url" gorm:"type:varchar(255);not null"`
	DemoURL           string           `json:"demo_url" binding:"omitempty,url" gorm:"type:varchar(255)"`
	DemoVideo         string           `json:"demo_video" binding:"omitempty,url" gorm:"type:varchar(255)"`
	Thumbnail         string           `json:"thumbnail" binding:"omitempty,url" gorm:"type:varchar(255)"`
}
