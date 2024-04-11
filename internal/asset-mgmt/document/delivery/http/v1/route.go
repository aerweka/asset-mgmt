package v1

import (
	"asset-management.com/internal/asset-mgmt/document"
	"github.com/gofiber/fiber/v2"
)

func NewDocumentRoutes(api fiber.Router, h document.Handlers) {
	document := api.Group("/documents")

	document.Get("/", h.Index)
	document.Get("/:id", h.GetDocument)
	document.Post("/", h.CreateDocument)
	document.Put("/:id", h.UpdateDocument)
}
