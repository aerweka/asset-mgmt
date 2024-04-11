package document

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	Index(*fiber.Ctx) error
	CreateDocument(*fiber.Ctx) error
	GetDocument(*fiber.Ctx) error
	UpdateDocument(*fiber.Ctx) error
}
