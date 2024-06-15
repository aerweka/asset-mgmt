package pbb

import (
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Usecase interface {
	Index(*fiber.Ctx) ([]*model.PBB, error)
	CreatePBB(*fiber.Ctx, *model.PBBRequest) (*model.PBB, error)
	GetPBB(*fiber.Ctx, uuid.UUID) (*model.PBB, error)
	UpdatePBB(*fiber.Ctx, *model.PBB, uuid.UUID) (*model.PBB, error)
}
