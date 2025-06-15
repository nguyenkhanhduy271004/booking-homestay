package handlers

import (
	"io"
	"net/http"
	"os"

	service "homestay.com/nguyenduy/internal/app/services"
)

type UploadHandler struct {
	hotelService service.HotelService
}

func NewUploadHandler(hotelService service.HotelService) *UploadHandler {
	return &UploadHandler{
		hotelService: hotelService,
	}
}

func (h *UploadHandler) UploadFile(w http.ResponseWriter, r *http.Request) (string, error) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File upload error", http.StatusBadRequest)
		return "", err
	}
	defer file.Close()

	// Create uploads directory if it doesn't exist
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
		return "", err
	}

	// Create a new file in the uploads directory
	filePath := uploadDir + "/" + handler.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return "", err
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Failed to copy file", http.StatusInternalServerError)
		return "", err
	}

	// Return the relative path to the uploaded file
	return filePath, nil
}
