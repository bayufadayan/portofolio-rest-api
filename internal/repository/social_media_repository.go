package repository

import (
	"portofolio-rest-api/internal/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	GetAll() ([]models.SocialMedia, error)
	GetById(id uint) (*models.SocialMedia, error)
	GetByPlatform(platform string) ([]models.SocialMedia, error)
	Create(socialMedia *models.SocialMedia) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) GetAll() ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err := r.db.Find(&socialMedias).Error
	return socialMedias, err
}

func (r *socialMediaRepository) GetById(id uint) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := r.db.First(&socialMedia, id).Error
	if err != nil {
		return nil, err
	}
	return &socialMedia, nil
}

func (r *socialMediaRepository) GetByPlatform(platform string) ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	searchPattern := "%" + platform + "%"

	err := r.db.Where("LOWER(platform) LIKE LOWER(?)", searchPattern).Find(&socialMedias).Error
	if err != nil {
		return nil, err
	}
	return socialMedias, nil
}

func (r *socialMediaRepository) Create(socialMedia *models.SocialMedia) error {
	return r.db.Create(socialMedia).Error
}

func (r *socialMediaRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.SocialMedia{}).Where("id = ?", id).Updates(updates).Error
}

func (r *socialMediaRepository) Delete(id uint) error {
    var socialMedia models.SocialMedia
    if err := r.db.First(&socialMedia, id).Error; err != nil {
        return err
    }

    return r.db.Delete(&socialMedia).Error
}