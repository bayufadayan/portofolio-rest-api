package repository

import (
	"portofolio-rest-api/internal/models"

	"gorm.io/gorm"
)

type SkillRepository interface {
	GetAll() ([]models.Skill, error)
	GetById(id uint) (*models.Skill, error)
	GetByName(name string) ([]models.Skill, error)
	Create(skill *models.Skill) error
	Update(id uint, updates map[string]interface{}) error
	// delete
}

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepository{db}
}

func (r *skillRepository) GetAll() ([]models.Skill, error) {
	var skills []models.Skill
	result := r.db.Find(&skills)
	return skills, result.Error
}

func (r *skillRepository) GetById(id uint) (*models.Skill, error) {
	var skill models.Skill
	err := r.db.First(&skill, id).Error
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

func (r *skillRepository) GetByName(name string) ([]models.Skill, error) {
	var skills []models.Skill
	searchPattern := "%" + name + "%"

	err := r.db.Where("LOWER(skill_name) LIKE LOWER(?)", searchPattern).Find(&skills).Error
	if err != nil {
		return nil, err
	}
	return skills, nil
}

func (r *skillRepository) Create(skill *models.Skill) error {
	return r.db.Create(skill).Error
}

func (r *skillRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.Skill{}).Where("id = ?", id).Updates(updates).Error
}

