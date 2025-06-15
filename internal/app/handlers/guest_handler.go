package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	service "homestay.com/nguyenduy/internal/app/services"
	"homestay.com/nguyenduy/internal/request"
)

type GuestHandler struct {
	guestService service.GuestService
}

func NewGuestHandler(guestService service.GuestService) *GuestHandler {
	return &GuestHandler{
		guestService: guestService,
	}
}

func (h *GuestHandler) CreateGuest(c *gin.Context) {
	var guest request.GuestRequest
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.guestService.CreateGuest(&guest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Guest created successfully"})
}

func (h *GuestHandler) GetAllGuests(c *gin.Context) {
	guests, err := h.guestService.GetAllGuests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guests)
}

func (h *GuestHandler) GetGuestByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid guest ID"})
		return
	}

	guest, err := h.guestService.GetGuestByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guest not found"})
		return
	}

	c.JSON(http.StatusOK, guest)
}

func (h *GuestHandler) UpdateGuest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid guest ID"})
		return
	}

	var guest request.GuestRequest
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.guestService.UpdateGuest(uint(id), &guest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Guest updated successfully"})
}

func (h *GuestHandler) DeleteGuest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid guest ID"})
		return
	}

	if err := h.guestService.DeleteGuest(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Guest deleted successfully"})
}
