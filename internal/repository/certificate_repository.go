package repository

import (
	"portofolio-rest-api/internal/models"

	"gorm.io/gorm"
)

type CertificateRepository interface {
	GetAll() ([]models.Certificate, error)
	GetById(id uint) (*models.Certificate, error)
	GetByTitle(title string) ([]models.Certificate, error)
	GetByIssuer(issuer string) ([]models.Certificate, error)
	Create(certificate *models.Certificate) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type certificateRepository struct {
	db *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) CertificateRepository {
	return &certificateRepository{db}
}

func (r *certificateRepository) GetAll() ([]models.Certificate, error) {
	var certificates []models.Certificate
	err := r.db.Find(&certificates).Error
	return certificates, err
}

func (r *certificateRepository) GetById(id uint) (*models.Certificate, error) {
	var certificate models.Certificate
	err := r.db.First(&certificate, id).Error
	if err != nil {
		return nil, err
	}
	return &certificate, nil
}

func (r *certificateRepository) GetByTitle(title string) ([]models.Certificate, error) {
	var certificates []models.Certificate
	searchPattern := "%" + title + "%"

	err := r.db.Where("LOWER(title) LIKE LOWER(?)", searchPattern).Find(&certificates).Error
	if err != nil {
		return nil, err
	}
	return certificates, nil
}

func (r *certificateRepository) GetByIssuer(issuer string) ([]models.Certificate, error) {
	var certificates []models.Certificate
	searchPattern := "%" + issuer + "%"

	err := r.db.Where("LOWER(issuer) LIKE LOWER(?)", searchPattern).Find(&certificates).Error
	if err != nil {
		return nil, err
	}
	return certificates, nil
}

func (r *certificateRepository) Create(certificate *models.Certificate) error {
	return r.db.Create(certificate).Error
}

func (r *certificateRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.Certificate{}).Where("id = ?", id).Updates(updates).Error
}

func (r *certificateRepository) Delete(id uint) error {
    var certificate models.Certificate
    if err := r.db.First(&certificate, id).Error; err != nil {
        return err
    }

    return r.db.Delete(&certificate).Error
}