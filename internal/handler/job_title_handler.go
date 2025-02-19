package handler

import (
	"net/http"
	"portofolio-rest-api/internal/models"
	"portofolio-rest-api/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobTitleRequest struct {
	Title string `json:"title" binding:"required"`
}

type JobTitleHandler struct {
	service service.JobTitleService
}

func NewJobTitleHandler(service service.JobTitleService) *JobTitleHandler {
	return &JobTitleHandler{service}
}

func (c *JobTitleHandler) GetAll(ctx *gin.Context) {
	jobTitles, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jobTitles)
}

func (c *JobTitleHandler) GetById(ctx *gin.Context) {
	id, err := getIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jobTitle, err := c.service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Job title not found"})
		return
	}
	ctx.JSON(http.StatusOK, jobTitle)
}

func getIDParam(ctx *gin.Context) (uint, error) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (c *JobTitleHandler) Create(ctx *gin.Context) {
	var req JobTitleRequest

	// Bind JSON ke struct (otomatis validasi "title" harus ada)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing, _ := c.service.GetByTitle(req.Title)
	if existing != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Job title already exists"})
		return
	}

	jobTitle := models.JobTitle{Title: req.Title}
	err := c.service.Create(&jobTitle)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Job title created", "data": jobTitle})
}

func (h *JobTitleHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var request struct {
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	existingJob, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job title not found"})
		return
	}

	duplicateJob, err := h.service.GetByTitle(request.Title)
	if err == nil && duplicateJob.ID != existingJob.ID {
		c.JSON(http.StatusConflict, gin.H{"error": "Job title already exists"})
		return
	}

	if err := h.service.UpdateTitle(uint(id), request.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job title"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job title updated successfully"})
}

func (h *JobTitleHandler) Delete(c *gin.Context) {
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

