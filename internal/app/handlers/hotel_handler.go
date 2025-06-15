package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"homestay.com/nguyenduy/internal/app/service"
	utils "homestay.com/nguyenduy/internal/pkg"
	"homestay.com/nguyenduy/internal/request"
)

type HotelHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type hotelHandler struct {
	hotelService service.HotelService
}

func NewHotelHandler(hotelService service.HotelService) HotelHandler {
	return &hotelHandler{hotelService: hotelService}
}

func (h *hotelHandler) Create(ctx *gin.Context) {
	var hotel request.HotelRequest
	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	if err := h.hotelService.CreateHotel(&hotel); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Hotel created successfully"})
}

func (h *hotelHandler) GetAll(ctx *gin.Context) {
	hotels, err := h.hotelService.GetAllHotels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get hotels"})
		return
	}

	ctx.JSON(http.StatusOK, hotels)
}

func (h *hotelHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	hotel, err := h.hotelService.GetHotelById(uint(idUint))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	ctx.JSON(http.StatusOK, hotel)
}

func (h *hotelHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}
	var hotel request.HotelRequest
	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	if err := h.hotelService.UpdateHotel(uint(idUint), &hotel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update hotel"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Hotel updated successfully"})
}

func (h *hotelHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}
	if err := h.hotelService.DeleteHotel(uint(idUint)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete hotel"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Hotel deleted successfully"})
}
