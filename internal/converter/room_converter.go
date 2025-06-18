package converter

import (
	"homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/dtos/response"
)

func ToRoomDTO(model *models.Room) response.RoomResponse {
	return response.RoomResponse{
		ID:        model.ID,
		HotelID:   model.HotelID,
		TypeID:    model.TypeID,
		Status:    model.Status,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
