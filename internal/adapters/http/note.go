package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/domain"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"github.com/rmarmolejo90/hvac/internal/config/log"
)

type NoteHandler struct {
	noteService *services.NoteService
}

func NewNoteHandler(noteService *services.NoteService) *NoteHandler {
	return &NoteHandler{noteService: noteService}
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
	var note domain.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		log.Error("Failed to bind note data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.noteService.CreateNote(c.Request.Context(), &note); err != nil {
		log.Error("Failed to create note:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	log.Info("Note created successfully:", note.ID)
	c.JSON(http.StatusCreated, note)
}

func (h *NoteHandler) GetNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid note ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	note, err := h.noteService.GetNoteByID(c.Request.Context(), uint(id))
	if err != nil {
		log.Error("Failed to get note:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get note"})
		return
	}

	if note == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	log.Info("Note retrieved successfully:", note.ID)
	c.JSON(http.StatusOK, note)
}

func (h *NoteHandler) UpdateNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid note ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var note domain.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		log.Error("Failed to bind note data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	note.ID = uint(id)
	if err := h.noteService.UpdateNote(c.Request.Context(), &note); err != nil {
		log.Error("Failed to update note:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	log.Info("Note updated successfully:", note.ID)
	c.JSON(http.StatusOK, note)
}

func (h *NoteHandler) DeleteNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error("Invalid note ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	if err := h.noteService.DeleteNote(c.Request.Context(), uint(id)); err != nil {
		log.Error("Failed to delete note:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	log.Info("Note deleted successfully:", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *NoteHandler) ListNotesByCustomerID(c *gin.Context) {
	customerID, err := strconv.ParseUint(c.Param("customer_id"), 10, 64)
	if err != nil {
		log.Error("Invalid customer ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	notes, err := h.noteService.ListNotesByCustomerID(c.Request.Context(), uint(customerID))
	if err != nil {
		log.Error("Failed to list notes:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list notes"})
		return
	}

	log.Info("Notes listed successfully for customer:", customerID)
	c.JSON(http.StatusOK, notes)
}

func (h *NoteHandler) ListNotesByLocationID(c *gin.Context) {
	locationID, err := strconv.ParseUint(c.Param("location_id"), 10, 64)
	if err != nil {
		log.Error("Invalid location ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
		return
	}

	notes, err := h.noteService.ListNotesByLocationID(c.Request.Context(), uint(locationID))
	if err != nil {
		log.Error("Failed to list notes:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list notes"})
		return
	}

	log.Info("Notes listed successfully for location:", locationID)
	c.JSON(http.StatusOK, notes)
}

func (h *NoteHandler) ListNotesByJobID(c *gin.Context) {
	jobID, err := strconv.ParseUint(c.Param("job_id"), 10, 64)
	if err != nil {
		log.Error("Invalid job ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	notes, err := h.noteService.ListNotesByJobID(c.Request.Context(), uint(jobID))
	if err != nil {
		log.Error("Failed to list notes:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list notes"})
		return
	}

	log.Info("Notes listed successfully for job:", jobID)
	c.JSON(http.StatusOK, notes)
}
