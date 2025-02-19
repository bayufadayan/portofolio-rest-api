package service

import (
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/repository"
)

type JobTitleService interface {
	GetAll() ([]models.JobTitle, error)
	GetById(id uint) (*models.JobTitle, error)
	Create(jobTitle *models.JobTitle) error
	GetByTitle(title string) (*models.JobTitle, error)
	UpdateTitle(id uint, title string) error
	Delete(id uint) error
}

func (s *jobTitleService) GetByTitle(title string) (*models.JobTitle, error) {
	return s.repo.GetByTitle(title)
}

type jobTitleService struct {
	repo repository.JobTitleRepository
}

func NewJobTitleService(repo repository.JobTitleRepository) JobTitleService {
	return &jobTitleService{repo}
}

func (s *jobTitleService) GetAll() ([]models.JobTitle, error) {
	return s.repo.GetAll()
}

func (s *jobTitleService) GetById(id uint) (*models.JobTitle, error) {
	return s.repo.GetById(id)
}

func (s *jobTitleService) Create(jobTitle *models.JobTitle) error {
	return s.repo.Create(jobTitle)
}

func (s *jobTitleService) UpdateTitle(id uint, title string) error {
	return s.repo.UpdateTitle(id, title)
}

func (s *jobTitleService) Delete(id uint) error {
    return s.repo.Delete(id)
}