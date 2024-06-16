package v1

import (
	mainasset "asset-management.com/internal/asset-mgmt/main_asset"
	"github.com/gofiber/fiber/v2"
)

func NewMainAssetRoutes(api fiber.Router, h mainasset.Handlers) {
	pbb := api.Group("/main-asset")

	pbb.Get("/", h.Index)
	pbb.Get("/:id", h.GetMainAsset)
	pbb.Post("/", h.CreateMainAsset)
	pbb.Patch("/:id", h.UpdateMainAsset)
}
