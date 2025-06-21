package services

import (
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/dtos/request"
	"homestay.com/nguyenduy/internal/dtos/response"
)

type GuestService interface {
	CreateGuest(guest *request.GuestRequest) error
	GetAllGuests() ([]response.GuestResponse, error)
	GetGuestByID(id uint) (*response.GuestResponse, error)
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

func (s *guestService) GetAllGuests() ([]response.GuestResponse, error) {
	guests, err := s.guestRepo.GetAll()

	if err != nil {
		return nil, err
	}

	guestDtos := make([]response.GuestResponse, len(guests))

	for i, guest := range guests {
		guestDtos[i] = response.GuestResponse{
			ID:        guest.ID,
			FirstName: guest.FirstName,
			LastName:  guest.LastName,
		}
	}

	return guestDtos, nil
}

func (s *guestService) GetGuestByID(id uint) (*response.GuestResponse, error) {
	guest, err := s.guestRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	guestDto := response.GuestResponse{
		ID:        guest.ID,
		FirstName: guest.FirstName,
		LastName:  guest.LastName,
		Phone:     guest.Phone,
		Email:     guest.Email,
	}
	return &guestDto, nil
}

func (s *guestService) UpdateGuest(id uint, guest *request.GuestRequest) error {
	return s.guestRepo.Update(id, guest)
}

func (s *guestService) DeleteGuest(id uint) error {
	return s.guestRepo.Delete(id)
}
