package pbb

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	Index(*fiber.Ctx) error
	CreatePBB(*fiber.Ctx) error
	GetPBB(*fiber.Ctx) error
	UpdatePBB(*fiber.Ctx) error
}
