package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	service "homestay.com/nguyenduy/internal/app/services"
	"homestay.com/nguyenduy/internal/request"
)

type RoomTypeHandler struct {
	roomTypeService service.RoomTypeService
}

func NewRoomTypeHandler(roomTypeService service.RoomTypeService) *RoomTypeHandler {
	return &RoomTypeHandler{
		roomTypeService: roomTypeService,
	}
}

func (h *RoomTypeHandler) CreateRoomType(c *gin.Context) {
	var roomType request.RoomTypeRequest
	if err := c.ShouldBindJSON(&roomType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.roomTypeService.CreateRoomType(&roomType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Room type created successfully"})
}

func (h *RoomTypeHandler) GetAllRoomTypes(c *gin.Context) {
	roomTypes, err := h.roomTypeService.GetAllRoomTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roomTypes)
}

func (h *RoomTypeHandler) GetRoomTypeByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room type ID"})
		return
	}

	roomType, err := h.roomTypeService.GetRoomTypeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room type not found"})
		return
	}

	c.JSON(http.StatusOK, roomType)
}

func (h *RoomTypeHandler) UpdateRoomType(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room type ID"})
		return
	}

	var roomType request.RoomTypeRequest
	if err := c.ShouldBindJSON(&roomType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.roomTypeService.UpdateRoomType(uint(id), &roomType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room type updated successfully"})
}

func (h *RoomTypeHandler) DeleteRoomType(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room type ID"})
		return
	}

	if err := h.roomTypeService.DeleteRoomType(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room type deleted successfully"})
}
