package service

import (
	"errors"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/repository"
)

type SocialMediaService interface {
	GetAll() ([]models.SocialMedia, error)
	GetById(id uint) (*models.SocialMedia, error)
	GetByPlatform(platform string) ([]models.SocialMedia, error)
	Create(socialMedia *models.SocialMedia) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type socialMediaService struct {
	repo repository.SocialMediaRepository
}

func NewSocialMediaService(repo repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{repo}
}

func (s *socialMediaService) GetAll() ([]models.SocialMedia, error) {
	return s.repo.GetAll()
}

func (s *socialMediaService) GetById(id uint) (*models.SocialMedia, error) {
	return s.repo.GetById(id)
}

func (s *socialMediaService) GetByPlatform(platform string) ([]models.SocialMedia, error) {
	return s.repo.GetByPlatform(platform)
}

func (s *socialMediaService) Create(socialMedia *models.SocialMedia) error {
	return s.repo.Create(socialMedia)
}

func (s *socialMediaService) Update(id uint, updates map[string]interface{}) error {
	existing, err := s.repo.GetById(id)
	if err != nil {
		return errors.New("social media not found")
	}

	if t, ok := updates["type"].(string); ok {
		validTypes := map[string]bool{"primary": true, "secondary": true, "tertiary": true}
		if !validTypes[t] {
			return errors.New("invalid type: must be 'primary', 'secondary', or 'tertiary'")
		}
	}

	return s.repo.Update(existing.ID, updates)
}

func (s *socialMediaService) Delete(id uint) error {
	return s.repo.Delete(id)
}
