package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type StockHandler struct {
	stockService *services.StockService
}

func NewStockHandler(stockService *services.StockService) *StockHandler {
	return &StockHandler{stockService: stockService}
}

func (h *StockHandler) CreateStock(c *gin.Context) {
	var stock domain.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		log.Error("Failed to bind stock data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.stockService.CreateStock(c.Request.Context(), &stock); err != nil {
		log.Error("Failed to create stock:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock"})
		return
	}

	log.Info("Stock created successfully:", stock.ID)
	c.JSON(http.StatusCreated, stock)
}

func (h *StockHandler) GetStock(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid stock ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	stock, err := h.stockService.GetStockByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get stock:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stock"})
		return
	}

	if stock == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	log.Info("Stock retrieved successfully:", stock.ID)
	c.JSON(http.StatusOK, stock)
}

func (h *StockHandler) UpdateStock(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid stock ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	var stock domain.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		log.Error("Failed to bind stock data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	stock.ID = uint(id)
	if err := h.stockService.UpdateStock(c.Request.Context(), &stock); err != nil {
		log.Error("Failed to update stock:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
		return
	}

	log.Info("Stock updated successfully:", stock.ID)
	c.JSON(http.StatusOK, stock)
}

func (h *StockHandler) DeleteStock(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid stock ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	if err := h.stockService.DeleteStock(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete stock:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete stock"})
		return
	}

	log.Info("Stock deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *StockHandler) ListStockByTechnicianID(c *gin.Context) {
	technicianID, err := strconv.ParseUint(c.Param("technician_id"), 10, 64)
	if err != nil {
		log.Error("Invalid technician ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid technician ID"})
		return
	}

	stocks, err := h.stockService.ListStockByTechnicianID(c.Request.Context(), uint(technicianID))
	if err != nil {
		log.Error("Failed to list stock:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list stock"})
		return
	}

	log.Info("Stock listed successfully for technician:", technicianID)
	c.JSON(http.StatusOK, stocks)
}
