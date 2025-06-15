package services

import (
	model "homestay.com/nguyenduy/internal/app/models"
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/request"
)

type BookingService interface {
	CreateBooking(req *request.BookingRequest) (*model.Booking, error)
	GetAllBookings() ([]model.Booking, error)
	GetBookingByID(id uint) (*model.Booking, error)
	UpdateBooking(id uint, req *request.BookingRequest) (*model.Booking, error)
	DeleteBooking(id uint) error
}

type bookingService struct {
	bookingRepo repository.BookingRepository
}

func NewBookingService(bookingRepo repository.BookingRepository) BookingService {
	return &bookingService{
		bookingRepo: bookingRepo,
	}
}

func (s *bookingService) CreateBooking(req *request.BookingRequest) (*model.Booking, error) {
	booking := &model.Booking{
		CheckinDate:  req.CheckInDate,
		CheckoutDate: req.CheckOutDate,
		GuestID:      req.GuestID,
		RoomID:       req.RoomID,
		PaymentID:    req.PaymentID,
		Status:       req.Status,
	}

	err := s.bookingRepo.Create(booking)
	if err != nil {
		return nil, err
	}

	return s.bookingRepo.GetByID(booking.ID)
}

func (s *bookingService) GetAllBookings() ([]model.Booking, error) {
	return s.bookingRepo.GetAll()
}

func (s *bookingService) GetBookingByID(id uint) (*model.Booking, error) {
	return s.bookingRepo.GetByID(id)
}

func (s *bookingService) UpdateBooking(id uint, req *request.BookingRequest) (*model.Booking, error) {
	booking := &model.Booking{
		CheckinDate:  req.CheckInDate,
		CheckoutDate: req.CheckOutDate,
		GuestID:      req.GuestID,
		RoomID:       req.RoomID,
		PaymentID:    req.PaymentID,
		Status:       req.Status,
	}

	err := s.bookingRepo.Update(id, booking)
	if err != nil {
		return nil, err
	}

	return s.bookingRepo.GetByID(id)
}

func (s *bookingService) DeleteBooking(id uint) error {
	return s.bookingRepo.Delete(id)
}
