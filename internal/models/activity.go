package models

import (
	"gorm.io/gorm"
)

type ActivityType string

const (
	Training   ActivityType = "pelatihan"
	Competition ActivityType = "kompetisi"
	Seminar    ActivityType = "seminar"
	Volunteer  ActivityType = "volunteer"
)

type Activity struct {
	gorm.Model
	Title       string       `gorm:"type:varchar(100);not null"`
	Organizer   string       `gorm:"type:varchar(100);not null"`
	Type        ActivityType `gorm:"type:activity_type;not null"`
	Achievement string       `gorm:"type:varchar(255)"`
	Date        string       `gorm:"type:date;not null"`
	Description string       `gorm:"type:text"`
}
