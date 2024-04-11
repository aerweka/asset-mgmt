package document

import (
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Usecase interface {
	Index(*fiber.Ctx) ([]model.Document, error)
	CreateDocument(*fiber.Ctx, *model.Document) error
	GetDocument(*fiber.Ctx, uuid.UUID) (*model.Document, error)
	UpdateDocument(*fiber.Ctx, uuid.UUID) error
}
