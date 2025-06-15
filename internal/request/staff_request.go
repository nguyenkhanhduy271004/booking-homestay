package request

import "time"

type StaffRequest struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Position    string    `json:"position"`
	Salary      float64   `json:"salary"`
	DateOfBirth time.Time `json:"date"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
}
