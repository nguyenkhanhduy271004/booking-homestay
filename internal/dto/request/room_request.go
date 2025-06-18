package request

type RoomRequest struct {
	HotelID uint `json:"hotel_id" binding:"required"`
	TypeID  uint `json:"type_id" binding:"required"`
	Status  bool `json:"status"`
}
