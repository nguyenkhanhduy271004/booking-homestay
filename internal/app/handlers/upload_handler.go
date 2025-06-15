package handlers

import (
	"io"
	"net/http"
	"os"

	service "homestay.com/nguyenduy/internal/app/services"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) (string, error) {

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File upload error", http.StatusBadRequest)
		return "", err
	}
	defer file.Close()

	tempFile, err := os.CreateTemp("temp-images", handler.Filename)
	if err != nil {
		http.Error(w, "Temp file error", http.StatusInternalServerError)
		return "", err
	}
	defer tempFile.Close()

	io.Copy(tempFile, file)

	url, err := service.UploadImage(tempFile.Name())
	if err != nil {
		http.Error(w, "Cloudinary upload error", http.StatusInternalServerError)
		return "", err
	}

	return url, nil
}
