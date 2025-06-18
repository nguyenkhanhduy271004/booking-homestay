package converter

import (
	"homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/dtos/response"
)

func ToBookingDTO(booking *models.Booking) response.BookingResponse {
	return response.BookingResponse{
		ID:           booking.ID,
		GuestID:      booking.GuestID,
		RoomID:       booking.RoomID,
		CheckinDate:  booking.CheckinDate,
		CheckoutDate: booking.CheckoutDate,
		Status:       booking.Status,
	}
}
