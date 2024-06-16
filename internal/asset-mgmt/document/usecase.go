package document

import (
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Usecase interface {
	Index(*fiber.Ctx, map[string]string) ([]*model.Document, error)
	CreateDocument(*fiber.Ctx, *model.DocumentRequest) ([]*model.Document, error)
	GetDocument(*fiber.Ctx, uuid.UUID) (*model.Document, error)
	UpdateDocument(*fiber.Ctx, *model.Document, uuid.UUID) (*model.Document, error)
}
