package repository

import (
	"portofolio-rest-api/internal/models"

	"gorm.io/gorm"
)

type ProjectCategoryRespository interface {
	GetAll() ([]models.ProjectCategory, error)
	GetById(id uint) (*models.ProjectCategory, error)
	Create(projectCategory *models.ProjectCategory) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type projectCategoryRepository struct {
	db *gorm.DB
}

func NewProjectCategoryRepository(db *gorm.DB) ProjectCategoryRespository  {
	return &projectCategoryRepository{db}
}

func (r *projectCategoryRepository) GetAll() ([]models.ProjectCategory, error) {
	var projectCategories []models.ProjectCategory
	err := r.db.Find(&projectCategories).Error
	return projectCategories, err
}

func (r *projectCategoryRepository) GetById(id uint) (*models.ProjectCategory, error) {
	var projectCategory models.ProjectCategory
	err := r.db.First(&projectCategory, id).Error
	if err != nil {
		return nil, err
	}
	return &projectCategory, nil
}

func (r *projectCategoryRepository) Create(projectCategory *models.ProjectCategory) error {
	return r.db.Create(projectCategory).Error
}

func (r *projectCategoryRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.ProjectCategory{}).Where("id = ?", id).Updates(updates).Error
}

func (r *projectCategoryRepository) Delete(id uint) error {
    var projectCategory models.ProjectCategory
    if err := r.db.First(&projectCategory, id).Error; err != nil {
        return err
    }

    return r.db.Delete(&projectCategory).Error
}