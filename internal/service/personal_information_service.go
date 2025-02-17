package service

import (
	"log"
	"portofolio-rest-api/internal/model"
	"portofolio-rest-api/internal/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdbname port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}
}

func GetPersonalInformation() (*model.PersonalInformation, error) {
	// Pastikan database sudah terhubung
	if DB == nil {
		InitDB()
	}

	// Get data dari repository
	personalInfo, err := repository.GetPersonalInformation(DB)
	if err != nil {
		return nil, err
	}

	return personalInfo, nil
}
