package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type JobStatusHistoryHandler struct {
	jobStatusHistoryService *services.JobStatusHistoryService
}

func NewJobStatusHistoryHandler(jobStatusHistoryService *services.JobStatusHistoryService) *JobStatusHistoryHandler {
	return &JobStatusHistoryHandler{jobStatusHistoryService: jobStatusHistoryService}
}

func (h *JobStatusHistoryHandler) CreateJobStatusHistory(c *gin.Context) {
	var history domain.JobStatusHistory
	if err := c.ShouldBindJSON(&history); err != nil {
		log.Error("Failed to bind job status history data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.jobStatusHistoryService.CreateJobStatusHistory(c.Request.Context(), &history); err != nil {
		log.Error("Failed to create job status history:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job status history"})
		return
	}

	log.Info("Job status history created successfully:", history.ID)
	c.JSON(http.StatusCreated, history)
}

func (h *JobStatusHistoryHandler) ListJobStatusHistoryByJobID(c *gin.Context) {
	jobID, err := strconv.ParseUint(c.Param("job_id"), 10, 64)
	if err != nil {
		log.Error("Invalid job ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	history, err := h.jobStatusHistoryService.ListJobStatusHistoryByJobID(c.Request.Context(), uint(jobID))
	if err != nil {
		log.Error("Failed to list job status history:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list job status history"})
		return
	}

	log.Info("Job status history listed successfully for job:", jobID)
	c.JSON(http.StatusOK, history)
}
