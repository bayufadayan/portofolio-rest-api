package service

import (
	"errors"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/repository"
)

type ProjectService interface {
	GetAll() ([]models.Project, error)
	GetById(id uint) (*models.Project, error)
	GetByTitle(title string) ([]models.Project, error)
	Create(project *models.Project) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo}
}

func (s *projectService) GetAll() ([]models.Project, error) {
	return s.repo.GetAll()
}

func (s *projectService) GetById(id uint) (*models.Project, error) {
	return s.repo.GetById(id)
}

func (s *projectService) GetByTitle(title string) ([]models.Project, error) {
	return s.repo.GetByTitle(title)
}

func (s *projectService) Create(project *models.Project) error {
	return s.repo.Create(project)
}

func (s *projectService) Update(id uint, updates map[string]interface{}) error {
	existing, err := s.repo.GetById(id)
	if err != nil {
		return errors.New("project not found")
	}

	return s.repo.Update(existing.ID, updates)
}

func (s *projectService) Delete(id uint) error {
	return s.repo.Delete(id)
}