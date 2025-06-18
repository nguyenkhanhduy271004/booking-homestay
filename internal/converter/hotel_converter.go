package converter

import (
	"homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/dtos/response"
)

func ToHotelDTO(user models.Hotel) response.HotelResponse {
	return response.HotelResponse{
		ID:      user.ID,
		Name:    user.Name,
		Address: user.Address,
		Phone:   user.Phone,
		Email:   user.Email,
		Stars:   user.Stars,
		// Image:        config.API_URL + user.Image,
		Image:        user.Image,
		CheckinTime:  user.CheckinTime,
		CheckoutTime: user.CheckoutTime,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
