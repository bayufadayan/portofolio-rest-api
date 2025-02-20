package service

import (
	"errors"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/repository"
)

type ProjectCategoryService interface {
	GetAll() ([]models.ProjectCategory, error)
	GetById(id uint) (*models.ProjectCategory, error)
	Create(projectCategory *models.ProjectCategory) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type projectCategoryService struct {
	repo repository.ProjectCategoryRespository
}

func NewProjectCategoryService(repo repository.ProjectCategoryRespository) ProjectCategoryService {
	return &projectCategoryService{repo}
}

func (s *projectCategoryService) GetAll() ([]models.ProjectCategory, error) {
	return s.repo.GetAll()
}

func (s *projectCategoryService) GetById(id uint) (*models.ProjectCategory, error) {
	return s.repo.GetById(id)
}

func (s *projectCategoryService) Create(projectCategory *models.ProjectCategory) error {
	return s.repo.Create(projectCategory)
}

func (s *projectCategoryService) Update(id uint, updates map[string]interface{}) error {
	existing, err := s.repo.GetById(id)
	if err != nil {
		return errors.New("personal information not found")
	}

	return s.repo.Update(existing.ID, updates)
}

func (s *projectCategoryService) Delete(id uint) error {
	return s.repo.Delete(id)
}