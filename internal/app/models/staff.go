package models

import (
	"time"

	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	HotelID     uint      `json:"hotel_id"`
	Hotel       Hotel     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Position    string    `json:"position"`
	Salary      float64   `json:"salary"`
	DateOfBirth time.Time `json:"date"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	HireDate    time.Time `json:"hire_date"`
	RoleID      uint      `json:"role_id"`
	Role        Role      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
}
