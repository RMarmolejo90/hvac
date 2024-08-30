package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type CustomerHandler struct {
	customerService *services.CustomerService
}

func NewCustomerHandler(customerService *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer domain.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		log.Error("Failed to bind customer data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.customerService.CreateCustomer(c.Request.Context(), &customer); err != nil {
		log.Error("Failed to create customer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	log.Info("Customer created successfully:", customer.ID)
	c.JSON(http.StatusCreated, customer)
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid customer ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	customer, err := h.customerService.GetCustomerByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get customer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get customer"})
		return
	}

	if customer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	log.Info("Customer retrieved successfully:", customer.ID)
	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid customer ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	var customer domain.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		log.Error("Failed to bind customer data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	customer.ID = uint(id)
	if err := h.customerService.UpdateCustomer(c.Request.Context(), &customer); err != nil {
		log.Error("Failed to update customer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	log.Info("Customer updated successfully:", customer.ID)
	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid customer ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	if err := h.customerService.DeleteCustomer(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete customer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	log.Info("Customer deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	customers, err := h.customerService.ListCustomers(c.Request.Context())
	if err != nil {
		log.Error("Failed to list customers:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list customers"})
		return
	}

	log.Info("Customers listed successfully")
	c.JSON(http.StatusOK, customers)
}
