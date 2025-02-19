package handler

import (
	"net/http"
	"strconv"

	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/service"

	"github.com/gin-gonic/gin"
)

type SocialMediaHandler struct {
	service service.SocialMediaService
}

func NewSocialMediaHandler(service service.SocialMediaService) *SocialMediaHandler {
	return &SocialMediaHandler{service}
}

func (h *SocialMediaHandler) GetAll(c *gin.Context) {
	socialMedias, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, socialMedias)
}

func (h *SocialMediaHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	socialMedia, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social media not found"})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func (h *SocialMediaHandler) GetByPlatform(c *gin.Context) {
	platform := c.Query("platform")
	if platform == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Platform parameter is required"})
		return
	}

	socialMedias, err := h.service.GetByPlatform(platform)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, socialMedias)
}

func (h *SocialMediaHandler) Create(c *gin.Context) {
	var socialMedia models.SocialMedia

	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidSocialMediaType(socialMedia.Type) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type value. Must be 'primary', 'secondary', or 'tertiary'"})
		return
	}

	if err := h.service.Create(&socialMedia); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Social media created successfully",
		"data":    socialMedia,
	})
}

func isValidSocialMediaType(t models.SocialMediaType) bool {
	switch t {
	case models.Primary, models.Secondary, models.Tertiary:
		return true
	default:
		return false
	}
}

func (h *SocialMediaHandler) Update(c *gin.Context) {
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

	err = h.service.Update(uint(id), updates)
	if err != nil {
		switch err.Error() {
		case "social media not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "invalid type: must be 'primary', 'secondary', or 'tertiary'":
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Social media updated successfully"})
}

func (h *SocialMediaHandler) Delete(c *gin.Context) {
	id, err := getIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, err := h.service.GetById(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job title not found"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job title"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job title deleted successfully"})
}
