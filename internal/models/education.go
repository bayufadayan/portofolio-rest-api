package models

import (
	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	Institution  string `gorm:"type:varchar(100);not null"`
	Degree       string `gorm:"type:varchar(20)"`
	Major        string `gorm:"type:varchar(100)"`
	StartYear    int    `gorm:"type:int;not null"`
	EndYear      int    `gorm:"type:int"`
	Description  string `gorm:"type:text"`
}
