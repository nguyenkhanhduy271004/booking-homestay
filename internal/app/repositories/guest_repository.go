package repositories

import (
	"gorm.io/gorm"
	model "homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/dtos/request"
)

type GuestRepository interface {
	Create(guest *request.GuestRequest) error
	GetAll() ([]model.Guest, error)
	GetByID(id uint) (*model.Guest, error)
	Update(id uint, guest *request.GuestRequest) error
	Delete(id uint) error
}

type guestRepository struct {
	db *gorm.DB
}

func NewGuestRepository(db *gorm.DB) GuestRepository {
	return &guestRepository{db: db}
}

func (r *guestRepository) Create(guest *request.GuestRequest) error {
	guestModel := &model.Guest{
		FirstName: guest.FirstName,
		LastName:  guest.LastName,
		Phone:     guest.Phone,
		Email:     guest.Email,
	}
	return r.db.Create(guestModel).Error
}

func (r *guestRepository) GetAll() ([]model.Guest, error) {
	var guests []model.Guest
	if err := r.db.Find(&guests).Error; err != nil {
		return nil, err
	}
	return guests, nil
}

func (r *guestRepository) GetByID(id uint) (*model.Guest, error) {
	var guest model.Guest
	if err := r.db.First(&guest, id).Error; err != nil {
		return nil, err
	}
	return &guest, nil
}

func (r *guestRepository) Update(id uint, guest *request.GuestRequest) error {
	return r.db.Model(&model.Guest{}).Where("id = ?", id).Updates(map[string]interface{}{
		"first_name": guest.FirstName,
		"last_name":  guest.LastName,
		"phone":      guest.Phone,
		"email":      guest.Email,
	}).Error
}

func (r *guestRepository) Delete(id uint) error {
	return r.db.Delete(&model.Guest{}, id).Error
}
