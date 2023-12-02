package database

import (
	"capstone/config"
	"fmt"

	"github.com/cloudinary/cloudinary-go"
)

func CloudinaryInstance(config config.DatabaseConfig) *cloudinary.Cloudinary {
	var urlCoudinary = fmt.Sprintf("cloudinary://%s:%s@%s",
		config.CDN_API_KEY,
		config.CDN_API_SECRET,
		config.CDN_CLOUD_NAME)

	CDNService, err := cloudinary.NewFromURL(urlCoudinary)
	if err != nil {
		return nil
	}

	return CDNService
}
