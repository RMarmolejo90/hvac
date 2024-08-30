package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type HourlyRateHandler struct {
	hourlyRateService *services.HourlyRateService
}

func NewHourlyRateHandler(hourlyRateService *services.HourlyRateService) *HourlyRateHandler {
	return &HourlyRateHandler{hourlyRateService: hourlyRateService}
}

func (h *HourlyRateHandler) CreateHourlyRate(c *gin.Context) {
	var hourlyRate domain.HourlyRate
	if err := c.ShouldBindJSON(&hourlyRate); err != nil {
		log.Error("Failed to bind hourly rate data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.hourlyRateService.CreateHourlyRate(c.Request.Context(), &hourlyRate); err != nil {
		log.Error("Failed to create hourly rate:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hourly rate"})
		return
	}

	log.Info("Hourly rate created successfully:", hourlyRate.ID)
	c.JSON(http.StatusCreated, hourlyRate)
}

func (h *HourlyRateHandler) GetHourlyRate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid hourly rate ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hourly rate ID"})
		return
	}

	hourlyRate, err := h.hourlyRateService.GetHourlyRateByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get hourly rate:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get hourly rate"})
		return
	}

	if hourlyRate == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hourly rate not found"})
		return
	}

	log.Info("Hourly rate retrieved successfully:", hourlyRate.ID)
	c.JSON(http.StatusOK, hourlyRate)
}

func (h *HourlyRateHandler) UpdateHourlyRate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid hourly rate ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hourly rate ID"})
		return
	}

	var hourlyRate domain.HourlyRate
	if err := c.ShouldBindJSON(&hourlyRate); err != nil {
		log.Error("Failed to bind hourly rate data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hourlyRate.ID = uint(id)
	if err := h.hourlyRateService.UpdateHourlyRate(c.Request.Context(), &hourlyRate); err != nil {
		log.Error("Failed to update hourly rate:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update hourly rate"})
		return
	}

	log.Info("Hourly rate updated successfully:", hourlyRate.ID)
	c.JSON(http.StatusOK, hourlyRate)
}

func (h *HourlyRateHandler) DeleteHourlyRate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid hourly rate ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hourly rate ID"})
		return
	}

	if err := h.hourlyRateService.DeleteHourlyRate(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete hourly rate:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete hourly rate"})
		return
	}

	log.Info("Hourly rate deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *HourlyRateHandler) ListHourlyRates(c *gin.Context) {
	hourlyRates, err := h.hourlyRateService.ListHourlyRates(c.Request.Context())
	if err != nil {
		log.Error("Failed to list hourly rates:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list hourly rates"})
		return
	}

	log.Info("Hourly rates listed successfully")
	c.JSON(http.StatusOK, hourlyRates)
}
