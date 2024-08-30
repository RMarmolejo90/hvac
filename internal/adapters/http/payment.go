package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type PaymentHandler struct {
	paymentService *services.PaymentService
}

func NewPaymentHandler(paymentService *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var payment domain.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		log.Error("Failed to bind payment data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.paymentService.CreatePayment(c.Request.Context(), &payment); err != nil {
		log.Error("Failed to create payment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	log.Info("Payment created successfully:", payment.ID)
	c.JSON(http.StatusCreated, payment)
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid payment ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	payment, err := h.paymentService.GetPaymentByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get payment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get payment"})
		return
	}

	if payment == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	log.Info("Payment retrieved successfully:", payment.ID)
	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) UpdatePayment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid payment ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	var payment domain.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		log.Error("Failed to bind payment data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	payment.ID = uint(id)
	if err := h.paymentService.UpdatePayment(c.Request.Context(), &payment); err != nil {
		log.Error("Failed to update payment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment"})
		return
	}

	log.Info("Payment updated successfully:", payment.ID)
	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) DeletePayment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid payment ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	if err := h.paymentService.DeletePayment(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete payment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete payment"})
		return
	}

	log.Info("Payment deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *PaymentHandler) ListPaymentsByInvoiceID(c *gin.Context) {
	invoiceID, err := strconv.ParseUint(c.Param("invoice_id"), 10, 64)
	if err != nil {
		log.Error("Invalid invoice ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	payments, err := h.paymentService.ListPaymentsByInvoiceID(c.Request.Context(), uint(invoiceID))
	if err != nil {
		log.Error("Failed to list payments:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list payments"})
		return
	}

	log.Info("Payments listed successfully for invoice:", invoiceID)
	c.JSON(http.StatusOK, payments)
}
