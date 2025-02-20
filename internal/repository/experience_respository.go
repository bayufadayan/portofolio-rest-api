package repository

import (
	"portofolio-rest-api/internal/models"

	"gorm.io/gorm"
)

type ExperienceRepository interface {
	GetAll() ([]models.Experience, error)
	GetById(id uint) (*models.Experience, error)
	GetByTitle(title string) ([]models.Experience, error)
	GetByCompany(issuer string) ([]models.Experience, error)
	Create(experience *models.Experience) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type experienceRepository struct {
	db *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) ExperienceRepository {
	return &experienceRepository{db}
}

func (r *experienceRepository) GetAll() ([]models.Experience, error) {
	var experiences []models.Experience
	result := r.db.Find(&experiences)
	return experiences, result.Error
}

func (r *experienceRepository) GetById(id uint) (*models.Experience, error) {
	var experience models.Experience
	err := r.db.First(&experience, id).Error
	if err != nil {
		return nil, err
	}
	return &experience, nil
}

func (r *experienceRepository) GetByTitle(title string) ([]models.Experience, error) {
	var experiences []models.Experience
	searchPattern := "%" + title + "%"

	err := r.db.Where("LOWER(title) LIKE LOWER(?)", searchPattern).Find(&experiences).Error
	if err != nil {
		return nil, err
	}
	return experiences, nil
}

func (r *experienceRepository) GetByCompany(company string) ([]models.Experience, error) {
	var experiences []models.Experience
	searchPattern := "%" + company + "%"

	err := r.db.Where("LOWER(company) LIKE LOWER(?)", searchPattern).Find(&experiences).Error
	if err != nil {
		return nil, err
	}
	return experiences, nil
}

func (r *experienceRepository) Create(experience *models.Experience) error {
	return r.db.Create(experience).Error
}

func (r *experienceRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.Experience{}).Where("id = ?", id).Updates(updates).Error
}

func (r *experienceRepository) Delete(id uint) error {
    var experience models.Experience
    if err := r.db.First(&experience, id).Error; err != nil {
        return err
    }

    return r.db.Delete(&experience).Error
}