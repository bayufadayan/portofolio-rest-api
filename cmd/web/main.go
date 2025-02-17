package main

import (
	"github.com/gin-gonic/gin"
	"portofolio-rest-api/internal/api"  // Sesuaikan dengan path project kamu
	"portofolio-rest-api/internal/service"
)

func main() {
	service.InitDB()

	r := gin.Default()

	api.InitRoutes(r)

	r.Run(":8080")
}
