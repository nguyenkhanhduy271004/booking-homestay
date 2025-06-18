package repositories

import (
	"gorm.io/gorm"
	model "homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/request"
)

type StaffRepository interface {
	Create(staff *request.StaffRequest) error
	GetAll() ([]model.Staff, error)
	GetByID(id uint) (*model.Staff, error)
	Update(id uint, staff *request.StaffRequest) error
	Delete(id uint) error
}

type staffRepository struct {
	db              *gorm.DB
	hotelRepository HotelRepository
	userRepository  UserRepository
}

func NewStaffRepository(db *gorm.DB, hotelRepository HotelRepository, userRepository UserRepository) StaffRepository {
	return &staffRepository{db: db, hotelRepository: hotelRepository, userRepository: userRepository}
}

func (r *staffRepository) Create(staff *request.StaffRequest) error {
	hotel, err := r.hotelRepository.GetByID(staff.HotelID)
	if err != nil {
		return err
	}

	user, err := r.userRepository.FindByID(staff.UserID)
	if err != nil {
		return err
	}

	// Check if role exists
	var role model.Role
	if err := r.db.First(&role, staff.RoleID).Error; err != nil {
		return err
	}

	staffModel := &model.Staff{
		HotelID:     hotel.ID,
		Hotel:       *hotel,
		UserID:      user.ID,
		User:        *user,
		FirstName:   staff.FirstName,
		LastName:    staff.LastName,
		Position:    staff.Position,
		Salary:      staff.Salary,
		DateOfBirth: staff.DateOfBirth,
		Phone:       staff.Phone,
		Email:       staff.Email,
		HireDate:    staff.HireDate,
		RoleID:      staff.RoleID,
		Role:        role,
	}
	return r.db.Create(staffModel).Error
}

func (r *staffRepository) GetAll() ([]model.Staff, error) {
	var staff []model.Staff
	if err := r.db.Preload("User").
		Preload("Hotel").
		Preload("Role").
		Find(&staff).Error; err != nil {
		return nil, err
	}
	return staff, nil
}

func (r *staffRepository) GetByID(id uint) (*model.Staff, error) {
	var staff model.Staff
	if err := r.db.Preload("User").
		Preload("Hotel").
		Preload("Role").
		First(&staff, id).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}

func (r *staffRepository) Update(id uint, staff *request.StaffRequest) error {
	// Check if role exists
	var role model.Role
	if err := r.db.First(&role, staff.RoleID).Error; err != nil {
		return err
	}

	// Check if hotel exists
	hotel, err := r.hotelRepository.GetByID(staff.HotelID)
	if err != nil {
		return err
	}

	// Check if user exists
	user, err := r.userRepository.FindByID(staff.UserID)
	if err != nil {
		return err
	}

	return r.db.Model(&model.Staff{}).Where("id = ?", id).Updates(map[string]interface{}{
		"user_id":       user.ID,
		"hotel_id":      hotel.ID,
		"first_name":    staff.FirstName,
		"last_name":     staff.LastName,
		"position":      staff.Position,
		"salary":        staff.Salary,
		"date_of_birth": staff.DateOfBirth,
		"phone":         staff.Phone,
		"email":         staff.Email,
		"hire_date":     staff.HireDate,
		"role_id":       role.ID,
	}).Error
}

func (r *staffRepository) Delete(id uint) error {
	return r.db.Delete(&model.Staff{}, id).Error
}
