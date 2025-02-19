package api

import (
	"net/http"

	"portofolio-rest-api/internal/handler"
	"portofolio-rest-api/internal/repository"
	"portofolio-rest-api/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handlers struct {
	JobTitleHandler            handler.JobTitleHandler
	SocialMediaHandler         handler.SocialMediaHandler
	PersonalInformationHandler handler.PersonalInformationHandler
}

func InitRoutes(db *gorm.DB) *gin.Engine {
	return setupRoutes(*initHandler(db))
}

func initHandler(db *gorm.DB) *Handlers {
	// JobTitle
	jobTitleRepository := repository.NewJobTitleRepository(db)
	jobTitleService := service.NewJobTitleService(jobTitleRepository)
	jobTitleHandler := handler.NewJobTitleHandler(jobTitleService)

	// SocialMedia
	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	socialMediaHandler := handler.NewSocialMediaHandler(socialMediaService)

	// SocialMedia
	personalInformationRepository := repository.NewPersonalInformationRepository(db)
	personalInformationService := service.NewPersonalInformationService(personalInformationRepository)
	personalInformationHandler := handler.NewPersonalInformationHandler(personalInformationService)

	return &Handlers{
		JobTitleHandler:            *jobTitleHandler,
		SocialMediaHandler:         *socialMediaHandler,
		PersonalInformationHandler: *personalInformationHandler,
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

	// SocialMedia
	socialMedia := api.Group("/social-medias")
	socialMedia.GET("", func(c *gin.Context) {
		platform := c.Query("platform")
		if platform != "" {
			handler.SocialMediaHandler.GetByPlatform(c)
		} else {
			handler.SocialMediaHandler.GetAll(c)
		}
	})
	socialMedia.GET("/:id", handler.SocialMediaHandler.GetById)
	socialMedia.POST("", handler.SocialMediaHandler.Create)
	socialMedia.PATCH("/:id", handler.SocialMediaHandler.Update)
	socialMedia.DELETE("/:id", handler.SocialMediaHandler.Delete)

	// PersonalInformation
	personalInformation := api.Group("/personal-informations")
	personalInformation.GET("", handler.PersonalInformationHandler.GetAll)
	personalInformation.GET("/:id", handler.PersonalInformationHandler.GetById)
	personalInformation.POST("", handler.PersonalInformationHandler.Create)
	personalInformation.PATCH("/:id", handler.PersonalInformationHandler.Update)
	personalInformation.DELETE("/:id", handler.PersonalInformationHandler.Delete)

	return router
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ping success",
	})
}
