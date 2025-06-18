package converter

import (
	"homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/dtos/response"
)

func ToRoomTypeDTO(RoomType *models.RoomType) response.RoomTypeResponse {
	return response.RoomTypeResponse{
		ID:        RoomType.ID,
		Name:      RoomType.Name,
		CreatedAt: RoomType.CreatedAt,
		UpdatedAt: RoomType.UpdatedAt,
	}
}
