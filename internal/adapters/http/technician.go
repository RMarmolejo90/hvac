package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type TechnicianHandler struct {
	technicianService *services.TechnicianService
}

func NewTechnicianHandler(technicianService *services.TechnicianService) *TechnicianHandler {
	return &TechnicianHandler{technicianService: technicianService}
}

func (h *TechnicianHandler) CreateTechnician(c *gin.Context) {
	var technician domain.Technician
	if err := c.ShouldBindJSON(&technician); err != nil {
		log.Error("Failed to bind technician data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.technicianService.CreateTechnician(c.Request.Context(), &technician); err != nil {
		log.Error("Failed to create technician:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create technician"})
		return
	}

	log.Info("Technician created successfully:", technician.ID)
	c.JSON(http.StatusCreated, technician)
}

func (h *TechnicianHandler) GetTechnician(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid technician ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid technician ID"})
		return
	}

	technician, err := h.technicianService.GetTechnicianByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get technician:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get technician"})
		return
	}

	if technician == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Technician not found"})
		return
	}

	log.Info("Technician retrieved successfully:", technician.ID)
	c.JSON(http.StatusOK, technician)
}

func (h *TechnicianHandler) UpdateTechnician(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid technician ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid technician ID"})
		return
	}

	var technician domain.Technician
	if err := c.ShouldBindJSON(&technician); err != nil {
		log.Error("Failed to bind technician data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	technician.ID = uint(id)
	if err := h.technicianService.UpdateTechnician(c.Request.Context(), &technician); err != nil {
		log.Error("Failed to update technician:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update technician"})
		return
	}

	log.Info("Technician updated successfully:", technician.ID)
	c.JSON(http.StatusOK, technician)
}

func (h *TechnicianHandler) DeleteTechnician(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid technician ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid technician ID"})
		return
	}

	if err := h.technicianService.DeleteTechnician(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete technician:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete technician"})
		return
	}

	log.Info("Technician deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *TechnicianHandler) ListTechnicians(c *gin.Context) {
	technicians, err := h.technicianService.ListTechnicians(c.Request.Context())
	if err != nil {
		log.Error("Failed to list technicians:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list technicians"})
		return
	}

	log.Info("Technicians listed successfully")
	c.JSON(http.StatusOK, technicians)
}
