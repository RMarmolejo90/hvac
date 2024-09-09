package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
)

type ContactMethodHandler struct {
	contactMethodService *services.ContactMethodService
}

func NewContactMethodHandler(contactMethodService *services.ContactMethodService) *ContactMethodHandler {
	return &ContactMethodHandler{contactMethodService: contactMethodService}
}

func (h *ContactMethodHandler) CreateContactMethod(c *gin.Context) {
	var contactMethod domain.ContactMethod
	if err := c.ShouldBindJSON(&contactMethod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.contactMethodService.CreateContactMethod(c.Request.Context(), &contactMethod); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contact method"})
		return
	}

	c.JSON(http.StatusCreated, contactMethod)
}

func (h *ContactMethodHandler) GetContactMethod(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact method ID"})
		return
	}

	contactMethod, err := h.contactMethodService.FindContactMethodByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get contact method"})
		return
	}

	if contactMethod == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact method not found"})
		return
	}

	c.JSON(http.StatusOK, contactMethod)
}

func (h *ContactMethodHandler) UpdateContactMethod(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact method ID"})
		return
	}

	var contactMethod domain.ContactMethod
	if err := c.ShouldBindJSON(&contactMethod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	contactMethod.ID = uint(id)
	if err := h.contactMethodService.UpdateContactMethod(c.Request.Context(), &contactMethod); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update contact method"})
		return
	}

	c.JSON(http.StatusOK, contactMethod)
}

func (h *ContactMethodHandler) DeleteContactMethod(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact method ID"})
		return
	}

	if err := h.contactMethodService.DeleteContactMethod(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contact method"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *ContactMethodHandler) ListContactMethodsByCustomerID(c *gin.Context) {
	customerID, err := strconv.ParseUint(c.Param("customer_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	contactMethods, err := h.contactMethodService.ListContactMethodsByCustomerID(c.Request.Context(), uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list contact methods"})
		return
	}

	c.JSON(http.StatusOK, contactMethods)
}
