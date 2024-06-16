package mainasset

import (
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Usecase interface {
	Index(*fiber.Ctx) ([]*model.MainAsset, error)
	CreateMainAsset(*fiber.Ctx, *model.MainAssetRequest) (*model.MainAsset, error)
	GetMainAsset(*fiber.Ctx, uuid.UUID) (*model.MainAsset, error)
	UpdateMainAsset(*fiber.Ctx, *model.MainAsset, uuid.UUID) (*model.MainAsset, error)
}
