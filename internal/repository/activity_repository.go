package repository

import (
	"portofolio-rest-api/internal/models"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	GetAll() ([]models.Activity, error)
	GetById(id uint) (*models.Activity, error)
	GetByTitle(title string) ([]models.Activity, error)
	GetByOrganizer(organizer string) ([]models.Activity, error)
	GetByType(tipe string) ([]models.Activity, error)
	Create(activity *models.Activity) error
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{db}
}

func (r *activityRepository) GetAll() ([]models.Activity, error) {
	var activities []models.Activity
	err := r.db.Find(&activities).Error
	return activities, err
}

func (r *activityRepository) GetById(id uint) (*models.Activity, error) {
	var activity models.Activity
	err := r.db.First(&activity, id).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *activityRepository) GetByTitle(title string) ([]models.Activity, error) {
	var activities []models.Activity
	searchPattern := "%" + title + "%"

	err := r.db.Where("LOWER(title) LIKE LOWER(?)", searchPattern).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *activityRepository) GetByOrganizer(organizer string) ([]models.Activity, error) {
	var activities []models.Activity
	searchPattern := "%" + organizer + "%"

	err := r.db.Where("LOWER(organizer) LIKE LOWER(?)", searchPattern).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *activityRepository) GetByType(tipe string) ([]models.Activity, error) {
	var activities []models.Activity
	searchPattern := "%" + tipe + "%"

	err := r.db.Where("LOWER(type::TEXT) LIKE LOWER(?)", searchPattern).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *activityRepository) Delete(id uint) error {
    var activity models.Activity
    if err := r.db.First(&activity, id).Error; err != nil {
        return err
    }

    return r.db.Delete(&activity).Error
}

func (r *activityRepository) Create(activity *models.Activity) error {
	return r.db.Create(activity).Error
}

func (r *activityRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.Activity{}).Where("id = ?", id).Updates(updates).Error
}