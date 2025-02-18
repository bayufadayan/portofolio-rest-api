package infrastructure

import (
	"log"
	"portofolio-rest-api/internal/models"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.PersonalInformation{},
		&models.SocialMedia{},
		&models.JobTitle{},
		&models.ProjectCategory{},
		&models.Project{},
		&models.Certificate{},
		&models.Skill{},
		&models.Experience{},
		&models.Activity{},
		&models.Education{},
		&models.Admin{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database Migrated Successfully!")
}
