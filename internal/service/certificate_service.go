package service

import (
	// "errors"
	"errors"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/repository"
)

type CertificateService interface {
	GetAll() ([]models.Certificate, error)
	GetById(id uint) (*models.Certificate, error)
	GetByTitle(title string) ([]models.Certificate, error)
	GetByIssuer(issuer string) ([]models.Certificate, error)
	Create(certificate *models.Certificate) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type certificateService struct {
	repo repository.CertificateRepository
}

func NewCertificateService(repo repository.CertificateRepository) CertificateService {
	return &certificateService{repo}
}

func (s *certificateService) GetAll() ([]models.Certificate, error) {
	return s.repo.GetAll()
}

func (s *certificateService) GetById(id uint) (*models.Certificate, error) {
	return s.repo.GetById(id)
}

func (s *certificateService) GetByTitle(title string) ([]models.Certificate, error) {
	return s.repo.GetByTitle(title)
}

func (s *certificateService) GetByIssuer(issuer string) ([]models.Certificate, error) {
	return s.repo.GetByIssuer(issuer)
}

func (s *certificateService) Create(certificate *models.Certificate) error {
	return s.repo.Create(certificate)
}

func (s *certificateService) Update(id uint, updates map[string]interface{}) error {
	existing, err := s.repo.GetById(id)
	if err != nil {
		return errors.New("certificate not found")
	}

	return s.repo.Update(existing.ID, updates)
}

func (s *certificateService) Delete(id uint) error {
	return s.repo.Delete(id)
}