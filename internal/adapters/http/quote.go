package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type QuoteHandler struct {
	quoteService *services.QuoteService
}

func NewQuoteHandler(quoteService *services.QuoteService) *QuoteHandler {
	return &QuoteHandler{quoteService: quoteService}
}

func (h *QuoteHandler) CreateQuote(c *gin.Context) {
	var quote domain.Quote
	if err := c.ShouldBindJSON(&quote); err != nil {
		log.Error("Failed to bind quote data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.quoteService.CreateQuote(c.Request.Context(), &quote); err != nil {
		log.Error("Failed to create quote:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quote"})
		return
	}

	log.Info("Quote created successfully:", quote.ID)
	c.JSON(http.StatusCreated, quote)
}

func (h *QuoteHandler) GetQuote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid quote ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quote ID"})
		return
	}

	quote, err := h.quoteService.GetQuoteByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get quote:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get quote"})
		return
	}

	if quote == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quote not found"})
		return
	}

	log.Info("Quote retrieved successfully:", quote.ID)
	c.JSON(http.StatusOK, quote)
}

func (h *QuoteHandler) UpdateQuote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid quote ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quote ID"})
		return
	}

	var quote domain.Quote
	if err := c.ShouldBindJSON(&quote); err != nil {
		log.Error("Failed to bind quote data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	quote.ID = uint(id)
	if err := h.quoteService.UpdateQuote(c.Request.Context(), &quote); err != nil {
		log.Error("Failed to update quote:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update quote"})
		return
	}

	log.Info("Quote updated successfully:", quote.ID)
	c.JSON(http.StatusOK, quote)
}

func (h *QuoteHandler) DeleteQuote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid quote ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quote ID"})
		return
	}

	if err := h.quoteService.DeleteQuote(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete quote:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete quote"})
		return
	}

	log.Info("Quote deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *QuoteHandler) ListQuotes(c *gin.Context) {
	quotes, err := h.quoteService.ListQuotes(c.Request.Context())
	if err != nil {
		log.Error("Failed to list quotes:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list quotes"})
		return
	}

	log.Info("Quotes listed successfully")
	c.JSON(http.StatusOK, quotes)
}
