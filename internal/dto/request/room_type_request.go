package request

type RoomTypeRequest struct {
	Name          string  `json:"name" binding:"required"`
	Description   string  `json:"description"`
	PricePerNight float64 `json:"price_per_night" binding:"required"`
	Capacity      int     `json:"capacity" binding:"required"`
}
