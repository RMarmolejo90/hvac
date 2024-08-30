package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type InvoiceHandler struct {
	invoiceService *services.InvoiceService
}

func NewInvoiceHandler(invoiceService *services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{invoiceService: invoiceService}
}

func (h *InvoiceHandler) CreateInvoice(c *gin.Context) {
	var invoice domain.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		log.Error("Failed to bind invoice data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.invoiceService.CreateInvoice(c.Request.Context(), &invoice); err != nil {
		log.Error("Failed to create invoice:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
		return
	}

	log.Info("Invoice created successfully:", invoice.ID)
	c.JSON(http.StatusCreated, invoice)
}

func (h *InvoiceHandler) GetInvoice(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid invoice ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	invoice, err := h.invoiceService.GetInvoiceByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get invoice:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get invoice"})
		return
	}

	if invoice == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	log.Info("Invoice retrieved successfully:", invoice.ID)
	c.JSON(http.StatusOK, invoice)
}

func (h *InvoiceHandler) UpdateInvoice(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid invoice ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	var invoice domain.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		log.Error("Failed to bind invoice data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	invoice.ID = uint(id)
	if err := h.invoiceService.UpdateInvoice(c.Request.Context(), &invoice); err != nil {
		log.Error("Failed to update invoice:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update invoice"})
		return
	}

	log.Info("Invoice updated successfully:", invoice.ID)
	c.JSON(http.StatusOK, invoice)
}

func (h *InvoiceHandler) DeleteInvoice(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid invoice ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	if err := h.invoiceService.DeleteInvoice(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete invoice:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete invoice"})
		return
	}

	log.Info("Invoice deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *InvoiceHandler) ListInvoices(c *gin.Context) {
	invoices, err := h.invoiceService.ListInvoices(c.Request.Context())
	if err != nil {
		log.Error("Failed to list invoices:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list invoices"})
		return
	}

	log.Info("Invoices listed successfully")
	c.JSON(http.StatusOK, invoices)
}
