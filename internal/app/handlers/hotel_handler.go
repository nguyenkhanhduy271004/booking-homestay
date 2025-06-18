package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"homestay.com/nguyenduy/internal/app/services"
	"homestay.com/nguyenduy/internal/dtos/request"
	"homestay.com/nguyenduy/internal/helper"
)

type HotelHandler struct {
	hotelService services.HotelService
}

func NewHotelHandler(hotelService services.HotelService) *HotelHandler {
	return &HotelHandler{
		hotelService: hotelService,
	}
}

func (h *HotelHandler) CreateHotel(c *gin.Context) {
	var hotelRequest request.HotelRequest
	if err := c.ShouldBindJSON(&hotelRequest); err != nil {
		response := helper.BuildErrorResponse("Invalid request data", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := h.hotelService.CreateHotel(&hotelRequest); err != nil {
		response := helper.BuildErrorResponse("Failed to create hotel", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.BuildResponse(true, "Hotel created successfully", nil)
	c.JSON(http.StatusCreated, response)
}

func (h *HotelHandler) GetAllHotels(c *gin.Context) {
	hotels, err := h.hotelService.GetAllHotels()
	if err != nil {
		response := helper.BuildErrorResponse("Failed to fetch hotels", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.BuildResponse(true, "Hotels fetched successfully", hotels)
	c.JSON(http.StatusOK, response)
}

func (h *HotelHandler) GetHotelByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response := helper.BuildErrorResponse("Invalid hotel ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	hotel, err := h.hotelService.GetHotelByID(uint(id))
	if err != nil {
		response := helper.BuildErrorResponse("Hotel not found", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.BuildResponse(true, "Hotel fetched successfully", hotel)
	c.JSON(http.StatusOK, response)
}

func (h *HotelHandler) UpdateHotel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response := helper.BuildErrorResponse("Invalid hotel ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var hotelRequest request.HotelRequest
	if err := c.ShouldBindJSON(&hotelRequest); err != nil {
		response := helper.BuildErrorResponse("Invalid request data", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := h.hotelService.UpdateHotel(uint(id), &hotelRequest); err != nil {
		response := helper.BuildErrorResponse("Failed to update hotel", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.BuildResponse(true, "Hotel updated successfully", nil)
	c.JSON(http.StatusOK, response)
}

func (h *HotelHandler) DeleteHotel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response := helper.BuildErrorResponse("Invalid hotel ID", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := h.hotelService.DeleteHotel(uint(id)); err != nil {
		response := helper.BuildErrorResponse("Failed to delete hotel", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.BuildResponse(true, "Hotel deleted successfully", nil)
	c.JSON(http.StatusOK, response)
}
