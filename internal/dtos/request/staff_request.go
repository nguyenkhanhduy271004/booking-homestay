package request

import "time"

type StaffRequest struct {
	UserID      uint      `json:"user_id" binding:"required"`
	HotelID     uint      `json:"hotel_id" binding:"required"`
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	Position    string    `json:"position" binding:"required"`
	Salary      float64   `json:"salary"`
	DateOfBirth time.Time `json:"date"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	HireDate    time.Time `json:"hire_date"`
	RoleID      uint      `json:"role_id"`
}
