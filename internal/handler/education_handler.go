package handler

import (
	"net/http"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EducationHandler struct {
	service service.EducationService
}

func NewEducationHandler(service service.EducationService) *EducationHandler {
	return &EducationHandler{service}
}

func (c *EducationHandler) GetAll(ctx *gin.Context) {
	educations, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, educations)
}

func (h *EducationHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	education, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "education not found"})
		return
	}

	c.JSON(http.StatusOK, education)
}

func (h *EducationHandler) Delete(c *gin.Context) {
	id, err := getIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, err := h.service.GetById(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "education not found"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete education"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "education deleted successfully"})
}

func (h *EducationHandler) Create(c *gin.Context) {
	var education models.Education

	if err := c.ShouldBindJSON(&education); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&education); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "education created successfully",
		"data":    education,
	})
}

func (h *EducationHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requiredFields := []string{"institution", "start_year"}
	for _, field := range requiredFields {
		if value, exists := updates[field]; exists && value == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": field + " cannot be empty"})
			return
		}
	}

	err = h.service.Update(uint(id), updates)
	if err != nil {
		switch err.Error() {
		case "education not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "education updated successfully"})
}
