package service

import (
	"errors"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/repository"
)

type ActivityService interface {
	GetAll() ([]models.Activity, error)
	GetById(id uint) (*models.Activity, error)
	GetByTitle(title string) ([]models.Activity, error)
	GetByOrganizer(organizer string) ([]models.Activity, error)
	GetByType(tipe string) ([]models.Activity, error)
	Create(activity *models.Activity) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type activityService struct {
	repo repository.ActivityRepository
}

func NewActivityService(repo repository.ActivityRepository) ActivityService {
	return &activityService{repo}
}


func (s *activityService) GetAll() ([]models.Activity, error) {
	return s.repo.GetAll()
}

func (s *activityService) GetById(id uint) (*models.Activity, error) {
	return s.repo.GetById(id)
}

func (s *activityService) GetByTitle(title string) ([]models.Activity, error) {
	return s.repo.GetByTitle(title)
}

func (s *activityService) GetByOrganizer(organizer string) ([]models.Activity, error) {
	return s.repo.GetByOrganizer(organizer)
}

func (s *activityService) GetByType(tipe string) ([]models.Activity, error) {
	return s.repo.GetByType(tipe)
}

func (s *activityService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *activityService) Create(activity *models.Activity) error {
	return s.repo.Create(activity)
}

func (s *activityService) Update(id uint, updates map[string]interface{}) error {
	existing, err := s.repo.GetById(id)
	if err != nil {
		return errors.New("activity not found")
	}

	if t, ok := updates["type"].(string); ok {
		validTypes := map[string]bool{"pelatihan": true, "seminar": true, "kompetisi": true, "volunteer": true}
		if !validTypes[t] {
			return errors.New("invalid type: must be 'pelatihan', 'seminar', 'kompetisi' or 'volunteer'")
		}
	}

	return s.repo.Update(existing.ID, updates)
}