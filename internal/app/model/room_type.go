package model

import "gorm.io/gorm"

type RoomType struct {
	gorm.Model
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	PricePerNight float64 `json:"price_per_night"`
	Capacity      int     `json:"capacity"`
}
