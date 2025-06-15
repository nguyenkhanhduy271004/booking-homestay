package services

import (
	model "homestay.com/nguyenduy/internal/app/models"
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/request"
)

type GuestService interface {
	CreateGuest(guest *request.GuestRequest) error
	GetAllGuests() ([]model.Guest, error)
	GetGuestByID(id uint) (*model.Guest, error)
	UpdateGuest(id uint, guest *request.GuestRequest) error
	DeleteGuest(id uint) error
}

type guestService struct {
	guestRepo repository.GuestRepository
}

func NewGuestService(guestRepo repository.GuestRepository) GuestService {
	return &guestService{
		guestRepo: guestRepo,
	}
}

func (s *guestService) CreateGuest(guest *request.GuestRequest) error {
	return s.guestRepo.Create(guest)
}

func (s *guestService) GetAllGuests() ([]model.Guest, error) {
	return s.guestRepo.GetAll()
}

func (s *guestService) GetGuestByID(id uint) (*model.Guest, error) {
	return s.guestRepo.GetByID(id)
}

func (s *guestService) UpdateGuest(id uint, guest *request.GuestRequest) error {
	return s.guestRepo.Update(id, guest)
}

func (s *guestService) DeleteGuest(id uint) error {
	return s.guestRepo.Delete(id)
}
