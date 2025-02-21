package repository

import (
	"portofolio-rest-api/internal/models"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	GetAll() ([]models.Project, error)
	GetById(id uint) (*models.Project, error)
	GetByTitle(title string) ([]models.Project, error)
	Create(project *models.Project) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db}
}

func (r *projectRepository) GetAll() ([]models.Project, error) {
	var projects []models.Project
	result := r.db.Preload("ProjectCategory").Find(&projects)
	return projects, result.Error
}

func (r *projectRepository) GetById(id uint) (*models.Project, error) {
	var project models.Project
	err := r.db.Preload("ProjectCategory").First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) GetByTitle(title string) ([]models.Project, error) {
	var projects []models.Project
	searchPattern := "%" + title + "%"

	err := r.db.Where("LOWER(title) LIKE LOWER(?)", searchPattern).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *projectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *projectRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.Project{}).Where("id = ?", id).Updates(updates).Error
}

func (r *projectRepository) Delete(id uint) error {
    var project models.Project
    if err := r.db.First(&project, id).Error; err != nil {
        return err
    }

    return r.db.Delete(&project).Error
}
