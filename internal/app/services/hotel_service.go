package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	model "homestay.com/nguyenduy/internal/app/models"
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/request"
)

type HotelService interface {
	CreateHotel(hotel *request.HotelRequest, imageFile multipart.File, fileName string) error
	GetAllHotels() ([]model.Hotel, error)
	GetHotelByID(id uint) (*model.Hotel, error)
	UpdateHotel(id uint, hotel *request.HotelRequest) error
	DeleteHotel(id uint) error
}

type hotelService struct {
	hotelRepo repository.HotelRepository
}

func NewHotelService(hotelRepo repository.HotelRepository) HotelService {
	return &hotelService{
		hotelRepo: hotelRepo,
	}
}

func (s *hotelService) UploadImage(file multipart.File, fileName string) (string, error) {
	// Create uploads directory if it doesn't exist
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Create a new file in the uploads directory
	filePath := filepath.Join(uploadDir, fileName)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	// Return the relative path to the uploaded file
	return filePath, nil
}

func (s *hotelService) CreateHotel(hotel *request.HotelRequest, imageFile multipart.File, fileName string) error {
	imageURL, err := s.UploadImage(imageFile, fileName)
	if err != nil {
		return err
	}

	hotel.Image = imageURL

	return s.hotelRepo.Create(hotel)
}

func (s *hotelService) GetAllHotels() ([]model.Hotel, error) {
	return s.hotelRepo.GetAll()
}

func (s *hotelService) GetHotelByID(id uint) (*model.Hotel, error) {
	return s.hotelRepo.GetByID(id)
}

func (s *hotelService) UpdateHotel(id uint, hotel *request.HotelRequest) error {
	return s.hotelRepo.Update(id, hotel)
}

func (s *hotelService) DeleteHotel(id uint) error {
	return s.hotelRepo.Delete(id)
}
