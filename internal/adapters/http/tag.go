package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type TagHandler struct {
	tagService *services.TagService
}

func NewTagHandler(tagService *services.TagService) *TagHandler {
	return &TagHandler{tagService: tagService}
}

func (h *TagHandler) CreateTag(c *gin.Context) {
	var tag domain.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		log.Error("Failed to bind tag data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.tagService.CreateTag(c.Request.Context(), &tag); err != nil {
		log.Error("Failed to create tag:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	log.Info("Tag created successfully:", tag.ID)
	c.JSON(http.StatusCreated, tag)
}

func (h *TagHandler) GetTag(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid tag ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	tag, err := h.tagService.GetTagByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get tag:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tag"})
		return
	}

	if tag == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	log.Info("Tag retrieved successfully:", tag.ID)
	c.JSON(http.StatusOK, tag)
}

func (h *TagHandler) UpdateTag(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid tag ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	var tag domain.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		log.Error("Failed to bind tag data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	tag.ID = uint(id)
	if err := h.tagService.UpdateTag(c.Request.Context(), &tag); err != nil {
		log.Error("Failed to update tag:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tag"})
		return
	}

	log.Info("Tag updated successfully:", tag.ID)
	c.JSON(http.StatusOK, tag)
}

func (h *TagHandler) DeleteTag(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid tag ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	if err := h.tagService.DeleteTag(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete tag:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
		return
	}

	log.Info("Tag deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *TagHandler) ListTags(c *gin.Context) {
	tags, err := h.tagService.ListTags(c.Request.Context())
	if err != nil {
		log.Error("Failed to list tags:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list tags"})
		return
	}

	log.Info("Tags listed successfully")
	c.JSON(http.StatusOK, tags)
}
