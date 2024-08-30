package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type ConsumablesHandler struct {
	consumablesService *services.ConsumablesService
}

func NewConsumablesHandler(consumablesService *services.ConsumablesService) *ConsumablesHandler {
	return &ConsumablesHandler{consumablesService: consumablesService}
}

func (h *ConsumablesHandler) CreateConsumable(c *gin.Context) {
	var consumable domain.Consumables
	if err := c.ShouldBindJSON(&consumable); err != nil {
		log.Error("Failed to bind consumable data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.consumablesService.CreateConsumable(c.Request.Context(), &consumable); err != nil {
		log.Error("Failed to create consumable:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create consumable"})
		return
	}

	log.Info("Consumable created successfully:", consumable.ID)
	c.JSON(http.StatusCreated, consumable)
}

func (h *ConsumablesHandler) GetConsumable(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid consumable ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid consumable ID"})
		return
	}

	consumable, err := h.consumablesService.GetConsumableByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get consumable:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get consumable"})
		return
	}

	if consumable == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Consumable not found"})
		return
	}

	log.Info("Consumable retrieved successfully:", consumable.ID)
	c.JSON(http.StatusOK, consumable)
}

func (h *ConsumablesHandler) UpdateConsumable(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid consumable ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid consumable ID"})
		return
	}

	var consumable domain.Consumables
	if err := c.ShouldBindJSON(&consumable); err != nil {
		log.Error("Failed to bind consumable data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	consumable.ID = uint(id)
	if err := h.consumablesService.UpdateConsumable(c.Request.Context(), &consumable); err != nil {
		log.Error("Failed to update consumable:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update consumable"})
		return
	}

	log.Info("Consumable updated successfully:", consumable.ID)
	c.JSON(http.StatusOK, consumable)
}

func (h *ConsumablesHandler) DeleteConsumable(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid consumable ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid consumable ID"})
		return
	}

	if err := h.consumablesService.DeleteConsumable(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete consumable:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete consumable"})
		return
	}

	log.Info("Consumable deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *ConsumablesHandler) ListConsumablesByEquipmentID(c *gin.Context) {
	equipmentID, err := strconv.ParseUint(c.Param("equipment_id"), 10, 64)
	if err != nil {
		log.Error("Invalid equipment ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	consumables, err := h.consumablesService.ListConsumablesByEquipmentID(c.Request.Context(), uint(equipmentID))
	if err != nil {
		log.Error("Failed to list consumables:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list consumables"})
		return
	}

	log.Info("Consumables listed successfully for equipment:", equipmentID)
	c.JSON(http.StatusOK, consumables)
}
