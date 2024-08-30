package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type ServiceHandler struct {
	serviceService *services.ServiceService
}

func NewServiceHandler(serviceService *services.ServiceService) *ServiceHandler {
	return &ServiceHandler{serviceService: serviceService}
}

func (h *ServiceHandler) CreateService(c *gin.Context) {
	var service domain.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		log.Error("Failed to bind service data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.serviceService.CreateService(c.Request.Context(), &service); err != nil {
		log.Error("Failed to create service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service"})
		return
	}

	log.Info("Service created successfully:", service.ID)
	c.JSON(http.StatusCreated, service)
}

func (h *ServiceHandler) GetService(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid service ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	service, err := h.serviceService.GetServiceByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get service"})
		return
	}

	if service == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	log.Info("Service retrieved successfully:", service.ID)
	c.JSON(http.StatusOK, service)
}

func (h *ServiceHandler) UpdateService(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid service ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	var service domain.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		log.Error("Failed to bind service data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	service.ID = uint(id)
	if err := h.serviceService.UpdateService(c.Request.Context(), &service); err != nil {
		log.Error("Failed to update service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update service"})
		return
	}

	log.Info("Service updated successfully:", service.ID)
	c.JSON(http.StatusOK, service)
}

func (h *ServiceHandler) DeleteService(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid service ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	if err := h.serviceService.DeleteService(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete service"})
		return
	}

	log.Info("Service deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *ServiceHandler) ListServices(c *gin.Context) {
	services, err := h.serviceService.ListServices(c.Request.Context())
	if err != nil {
		log.Error("Failed to list services:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list services"})
		return
	}

	log.Info("Services listed successfully")
	c.JSON(http.StatusOK, services)
}
