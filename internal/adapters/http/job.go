package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type JobHandler struct {
	jobService *services.JobService
}

func NewJobHandler(jobService *services.JobService) *JobHandler {
	return &JobHandler{jobService: jobService}
}

func (h *JobHandler) CreateJob(c *gin.Context) {
	var job domain.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		log.Error("Failed to bind job data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.jobService.CreateJob(c.Request.Context(), &job); err != nil {
		log.Error("Failed to create job:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job"})
		return
	}

	log.Info("Job created successfully:", job.ID)
	c.JSON(http.StatusCreated, job)
}

func (h *JobHandler) GetJob(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid job ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	job, err := h.jobService.GetJobByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get job:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get job"})
		return
	}

	if job == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	log.Info("Job retrieved successfully:", job.ID)
	c.JSON(http.StatusOK, job)
}

func (h *JobHandler) UpdateJob(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid job ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	var job domain.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		log.Error("Failed to bind job data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	job.ID = uint(id)
	if err := h.jobService.UpdateJob(c.Request.Context(), &job); err != nil {
		log.Error("Failed to update job:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job"})
		return
	}

	log.Info("Job updated successfully:", job.ID)
	c.JSON(http.StatusOK, job)
}

func (h *JobHandler) DeleteJob(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid job ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	if err := h.jobService.DeleteJob(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete job:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job"})
		return
	}

	log.Info("Job deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *JobHandler) ListJobs(c *gin.Context) {
	jobs, err := h.jobService.ListJobs(c.Request.Context())
	if err != nil {
		log.Error("Failed to list jobs:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list jobs"})
		return
	}

	log.Info("Jobs listed successfully")
	c.JSON(http.StatusOK, jobs)
}
