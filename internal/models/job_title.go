package models

import (
	"gorm.io/gorm"
)

type JobTitle struct {
	gorm.Model
	Title string `json:"title" binding:"required" gorm:"type:varchar(100);not null;unique"`
}
