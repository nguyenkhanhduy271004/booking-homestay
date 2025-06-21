package services

import (
	"errors"
	"strings"
	"time"

	"homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/converter"
	"homestay.com/nguyenduy/internal/dtos/request"
	"homestay.com/nguyenduy/internal/dtos/response"
)

type StaffService interface {
	CreateStaff(staff *request.StaffRequest) error
	GetAllStaff() ([]response.StaffResponse, error)
	GetStaffByID(id uint) (*response.StaffResponse, error)
	UpdateStaff(id uint, staff *request.StaffRequest) error
	DeleteStaff(id uint) error
}

type staffService struct {
	staffRepo repositories.StaffRepository
}

func NewStaffService(staffRepo repositories.StaffRepository) StaffService {
	return &staffService{
		staffRepo: staffRepo,
	}
}

func (s *staffService) validateStaff(staff *request.StaffRequest) error {
	if staff.FirstName == "" {
		return errors.New("first name is required")
	}
	if staff.LastName == "" {
		return errors.New("last name is required")
	}
	if staff.Position == "" {
		return errors.New("position is required")
	}
	if staff.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	if staff.Email != "" {
		if !strings.Contains(staff.Email, "@") {
			return errors.New("invalid email format")
		}
	}
	return nil
}

func (s *staffService) CreateStaff(staff *request.StaffRequest) error {
	if err := s.validateStaff(staff); err != nil {
		return err
	}

	if staff.HireDate.IsZero() {
		staff.HireDate = time.Now()
	}

	if staff.DateOfBirth.IsZero() {
		staff.DateOfBirth = time.Now().AddDate(-18, 0, 0)
	}

	return s.staffRepo.Create(staff)
}

func (s *staffService) GetAllStaff() ([]response.StaffResponse, error) {
	staffs, err := s.staffRepo.GetAll()

	if err != nil {
		return nil, err
	}

	staffDtos := make([]response.StaffResponse, len(staffs))

	for i, staff := range staffs {
		staffDtos[i] = converter.ToStaffDTO(&staff)
	}

	return staffDtos, nil
}

func (s *staffService) GetStaffByID(id uint) (*response.StaffResponse, error) {
	if id == 0 {
		return nil, errors.New("invalid staff ID")
	}

	staff, err := s.staffRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	staffDto := converter.ToStaffDTO(staff)

	return &staffDto, nil
}

func (s *staffService) UpdateStaff(id uint, staff *request.StaffRequest) error {
	if id == 0 {
		return errors.New("invalid staff ID")
	}

	if err := s.validateStaff(staff); err != nil {
		return err
	}

	if staff.HireDate.IsZero() {
		staff.HireDate = time.Now()
	}

	if staff.DateOfBirth.IsZero() {
		staff.DateOfBirth = time.Now().AddDate(-18, 0, 0)
	}

	return s.staffRepo.Update(id, staff)
}

func (s *staffService) DeleteStaff(id uint) error {
	if id == 0 {
		return errors.New("invalid staff ID")
	}
	return s.staffRepo.Delete(id)
}
