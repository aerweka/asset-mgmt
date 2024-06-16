package mainasset

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	Index(*fiber.Ctx) error
	CreateMainAsset(*fiber.Ctx) error
	GetMainAsset(*fiber.Ctx) error
	UpdateMainAsset(*fiber.Ctx) error
}
