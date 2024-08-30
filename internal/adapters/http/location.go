package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type LocationHandler struct {
	locationService *services.LocationService
}

func NewLocationHandler(locationService *services.LocationService) *LocationHandler {
	return &LocationHandler{locationService: locationService}
}

func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var location domain.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		log.Error("Failed to bind location data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.locationService.CreateLocation(c.Request.Context(), &location); err != nil {
		log.Error("Failed to create location:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create location"})
		return
	}

	log.Info("Location created successfully:", location.ID)
	c.JSON(http.StatusCreated, location)
}

func (h *LocationHandler) GetLocation(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid location ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
		return
	}

	location, err := h.locationService.GetLocationByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get location:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get location"})
		return
	}

	if location == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	log.Info("Location retrieved successfully:", location.ID)
	c.JSON(http.StatusOK, location)
}

func (h *LocationHandler) UpdateLocation(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid location ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
		return
	}

	var location domain.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		log.Error("Failed to bind location data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	location.ID = uint(id)
	if err := h.locationService.UpdateLocation(c.Request.Context(), &location); err != nil {
		log.Error("Failed to update location:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update location"})
		return
	}

	log.Info("Location updated successfully:", location.ID)
	c.JSON(http.StatusOK, location)
}

func (h *LocationHandler) DeleteLocation(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid location ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
		return
	}

	if err := h.locationService.DeleteLocation(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete location:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete location"})
		return
	}

	log.Info("Location deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *LocationHandler) ListLocationsByCustomerID(c *gin.Context) {
	customerID, err := strconv.ParseUint(c.Param("customer_id"), 10, 64)
	if err != nil {
		log.Error("Invalid customer ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	locations, err := h.locationService.ListLocationsByCustomerID(c.Request.Context(), uint(customerID))
	if err != nil {
		log.Error("Failed to list locations:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list locations"})
		return
	}

	log.Info("Locations listed successfully for customer:", customerID)
	c.JSON(http.StatusOK, locations)
}
