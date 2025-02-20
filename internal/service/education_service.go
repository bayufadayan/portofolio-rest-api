package service

import (
	"errors"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/repository"
)

type EducationService interface {
	GetAll() ([]models.Education, error)
	GetById(id uint) (*models.Education, error)
	Create(education *models.Education) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type educationService struct {
	repo repository.EducationRepository
}

func NewEducationService(repo repository.EducationRepository) EducationService {
	return &educationService{repo}
}

func (s *educationService) GetAll() ([]models.Education, error) {
	return s.repo.GetAll()
}

func (s *educationService) GetById(id uint) (*models.Education, error) {
	return s.repo.GetById(id)
}

func (s *educationService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *educationService) Create(education *models.Education) error {
	return s.repo.Create(education)
}

func (s *educationService) Update(id uint, updates map[string]interface{}) error {
	existing, err := s.repo.GetById(id)
	if err != nil {
		return errors.New("education not found")
	}

	return s.repo.Update(existing.ID, updates)
}