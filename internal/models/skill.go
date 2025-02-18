package models

import (
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	SkillName string `gorm:"type:varchar(100);not null"`
	IconURL   string `gorm:"type:varchar(255);not null"`
}
