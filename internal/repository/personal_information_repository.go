package repository

import (
	"portofolio-rest-api/internal/models"

	"gorm.io/gorm"
)

type PersonalInformationRespository interface {
	GetAll() ([]models.PersonalInformation, error)
	GetById(id uint) (*models.PersonalInformation, error)
	Create(personalInformation *models.PersonalInformation) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type personalInformationRepository struct {
	db *gorm.DB
}

func NewPersonalInformationRepository(db *gorm.DB) PersonalInformationRespository  {
	return &personalInformationRepository{db}
}

func (r *personalInformationRepository) GetAll() ([]models.PersonalInformation, error) {
	var personalInformations []models.PersonalInformation
	err := r.db.Find(&personalInformations).Error
	return personalInformations, err
}

func (r *personalInformationRepository) GetById(id uint) (*models.PersonalInformation, error) {
	var personalInformation models.PersonalInformation
	err := r.db.First(&personalInformation, id).Error
	if err != nil {
		return nil, err
	}
	return &personalInformation, nil
}

func (r *personalInformationRepository) Create(personalInformation *models.PersonalInformation) error {
	return r.db.Create(personalInformation).Error
}

func (r *personalInformationRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.PersonalInformation{}).Where("id = ?", id).Updates(updates).Error
}

func (r *personalInformationRepository) Delete(id uint) error {
    var personalInformation models.PersonalInformation
    if err := r.db.First(&personalInformation, id).Error; err != nil {
        return err
    }

    return r.db.Delete(&personalInformation).Error
}