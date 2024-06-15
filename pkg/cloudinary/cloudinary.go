package cloudinary

import (
	"context"

	"asset-management.com/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Cloudinary struct {
	App *cloudinary.Cloudinary
}

func InitInstance(cfg config.Config) (*Cloudinary, error) {
	app, err := cloudinary.NewFromParams(cfg.Cloudinary.CloudName, cfg.Cloudinary.APIKey, cfg.Cloudinary.APISecret)
	if err != nil {
		return nil, err
	}

	return &Cloudinary{
		App: app,
	}, nil
}

func (h *Cloudinary) SendImage(image string, publicId string) (*uploader.UploadResult, error) {
	ctx := context.Background()
	res, err := h.App.Upload.Upload(ctx, image, uploader.UploadParams{
		PublicID:     publicId,
		Overwrite:    api.Bool(true),
		ResourceType: "image",
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Cloudinary) SendPdf(base64 string, publicId string) (*uploader.UploadResult, error) {
	ctx := context.Background()
	res, err := h.App.Upload.Upload(ctx, base64, uploader.UploadParams{
		PublicID:     publicId,
		Overwrite:    api.Bool(true),
		ResourceType: "raw",
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// delete asset
/**
var ctx = context.Background()
result, _ := cld.Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{
  PublicIDs: []string{"aa2llhorihuyde4vlawg", "xlzldfihjpl4ymkpvhmn"},
  DeliveryType: "upload",
  AssetType: "image",
})
log.Println(result)
*/
