package handler

import (
	"net/http"
	"strconv"

	// "strconv"

	// "portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/service"

	"github.com/gin-gonic/gin"
)

type CertificateHandler struct {
	service service.CertificateService
}

func NewCertificateHandler(service service.CertificateService) *CertificateHandler {
	return &CertificateHandler{service}
}

func (c *CertificateHandler) GetAll(ctx *gin.Context) {
	certificates, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, certificates)
}

func (h *CertificateHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	certificate, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "certificate not found"})
		return
	}

	c.JSON(http.StatusOK, certificate)
}

func (h *CertificateHandler) GetByTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Platform parameter is required"})
		return
	}

	certificates, err := h.service.GetByTitle(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, certificates)
}

func (h *CertificateHandler) GetByIssuer(c *gin.Context) {
	issuer := c.Query("issuer")
	if issuer == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Platform parameter is required"})
		return
	}

	certificates, err := h.service.GetByIssuer(issuer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, certificates)
}

func (h *CertificateHandler) Create(c *gin.Context) {
	var certicate models.Certificate

	if err := c.ShouldBindJSON(&certicate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&certicate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "certificate created successfully",
		"data":    certicate,
	})
}

func (h *CertificateHandler) Update(c *gin.Context) {
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

	urlFields := []string{"image_url", "certificate_link"}
	for _, field := range urlFields {
		if value, ok := updates[field]; ok {
			if !isValidURL(value.(string)) {
				c.JSON(http.StatusBadRequest, gin.H{"error": field + " must be a valid URL"})
				return
			}
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
		case "certificate not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "certificate updated successfully"})
}

func (h *CertificateHandler) Delete(c *gin.Context) {
	id, err := getIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, err := h.service.GetById(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "certicates not found"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete certificate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "certificate deleted successfully"})
}
