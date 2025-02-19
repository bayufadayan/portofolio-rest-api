package models

import (
	"gorm.io/gorm"
)

type ActivityType string

const (
	Training    ActivityType = "pelatihan"
	Competition ActivityType = "kompetisi"
	Seminar     ActivityType = "seminar"
	Volunteer   ActivityType = "volunteer"
)

type Activity struct {
	gorm.Model
	Title       string       `json:"title" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	Organizer   string       `json:"organizer" binding:"required,max=100" gorm:"type:varchar(100);not null"`
	Type        ActivityType `json:"type" binding:"required" gorm:"type:activity_type;not null"`
	Achievement string       `json:"achievement" gorm:"type:varchar(255)"`
	Date        string       `json:"date" binding:"required" gorm:"type:date;not null"`
	Description string       `json:"description" gorm:"type:text"`
}
