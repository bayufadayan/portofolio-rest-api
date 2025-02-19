package service

import (
	"errors"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/repository"
)

type PersonalInformationService interface {
	GetAll() ([]models.PersonalInformation, error)
	GetById(id uint) (*models.PersonalInformation, error)
	Create(personalInformation *models.PersonalInformation) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type personalInformationService struct {
	repo repository.PersonalInformationRespository
}

func NewPersonalInformationService(repo repository.PersonalInformationRespository) PersonalInformationService {
	return &personalInformationService{repo}
}

func (s *personalInformationService) GetAll() ([]models.PersonalInformation, error) {
	return s.repo.GetAll()
}

func (s *personalInformationService) GetById(id uint) (*models.PersonalInformation, error) {
	return s.repo.GetById(id)
}

func (s *personalInformationService) Create(personalInformation *models.PersonalInformation) error {
	return s.repo.Create(personalInformation)
}

func (s *personalInformationService) Update(id uint, updates map[string]interface{}) error {
	existing, err := s.repo.GetById(id)
	if err != nil {
		return errors.New("personal information not found")
	}

	return s.repo.Update(existing.ID, updates)
}

func (s *personalInformationService) Delete(id uint) error {
	return s.repo.Delete(id)
}