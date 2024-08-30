package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type EquipmentHandler struct {
	equipmentService *services.EquipmentService
}

func NewEquipmentHandler(equipmentService *services.EquipmentService) *EquipmentHandler {
	return &EquipmentHandler{equipmentService: equipmentService}
}

func (h *EquipmentHandler) CreateEquipment(c *gin.Context) {
	var equipment domain.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		log.Error("Failed to bind equipment data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.equipmentService.CreateEquipment(c.Request.Context(), &equipment); err != nil {
		log.Error("Failed to create equipment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create equipment"})
		return
	}

	log.Info("Equipment created successfully:", equipment.ID)
	c.JSON(http.StatusCreated, equipment)
}

func (h *EquipmentHandler) GetEquipment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid equipment ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	equipment, err := h.equipmentService.GetEquipmentByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get equipment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get equipment"})
		return
	}

	if equipment == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Equipment not found"})
		return
	}

	log.Info("Equipment retrieved successfully:", equipment.ID)
	c.JSON(http.StatusOK, equipment)
}

func (h *EquipmentHandler) UpdateEquipment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid equipment ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	var equipment domain.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		log.Error("Failed to bind equipment data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	equipment.ID = uint(id)
	if err := h.equipmentService.UpdateEquipment(c.Request.Context(), &equipment); err != nil {
		log.Error("Failed to update equipment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update equipment"})
		return
	}

	log.Info("Equipment updated successfully:", equipment.ID)
	c.JSON(http.StatusOK, equipment)
}

func (h *EquipmentHandler) DeleteEquipment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid equipment ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	if err := h.equipmentService.DeleteEquipment(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete equipment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete equipment"})
		return
	}

	log.Info("Equipment deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *EquipmentHandler) ListEquipmentByLocationID(c *gin.Context) {
	locationID, err := strconv.ParseUint(c.Param("location_id"), 10, 64)
	if err != nil {
		log.Error("Invalid location ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
		return
	}

	equipment, err := h.equipmentService.ListEquipmentByLocationID(c.Request.Context(), uint(locationID))
	if err != nil {
		log.Error("Failed to list equipment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list equipment"})
		return
	}

	log.Info("Equipment listed successfully for location:", locationID)
	c.JSON(http.StatusOK, equipment)
}
