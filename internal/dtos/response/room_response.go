package response

import "time"

type RoomResponse struct {
	ID        uint      `json:"id"`
	HotelID   uint      `json:"hotel_id"`
	TypeID    uint      `json:"type_id"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
