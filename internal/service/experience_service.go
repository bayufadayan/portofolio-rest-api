package service

import (
	"errors"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/repository"
)

type ExperienceService interface {
	GetAll() ([]models.Experience, error)
	GetById(id uint) (*models.Experience, error)
	GetByTitle(title string) ([]models.Experience, error)
	GetByCompany(issuer string) ([]models.Experience, error)
	Create(experience *models.Experience) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type experienceService struct {
	repo repository.ExperienceRepository
}

func NewExperienceService(repo repository.ExperienceRepository) ExperienceService {
	return &experienceService{repo}
}

func (s *experienceService) GetAll() ([]models.Experience, error) {
	return s.repo.GetAll()
}

func (s *experienceService) GetById(id uint) (*models.Experience, error) {
	return s.repo.GetById(id)
}

func (s *experienceService) GetByTitle(title string) ([]models.Experience, error) {
	return s.repo.GetByTitle(title)
}

func (s *experienceService) GetByCompany(company string) ([]models.Experience, error) {
	return s.repo.GetByCompany(company)
}

func (s *experienceService) Create(experience *models.Experience) error {
	return s.repo.Create(experience)
}

func (s *experienceService) Update(id uint, updates map[string]interface{}) error {
	existing, err := s.repo.GetById(id)
	if err != nil {
		return errors.New("experience not found")
	}

	return s.repo.Update(existing.ID, updates)
}

func (s *experienceService) Delete(id uint) error {
	return s.repo.Delete(id)
}