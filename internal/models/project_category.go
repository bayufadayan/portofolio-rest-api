package models

import (
	"gorm.io/gorm"
)

type ProjectCategory struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100);not null"`
	Projects []Project `gorm:"foreignKey:ProjectCategoryID"`
}
