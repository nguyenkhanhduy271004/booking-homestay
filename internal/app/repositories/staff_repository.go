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
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) StaffRepository {
	return &staffRepository{db: db}
}

func (r *staffRepository) Create(staff *request.StaffRequest) error {
	staffModel := &model.Staff{
		FirstName:   staff.FirstName,
		LastName:    staff.LastName,
		Position:    staff.Position,
		Salary:      staff.Salary,
		DateOfBirth: staff.DateOfBirth,
		Phone:       staff.Phone,
		Email:       staff.Email,
	}
	return r.db.Create(staffModel).Error
}

func (r *staffRepository) GetAll() ([]model.Staff, error) {
	var staff []model.Staff
	if err := r.db.Find(&staff).Error; err != nil {
		return nil, err
	}
	return staff, nil
}

func (r *staffRepository) GetByID(id uint) (*model.Staff, error) {
	var staff model.Staff
	if err := r.db.First(&staff, id).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}

func (r *staffRepository) Update(id uint, staff *request.StaffRequest) error {
	return r.db.Model(&model.Staff{}).Where("id = ?", id).Updates(staff).Error

}

func (r *staffRepository) Delete(id uint) error {
	return r.db.Delete(&model.Staff{}, id).Error
}
