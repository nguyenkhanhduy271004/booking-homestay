package services

import (
	"errors"
	"strings"
	"time"

	model "homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/request"
)

type StaffService interface {
	CreateStaff(staff *request.StaffRequest) error
	GetAllStaff() ([]model.Staff, error)
	GetStaffByID(id uint) (*model.Staff, error)
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
		// Basic email validation
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

	// Set hire date to current time if not provided
	if staff.HireDate.IsZero() {
		staff.HireDate = time.Now()
	}

	// Set a default date of birth if not provided (e.g., 18 years ago)
	if staff.DateOfBirth.IsZero() {
		staff.DateOfBirth = time.Now().AddDate(-18, 0, 0)
	}

	return s.staffRepo.Create(staff)
}

func (s *staffService) GetAllStaff() ([]model.Staff, error) {
	return s.staffRepo.GetAll()
}

func (s *staffService) GetStaffByID(id uint) (*model.Staff, error) {
	if id == 0 {
		return nil, errors.New("invalid staff ID")
	}
	return s.staffRepo.GetByID(id)
}

func (s *staffService) UpdateStaff(id uint, staff *request.StaffRequest) error {
	if id == 0 {
		return errors.New("invalid staff ID")
	}

	if err := s.validateStaff(staff); err != nil {
		return err
	}

	// Set hire date to current time if not provided
	if staff.HireDate.IsZero() {
		staff.HireDate = time.Now()
	}

	// Set a default date of birth if not provided (e.g., 18 years ago)
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
