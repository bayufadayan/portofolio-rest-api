package repository

import (
	"portofolio-rest-api/internal/models"

	"gorm.io/gorm"
)

type EducationRepository interface {
	GetAll() ([]models.Education, error)
	GetById(id uint) (*models.Education, error)
	Create(education *models.Education) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type educationRepository struct {
	db *gorm.DB
}

func NewEducationRepository(db *gorm.DB) EducationRepository {
	return &educationRepository{db}
}

func (r *educationRepository) GetAll() ([]models.Education, error) {
	var educations []models.Education
	result := r.db.Find(&educations)
	return educations, result.Error
}

func (r *educationRepository) GetById(id uint) (*models.Education, error) {
	var education models.Education
	err := r.db.First(&education, id).Error
	if err != nil {
		return nil, err
	}
	return &education, nil
}

func (r *educationRepository) Delete(id uint) error {
    var education models.Education
    if err := r.db.First(&education, id).Error; err != nil {
        return err
    }

    return r.db.Delete(&education).Error
}

func (r *educationRepository) Create(education *models.Education) error {
	return r.db.Create(education).Error
}

func (r *educationRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.Education{}).Where("id = ?", id).Updates(updates).Error
}