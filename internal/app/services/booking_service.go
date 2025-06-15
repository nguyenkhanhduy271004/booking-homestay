package services

import (
	model "homestay.com/nguyenduy/internal/app/models"
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/request"
)

type BookingService interface {
	CreateBooking(booking *request.BookingRequest) error
	GetAllBookings() ([]model.Booking, error)
	GetBookingByID(id uint) (*model.Booking, error)
	UpdateBooking(id uint, booking *request.BookingRequest) error
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

func (s *bookingService) CreateBooking(booking *request.BookingRequest) error {
	bookingModel := &model.Booking{
		GuestID:      booking.GuestID,
		RoomID:       booking.RoomID,
		PaymentID:    booking.PaymentID,
		CheckinDate:  booking.CheckinDate,
		CheckoutDate: booking.CheckoutDate,
		Status:       booking.Status,
	}
	return s.bookingRepo.Create(bookingModel)
}

func (s *bookingService) GetAllBookings() ([]model.Booking, error) {
	return s.bookingRepo.GetAll()
}

func (s *bookingService) GetBookingByID(id uint) (*model.Booking, error) {
	return s.bookingRepo.GetByID(id)
}

func (s *bookingService) UpdateBooking(id uint, booking *request.BookingRequest) error {
	bookingModel := &model.Booking{
		GuestID:      booking.GuestID,
		RoomID:       booking.RoomID,
		PaymentID:    booking.PaymentID,
		CheckinDate:  booking.CheckinDate,
		CheckoutDate: booking.CheckoutDate,
		Status:       booking.Status,
	}
	return s.bookingRepo.Update(id, bookingModel)
}

func (s *bookingService) DeleteBooking(id uint) error {
	return s.bookingRepo.Delete(id)
}
