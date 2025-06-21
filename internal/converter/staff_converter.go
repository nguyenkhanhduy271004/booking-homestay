package converter

import (
	"homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/dtos/response"
)

func ToStaffDTO(staff *models.Staff) response.StaffResponse {
	return response.StaffResponse{
		ID:          staff.ID,
		FirstName:   staff.FirstName,
		LastName:    staff.LastName,
		Position:    staff.Position,
		Salary:      staff.Salary,
		Email:       staff.Email,
		Phone:       staff.Phone,
		DateOfBirth: staff.DateOfBirth,
		HireDate:    staff.HireDate,
	}
}
