package mainasset

import (
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Usecase interface {
	Index(*fiber.Ctx) ([]*model.PBB, error)
	CreateMainAsset(*fiber.Ctx, *model.PBBRequest) (*model.PBB, error)
	GetMainAsset(*fiber.Ctx, uuid.UUID) (*model.PBB, error)
	UpdateMainAsset(*fiber.Ctx, *model.PBB, uuid.UUID) (*model.PBB, error)
}
