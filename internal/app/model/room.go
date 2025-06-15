package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	HotelID uint     `json:"hotel_id"`
	Hotel   Hotel    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TypeID  uint     `json:"type_id"`
	Type    RoomType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status  bool     `json:"status"`
}
