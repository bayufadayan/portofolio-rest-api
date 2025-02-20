package service

import (
	"errors"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/repository"
)

type SkillService interface {
	GetAll() ([]models.Skill, error)
	GetById(id uint) (*models.Skill, error)
	GetByName(name string) ([]models.Skill, error)
	Create(skill *models.Skill) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type skillService struct {
	repo repository.SkillRepository
}

func NewSkillService(repo repository.SkillRepository) SkillService {
	return &skillService{repo}
}

func (s *skillService) GetAll() ([]models.Skill, error) {
	return s.repo.GetAll()
}

func (s *skillService) GetById(id uint) (*models.Skill, error) {
	return s.repo.GetById(id)
}

func (s *skillService) GetByName(name string) ([]models.Skill, error) {
	return s.repo.GetByName(name)
}

func (s *skillService) Create(skill *models.Skill) error {
	return s.repo.Create(skill)
}

func (s *skillService) Update(id uint, updates map[string]interface{}) error {
	existing, err := s.repo.GetById(id)
	if err != nil {
		return errors.New("skill not found")
	}

	return s.repo.Update(existing.ID, updates)
}

func (s *skillService) Delete(id uint) error {
	return s.repo.Delete(id)
}