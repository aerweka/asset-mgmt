package v1

import (
	"asset-management.com/internal/asset-mgmt/pbb"
	"github.com/gofiber/fiber/v2"
)

func NewPbbRoutes(api fiber.Router, h pbb.Handlers) {
	pbb := api.Group("/pbb")

	pbb.Get("/", h.Index)
	pbb.Get("/:id", h.GetPBB)
	pbb.Post("/", h.CreatePBB)
	pbb.Patch("/:id", h.UpdatePBB)
}
