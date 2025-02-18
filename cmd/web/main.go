package main

import (
	"log"
	"portofolio-rest-api/infrastructure"
	"portofolio-rest-api/internal/api"
)

func main() {
	db := infrastructure.NewDBConnection()
	if db == nil {
		log.Fatal("Failed to connect to database")
	}
	// untuk menjalankan query sql native (tanpa gorm), karena ngga perlu di komenin dlu aja
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatal("Failed to get database instance:", err)
	// }
	// defer sqlDB.Close()

	infrastructure.MigrateDB(db)

	router := api.InitRoutes(db)
	log.Println("Server running on :8080")
	router.Run(":8080")
}
