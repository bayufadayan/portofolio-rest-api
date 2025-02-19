package models

import (
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	SkillName string `json:"skill_name" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	IconURL   string `json:"icon_url" binding:"required,url" gorm:"type:varchar(255);not null"`
}
