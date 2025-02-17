package repository

import (
	"portofolio-rest-api/internal/model"

	"gorm.io/gorm"
)

func GetPersonalInformation(db *gorm.DB) (*model.PersonalInformation, error) {
	var personalInfo model.PersonalInformation
	err := db.First(&personalInfo).Error
	if err != nil {
		return nil, err
	}
	return &personalInfo, nil
}
