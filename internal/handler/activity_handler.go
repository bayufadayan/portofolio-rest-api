package handler

import (
	"net/http"
	"strconv"

	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/service"

	"github.com/gin-gonic/gin"
)

type ActivityHandler struct {
	service service.ActivityService
}

func NewActivityHandler(service service.ActivityService) *ActivityHandler {
	return &ActivityHandler{service}
}

func (h *ActivityHandler) GetAll(c *gin.Context) {
	activities, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activities)
}

func (h *ActivityHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	activity, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "activity not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func (h *ActivityHandler) GetByTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title parameter is required"})
		return
	}

	activities, err := h.service.GetByTitle(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func (h *ActivityHandler) GetByOrganizer(c *gin.Context) {
	organizer := c.Query("organizer")
	if organizer == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Organizer parameter is required"})
		return
	}

	activities, err := h.service.GetByOrganizer(organizer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func (h *ActivityHandler) GetByType(c *gin.Context) {
	tipe := c.Query("type")
	if tipe == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type parameter is required"})
		return
	}

	activities, err := h.service.GetByType(tipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func (h *ActivityHandler) Delete(c *gin.Context) {
	id, err := getIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, err := h.service.GetById(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "activity not found"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "activity deleted successfully"})
}

func (h *ActivityHandler) Create(c *gin.Context) {
	var activity models.Activity

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidActivityType(activity.Type) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type value. Must be 'pelatihan', 'kompetisi', 'seminar', or 'volunter'"})
		return
	}

	if err := h.service.Create(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Activity created successfully",
		"data":    activity,
	})
}

func isValidActivityType(t models.ActivityType) bool {
	switch t {
	case models.Competition, models.Training, models.Seminar, models.Volunteer:
		return true
	default:
		return false
	}
}

func (h *ActivityHandler) Update(c *gin.Context) {
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

	requiredFields := []string{"title", "organizer", "type", "date"}
	for _, field := range requiredFields {
		if value, exists := updates[field]; exists && value == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": field + " cannot be empty"})
			return
		}
	}

	err = h.service.Update(uint(id), updates)
	if err != nil {
		switch err.Error() {
		case "activity not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "invalid type: must be 'pelatihan', 'seminar', 'kompetisi' or 'volunteer'":
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "activity updated successfully"})
}
