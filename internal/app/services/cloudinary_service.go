package services

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"homestay.com/nguyenduy/internal/config"
)

func UploadImage(localPath string) (string, error) {
	cld := config.InitCloudinary()

	resp, err := cld.Upload.Upload(context.Background(), localPath, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
