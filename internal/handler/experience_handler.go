package handler

import (
	"net/http"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExperienceHandler struct {
	service service.ExperienceService
}

func NewExperienceHandler(service service.ExperienceService) *ExperienceHandler {
	return &ExperienceHandler{service}
}

func (c *ExperienceHandler) GetAll(ctx *gin.Context) {
	experiences, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, experiences)
}

func (h *ExperienceHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	experience, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "experience not found"})
		return
	}

	c.JSON(http.StatusOK, experience)
}

func (h *ExperienceHandler) GetByTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title parameter is required"})
		return
	}

	experiences, err := h.service.GetByTitle(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, experiences)
}

func (h *ExperienceHandler) GetByCompany(c *gin.Context) {
	company := c.Query("company")
	if company == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company parameter is required"})
		return
	}

	experiences, err := h.service.GetByCompany(company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, experiences)
}

func (h *ExperienceHandler) Create(c *gin.Context) {
	var experience models.Experience

	if err := c.ShouldBindJSON(&experience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&experience); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "experience created successfully",
		"data":    experience,
	})
}

func (h *ExperienceHandler) Update(c *gin.Context) {
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

	requiredFields := []string{"title", "company", "start_date"}
	for _, field := range requiredFields {
		if value, exists := updates[field]; exists && value == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": field + " cannot be empty"})
			return
		}
	}

	err = h.service.Update(uint(id), updates)
	if err != nil {
		switch err.Error() {
		case "experience not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "experience updated successfully"})
}

func (h *ExperienceHandler) Delete(c *gin.Context) {
	id, err := getIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, err := h.service.GetById(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "experience not found"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete experience"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "experience deleted successfully"})
}
