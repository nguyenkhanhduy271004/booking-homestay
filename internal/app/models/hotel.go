package models

import (
	"time"

	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Stars        int       `json:"stars"`
	Image        string    `json:"image"`
	CheckinTime  time.Time `json:"checkin_time"`
	CheckoutTime time.Time `json:"checkout_time"`
	Staffs       []Staff   `gorm:"foreignKey:HotelID"`
}
