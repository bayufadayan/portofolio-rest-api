package repository

import (
	"errors"
	"portofolio-rest-api/internal/models"

	"gorm.io/gorm"
)

type JobTitleRepository interface {
	GetAll() ([]models.JobTitle, error)
	GetById(id uint) (*models.JobTitle, error)
	Create(jobTitle *models.JobTitle) error
	GetByTitle(title string) (*models.JobTitle, error)
	UpdateTitle(id uint, title string) error
	Delete(id uint) error
}

func (r *jobTitleRepository) GetByTitle(title string) (*models.JobTitle, error) {
	var jobTitle models.JobTitle
	if err := r.db.Where("title = ?", title).First(&jobTitle).Error; err != nil {
		return nil, err
	}
	return &jobTitle, nil
}

type jobTitleRepository struct {
	db *gorm.DB
}

func NewJobTitleRepository(db *gorm.DB) JobTitleRepository {
	return &jobTitleRepository{db}
}

func (r *jobTitleRepository) GetAll() ([]models.JobTitle, error) {
	var jobTitles []models.JobTitle
	result := r.db.Find(&jobTitles)
	return jobTitles, result.Error
}

func (r *jobTitleRepository) GetById(id uint) (*models.JobTitle, error) {
	var jobTitle models.JobTitle
	result := r.db.First(&jobTitle, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &jobTitle, nil
}

func (r *jobTitleRepository) Create(jobTitle *models.JobTitle) error {
	var existing models.JobTitle
	if err := r.db.Where("title = ?", jobTitle.Title).First(&existing).Error; err == nil {
		return errors.New("job title already exists")
	}

	result := r.db.Create(jobTitle)
	return result.Error
}

func (r *jobTitleRepository) UpdateTitle(id uint, title string) error {
	return r.db.Model(&models.JobTitle{}).
		Where("id = ?", id).
		Update("title", title).
		Error
}

func (r *jobTitleRepository) Delete(id uint) error {
    var jobTitle models.JobTitle
    if err := r.db.First(&jobTitle, id).Error; err != nil {
        return err
    }

    return r.db.Delete(&jobTitle).Error
}

