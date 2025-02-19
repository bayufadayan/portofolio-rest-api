package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"portofolio-rest-api/internal/handler"
	"portofolio-rest-api/internal/repository"
	"portofolio-rest-api/internal/service"
)

type Handlers struct {
	JobTitleHandler handler.JobTitleHandler
}

func InitRoutes(db *gorm.DB) *gin.Engine {
	return setupRoutes(*initHandler(db))
}

func initHandler(db *gorm.DB) *Handlers {
	// JobTitle
	jobTitleRepository := repository.NewJobTitleRepository(db)
	jobTitleService := service.NewJobTitleService(jobTitleRepository)
	jobTitleHandler := handler.NewJobTitleHandler(jobTitleService)

	return &Handlers{
		JobTitleHandler: *jobTitleHandler,
	}
}

func setupRoutes(handler Handlers) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	api.GET("/ping", pingHandler)

	// JobTitle
	jobTitle := api.Group("/job-titles")
	jobTitle.GET("", handler.JobTitleHandler.GetAll)
	jobTitle.GET("/:id", handler.JobTitleHandler.GetById)
	jobTitle.POST("", handler.JobTitleHandler.Create)
	jobTitle.PATCH("/:id", handler.JobTitleHandler.Update)
	jobTitle.DELETE("/:id", handler.JobTitleHandler.Delete)

	return router
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ping success",
	})
}
