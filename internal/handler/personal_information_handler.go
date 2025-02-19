package handler

import (
	"net/http"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"

	"net/mail"
	"net/url"
)

type PersonalInformationHandler struct {
	service service.PersonalInformationService
}

func NewPersonalInformationHandler(service service.PersonalInformationService) *PersonalInformationHandler {
	return &PersonalInformationHandler{service}
}

func (h *PersonalInformationHandler) GetAll(c *gin.Context) {
	personalInformations, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, personalInformations)
}

func (h *PersonalInformationHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	personalInformation, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Personal Information not found"})
		return
	}

	c.JSON(http.StatusOK, personalInformation)
}

func (h *PersonalInformationHandler) Create(c *gin.Context) {
	var personalInformation models.PersonalInformation

	if err := c.ShouldBindJSON(&personalInformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&personalInformation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Personal information created successfully",
		"data":    personalInformation,
	})
}

func (h *PersonalInformationHandler) Update(c *gin.Context) {
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

	if email, ok := updates["email"]; ok {
		if !isValidEmail(email.(string)) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email format"})
			return
		}
	}

	if url, ok := updates["profile_image"]; ok {
		if !isValidURL(url.(string)) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid URL format"})
			return
		}
	}

	for key, value := range updates {
		if value == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": key + " cannot be empty"})
			return
		}
	}

	err = h.service.Update(uint(id), updates)
	if err != nil {
		switch err.Error() {
		case "personal information not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "personal information updated successfully"})
}

// Validasi email
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Validasi URL
func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func (h *PersonalInformationHandler) Delete(c *gin.Context) {
	id, err := getIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, err := h.service.GetById(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "personal information not found"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete personal information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "personal information deleted successfully"})
}