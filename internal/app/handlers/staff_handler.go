package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"homestay.com/nguyenduy/internal/app/services"
	"homestay.com/nguyenduy/internal/request"
)

type StaffHandler struct {
	staffService services.StaffService
}

func NewStaffHandler(staffService services.StaffService) *StaffHandler {
	return &StaffHandler{
		staffService: staffService,
	}
}

func (h *StaffHandler) CreateStaff(c *gin.Context) {
	var staffRequest request.StaffRequest
	if err := c.ShouldBindJSON(&staffRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	if err := h.staffService.CreateStaff(&staffRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create staff",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Staff created successfully",
	})
}

func (h *StaffHandler) GetAllStaff(c *gin.Context) {
	staff, err := h.staffService.GetAllStaff()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch staff",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   staff,
	})
}

func (h *StaffHandler) GetStaffByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid staff ID",
			"error":   err.Error(),
		})
		return
	}

	staff, err := h.staffService.GetStaffByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Staff not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   staff,
	})
}

func (h *StaffHandler) UpdateStaff(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid staff ID",
			"error":   err.Error(),
		})
		return
	}

	var staffRequest request.StaffRequest
	if err := c.ShouldBindJSON(&staffRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	if err := h.staffService.UpdateStaff(uint(id), &staffRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to update staff",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Staff updated successfully",
	})
}

func (h *StaffHandler) DeleteStaff(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid staff ID",
			"error":   err.Error(),
		})
		return
	}

	if err := h.staffService.DeleteStaff(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to delete staff",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Staff deleted successfully",
	})
}
