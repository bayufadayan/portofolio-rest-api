package main

import (
	"portofolio-rest-api/infrastructure"
	"portofolio-rest-api/internal/api"
)

func main() {
	// Koneksi Database
	db := infrastructure.NewDBConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Setup Routes
	router := api.InitRoutes(db)
	router.Run(":8080")
}
