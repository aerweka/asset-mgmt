package cloudinary

import (
	"asset-management.com/config"
	"github.com/cloudinary/cloudinary-go/v2"
)

func InitInstance(cfg config.Config) (*cloudinary.Cloudinary, error) {
	app, err := cloudinary.NewFromParams(cfg.Cloudinary.CloudName, cfg.Cloudinary.APIKey, cfg.Cloudinary.APISecret)
	if err != nil {
		return nil, err
	}

	return app, nil
}
