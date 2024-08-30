package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type ScheduleHandler struct {
	scheduleService *services.ScheduleService
}

func NewScheduleHandler(scheduleService *services.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{scheduleService: scheduleService}
}

func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	var schedule domain.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		log.Error("Failed to bind schedule data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.scheduleService.CreateSchedule(c.Request.Context(), &schedule); err != nil {
		log.Error("Failed to create schedule:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule"})
		return
	}

	log.Info("Schedule created successfully:", schedule.ID)
	c.JSON(http.StatusCreated, schedule)
}

func (h *ScheduleHandler) GetSchedule(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid schedule ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	schedule, err := h.scheduleService.GetScheduleByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get schedule:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get schedule"})
		return
	}

	if schedule == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	log.Info("Schedule retrieved successfully:", schedule.ID)
	c.JSON(http.StatusOK, schedule)
}

func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid schedule ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	var schedule domain.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		log.Error("Failed to bind schedule data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	schedule.ID = uint(id)
	if err := h.scheduleService.UpdateSchedule(c.Request.Context(), &schedule); err != nil {
		log.Error("Failed to update schedule:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}

	log.Info("Schedule updated successfully:", schedule.ID)
	c.JSON(http.StatusOK, schedule)
}

func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid schedule ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	if err := h.scheduleService.DeleteSchedule(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete schedule:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule"})
		return
	}

	log.Info("Schedule deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *ScheduleHandler) ListSchedulesByTechnicianID(c *gin.Context) {
	technicianID, err := strconv.ParseUint(c.Param("technician_id"), 10, 64)
	if err != nil {
		log.Error("Invalid technician ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid technician ID"})
		return
	}

	schedules, err := h.scheduleService.ListSchedulesByTechnicianID(c.Request.Context(), uint(technicianID))
	if err != nil {
		log.Error("Failed to list schedules:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list schedules"})
		return
	}

	log.Info("Schedules listed successfully for technician:", technicianID)
	c.JSON(http.StatusOK, schedules)
}
